package benchmark

import (
    "net"
    "regexp"
    "strconv"
)


func GetOpenTcpPort() int {
    r, _ := regexp.Compile("([0-9]+)")

    l,_ := net.Listen("tcp",":0")
    defer l.Close()
    addr := l.Addr().String()

    match := r.FindString(addr)
    i, _ := strconv.Atoi(match)

    return i
}
