package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"runtime"
	"strconv"
	"sync"
)

const (
	PORT = 3540
	kMax = 100000
	kMin = 10000
)

var mu sync.Mutex
var num_connections int

func getGOMAXPROCS() int {
	return runtime.GOMAXPROCS(0)
}

func increaseConnectionCount() {
	mu.Lock()
	num_connections++
	if num_connections%100 == 0 {
		log.Printf("Total number of connections: %v", num_connections)
	}

	mu.Unlock()
}

func decreaseConnectionCount() {
	mu.Lock()
	num_connections--
	mu.Unlock()
}

func main() {
	fmt.Printf("GOMAXPROCS is %d\n", getGOMAXPROCS())
	runtime.GOMAXPROCS(16)
	num_connections = 0
	server, err := net.Listen("tcp", ":"+strconv.Itoa(PORT))
	if server == nil || err != nil {
		panic("couldn't start listening")
	}
	conns := clientConns(server)
	for {
		go handleConn(<-conns)
	}
}

func clientConns(listener net.Listener) chan net.Conn {
	ch := make(chan net.Conn)
	i := 0
	go func() {
		for {
			client, err := listener.Accept()
			if client == nil || err != nil {
				//fmt.Printf("couldn't accept: %v", err)
				continue
			}
			i++
			//fmt.Printf("%d: %v <-> %v\n", i, client.LocalAddr(), client.RemoteAddr())
			ch <- client
		}
	}()
	return ch
}

func handleConn(client net.Conn) {
	increaseConnectionCount()
	b := bufio.NewReader(client)
	data := make([]byte, 2048)
	for {
		n, err := b.Read(data)
		if err != nil { // EOF, or worse
			break
		}
		client.Write(data[:n])
	}
	decreaseConnectionCount()
}
