package main

import (
  "fmt"
  "net"
)

func main() {

  // listen on port 50222 on all interfaces
  pc,err := net.ListenPacket("udp4", "0.0.0.0:50222")

  if err != nil {
    panic(err)
  }

  defer pc.Close()

  buf := make([]byte, 1024) // this shld be plenty of space for Tempest msgs

  for { // forever

    n, _, err := pc.ReadFrom(buf) // read the broadcast messages

    if err != nil { // it is possible there is an error
      panic(err)
    }
    
    fmt.Printf("%s\n", buf[:n]) // output to stdout

  }

}
