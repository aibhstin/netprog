package main

import (
    "net"
    "os"
    "fmt"
)

func main() {
    if len(os.Args) != 2 {
        fmt.Fprintf(os.Stderr, "Usage: %s host:port\n", os.Args[0])
        os.Exit(1)
    }

    service := os.Args[1]

    udpAddr, err := net.ResolveUDPAddr("udp4", service)
    checkError(err)

    conn, err := net.DialUDP("udp", nil, udpAddr)
    checkError(err)

    _, err = conn.Write([]byte("anything"))
    checkError(err)

    var buf = make([]byte, 512)
    n, err := conn.Read(buf[0:])
    checkError(err)

    fmt.Printf("%s\n", string(buf[0:n]))

    os.Exit(1)
}

func checkError(err error) {
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
        os.Exit(1)
    }
}
