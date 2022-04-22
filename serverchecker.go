package main
import (
    "fmt"
    "net"
    "os"
)
func main() {
        ip:= os.Args[1]
        port := os.Args[2]

        ipandport := ip+":"+port


    conn, err := net.Dial("tcp", ipandport, )
    if err != nil {
        fmt.Println("Server is Down :C, L bozo ratio")
        return
    } else {
        fmt.Println(ipandport, "is online!")
    }
    defer conn.Close()
}
