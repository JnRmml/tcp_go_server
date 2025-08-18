package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
)

// ReadClose ist Funktion mit sowohl Read als auch Close
func getLinesChannel(f io.ReadCloser) <-chan string {
	out := make(chan string, 1)

	go func() {
		defer f.Close()
		defer close(out)
		str := ""
		for {
			data := make([]byte, 8)
			n, err := f.Read(data)
			if err != nil {
				break
			}

			data = data[:n]
			if i := bytes.IndexByte(data, '\n'); i != -1 {
				str += string(data[:i])
				data = data[i+1:]
				//fmt.Printf("read prnt: %s\n", str)
				out <- str
				str = ""
			}
			str += string(data)
		}
		if len(str) != 0 {
			//fmt.Printf("read: %s\n", str)
			out <- str
		}

	}()
	return out
}

func main() {
	listener, err := net.Listen("tcp", ":42069")
	//f, err := os.Open("message.txt")
	if err != nil {
		log.Fatal("error", "error, err")
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("error", "error, err")
		}
		for line := range getLinesChannel(conn) {
			//fmt.Printf("read: %q\n", line)
			fmt.Printf("read: %s\n", strconv.Quote(line)[1:len(strconv.Quote(line))-1]) //strconv.Quote(line)[1:len(strconv.Quote(line))-1])
		}
	}
}
