package lib

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

const port = "8080"

//RunHost  takes as ip argument
//and  listen fo connection with thr ip
func RunHost(ip string) {
	ipAndPort := ip + ":" + port
	listener, listenErr := net.Listen("tcp", ipAndPort)
	if listenErr != nil {
		log.Fatal("Error : ", listenErr)
	}
	fmt.Println("Listening on ", ipAndPort)
	conn, acceptErr := listener.Accept()
	if acceptErr != nil {
		log.Fatal("Error : ", acceptErr)
	}
	fmt.Println("New connection accepted")
	for {
		handleHost(conn)
	}

}
func handleHost(conn net.Conn) {
	reader := bufio.NewReader(conn)
	message, readErr := reader.ReadString('\n')
	if readErr != nil {
		log.Fatal("Error: ", readErr)
	}
	fmt.Println("Message received : ", message)
	fmt.Print("Send message : ")
	replayReader := bufio.NewReader(os.Stdin)
	replyMessage, replyError := replayReader.ReadString('\n')
	if replyError != nil {
		log.Fatal("Error :", replyError)
	}
	fmt.Fprint(conn, replyMessage)
}

//RunGuest conect this guest with our server
func RunGuest(ip string) {
	ipAndPort := ip + ":" + port
	conn, dialErr := net.Dial("tcp", ipAndPort)
	if dialErr != nil {
		log.Fatal("Error : ", dialErr)
	}
	for {
		handleGuest(conn)
	}

}

func handleGuest(conn net.Conn) {
	fmt.Print("Send message :")
	reader := bufio.NewReader(os.Stdin)
	message, readErr := reader.ReadString('\n')
	if readErr != nil {
		log.Fatal("Error", readErr)
	}
	fmt.Fprint(conn, message)
	replyReader := bufio.NewReader(conn)
	replyMessage, replyErr := replyReader.ReadString('\n')
	if replyErr != nil {
		log.Fatal("Error : ", replyErr)
	}
	fmt.Println("Message received :", replyMessage)
}
