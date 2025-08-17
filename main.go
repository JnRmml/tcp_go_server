package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	targetIP := "192.168.178.46:42069"
	
	conn, err := net.Dial("tcp", targetIP)
	if err != nil {
		fmt.Println("Fehler beim Verbinden", err)
		return
	}
	defer conn.Close()
	fmt.Println("TCP-Client hat auf Port 42069 gesendet")
	
	data, err := os.ReadFile("message.txt")
	if err != nil {
		conn.Write([]byte("Fehler beim Lesen der Datei\n"))
		return
	}
	
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("Fehler beim Senden der Daten:", err)
	}
	fmt.Println("Daten erfolgreich an", targetIP, " gesendet")
	
}
		
