package recorder

import (
	"net"
	"strconv"
	"time"

	"github.com/oleiade/lane"
)

var queue = lane.NewQueue()
var limit int = 1000

type Record struct {
	Timestamp  string
	UserHash   string
	ClientIP   string
	ClientPort string
	TargetHost string
	TargetPort string
}

func Add(hash string, clientAddr, targetAddr net.Addr) {
	if queue.Size() >= limit {
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
	}
	queue.Enqueue(record)
}

func Remove() interface{} {
	return queue.Dequeue()
}
