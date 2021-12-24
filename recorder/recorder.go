package recorder

import (
	"net"
	"strconv"
	"time"

	"github.com/oleiade/lane"
)

var queue = lane.NewQueue()
var Capacity int = 0

type Record struct {
	Timestamp  string
	UserHash   string
	ClientIP   string
	ClientPort string
	TargetHost string
	TargetPort string
	Transport  string
}

func Add(hash string, clientAddr, targetAddr net.Addr, transport string) {
	if queue.Size() >= Capacity {
		return
	}

	clientIP, clientPort, _ := net.SplitHostPort(clientAddr.String())
	targetHost, targetPort, _ := net.SplitHostPort(targetAddr.String())

	record := Record{
		Timestamp:  strconv.Itoa(int(time.Now().UnixMilli())),
		UserHash:   hash,
		ClientIP:   clientIP,
		ClientPort: clientPort,
		TargetHost: targetHost,
		TargetPort: targetPort,
		Transport:  transport,
	}
	queue.Enqueue(record)
}

func Remove() interface{} {
	return queue.Dequeue()
}
