package main

import (
    "fmt"
    "net"
)

func main() {

    conn, err := net.Dial("tcp", "78.108.218.23:25555")
    if err != nil {
        fmt.Println("Server is Down :C, L bozo ratio")
        return
    } else {
        fmt.Println("Server is Up!")
    }
    defer conn.Close()
}