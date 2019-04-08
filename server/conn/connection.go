package conn

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"io"
	"log"
	"net"
	"sync/atomic"
)

// 系統底層訊息
const SYSTEM_MESSAGE uint8 = 255
const ERROR_RECV uint8 = 254
const (
	MAX_RECV_QUEUE_SIZE = 64
	MAX_SEND_QUEUE_SIZE = 64
	MAX_RECV_BUFF_SIZE  = 65536 * 2
	MSG_SIZE_LEN        = 4
	CACHE_SIZE          = 1024
)

//Connection 為基礎連線結構資料
type Connection struct {
	conn   *net.TCPConn
	closed uint32

	prevMessage *[]byte
}
type IConnection interface {
	SendBytes([]byte)
	DoClose()
	SetClosed() bool
	IsClosed() bool
}

//NewConnection 建立新socket連線物件
func NewConnection(conn net.Conn) *Connection {
	task := new(Connection)
	task.conn = conn.(*net.TCPConn)
	task.closed = 0

	task.conn.SetWriteBuffer(65536)

	//go task.onRecv()

	return task
}

// StartReceiving 利用go routine處理conn接收到的client訊息
func (client *Connection) StartReceiving(recv chan<- []byte) {
	cache := make([]byte, CACHE_SIZE)
	buf := bytes.NewBuffer(make([]byte, 0, MAX_RECV_BUFF_SIZE))
	bufReader := bufio.NewReader(client.conn)
	var contentLen uint32
	packByteSize := make([]byte, MSG_SIZE_LEN)

	defer func() {
		if r := recover(); r != nil {
			log.Printf("recover onRecv error: %v \n", r.(error))
		}
	}()
	for {

		size, err := bufReader.Read(cache)

		if err == io.EOF {
			log.Println("Connection Closed")
			recv <- []byte{SYSTEM_MESSAGE}
			client.DoClose()
			break
		} else if err != nil {
			log.Printf("Connection Recv error: %v", err)
			recv <- []byte{SYSTEM_MESSAGE}
			client.DoClose()
			break
		}

		if size == 0 {
			log.Printf("packet size == 0, skip it\n")
			continue
		}

		buf.Write(cache[:size])
		for {
			//本次缓冲区数据包正好读完，重置内容长度
			if buf.Len() == 0 {
				contentLen = 0
				break
			}

			// 开始读取一个新的数据包
			if contentLen == 0 {
				// 判断缓冲区剩余数据是否足够读取一个包长
				if buf.Len() < MSG_SIZE_LEN {
					break
				}
				_, err = buf.Read(packByteSize)
				contentLen = binary.LittleEndian.Uint32(packByteSize)
				if contentLen > 50000 {
					if client.prevMessage != nil {
						log.Printf("wierd content size: %d, previous: %v", contentLen, *client.prevMessage)
					} else {
						log.Printf("wierd content size: %d, and prevMessage is nil now", contentLen)
					}

					recv <- []byte{ERROR_RECV}
				}
			}

			//判断缓冲区剩余数据是否足够读取一个完整的包
			//true -> 继续读取(contentLen - buf.Len())长度的字节数据
			if int(contentLen) > buf.Len() || contentLen == 0 {
				break
			}

			data := make([]byte, contentLen)
			//data为完整数据包
			_, err = buf.Read(data)
			msg := make([]byte, contentLen)
			copy(msg, data[:contentLen])
			client.prevMessage = &msg
			recv <- msg

			contentLen = 0
		}
	}
}

//Send 送字串封包
func (client *Connection) Send(message string) {
	client.SendBytes([]byte(message))
}

//SendBytes 推送byte封包
func (client *Connection) SendBytes(b []byte) {
	if client.IsClosed() {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			log.Printf("recover SendBytes error: %v \n", r.(error))
		}
	}()
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, uint32(len(b)))
	if err != nil {
		log.Println(err.Error())
		return
	}
	buf.Write(b)
	client.conn.Write(buf.Bytes())
}

//DoClose 當client斷線的處理
func (client *Connection) DoClose() {
	if client.SetClosed() {
		client.conn.Close()

		client.prevMessage = nil
		//client.conn = nil
	}
}

//SetClosed 設定client斷線旗標
func (client *Connection) SetClosed() bool {
	if atomic.CompareAndSwapUint32(&client.closed, 0, 1) {
		return true
	}

	return false
}

//IsClosed 取得斷線旗標
func (client *Connection) IsClosed() bool {
	return (atomic.LoadUint32(&client.closed) == 1)
}
