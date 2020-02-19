package main
import (
    "fmt"
    "net"
    "bufio"
    "os"
)

func main() {
    p :=  make([]byte, 1472)
    conn, err := net.Dial("udp", "127.0.0.1:1053")
    if err != nil {
        fmt.Printf("Some error %v", err)
        return
    }
    msg := os.Args[1]
    fmt.Fprintf(conn, msg)
    defer conn.Close()
    _, err = bufio.NewReader(conn).Read(p)
    if err == nil {
        fmt.Printf("%s\n", p)
    } else {
        fmt.Printf("Some error %v\n", err)
    }
}
