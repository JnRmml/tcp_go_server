package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	
	listener, err := net.Listen("tcp", ":3333")
	if err != nil {
		fmt.Println("Fehler beim Starten des Servers", err)
		return
	}
	defer listener.Close()
	fmt.Println("TCP-Server läuft auf Port 3333")
	
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Fehler beim Akzeptieren der Verbindung", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	
	data, err := os.ReadFile("message.txt")
	if err != nil {
		conn.Write([]byte("Fehler beim Lesen der Datei\n"))
		return
	}
	
	conn.Write(data)
	
}
		
