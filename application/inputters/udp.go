// Used code from: https://varshneyabhi.wordpress.com/2014/12/23/simple-udp-clientserver-in-golang/
package listeners

//func main() {
//    ServerAddr, err := net.ResolveUDPAddr("udp", ":8000")
//    if err != nil {
//        fmt.Println(err)
//    }
//
//    ServerConn, err := net.ListenUDP("udp", ServerAddr)
//    if err != nil {
//        fmt.Println(err)
//    }
//    defer ServerConn.Close()
//
//    buf := make([]byte, 1024)
//
//    for {
//        n, addr, err := ServerConn.ReadFromUDP(buf)
//        fmt.Println("Received ", string(buf[0:n]), " from ", addr)
//
//        if err != nil {
//            fmt.Println("Error: ", err)
//        }
//    }
//}


