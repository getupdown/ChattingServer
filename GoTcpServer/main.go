package main

import (
	"net"
	"fmt"
	"io"
	"GoTcpServer/MessageTreatChain"
	"GoTcpServer/Info/InfoParser"
	"GoTcpServer/Info"
)

const (
	maxByteCount = 1000
)

func handleConnection(con net.Conn, mp *MessageTreatChain.GlobalSocketMap) {
	var tmp InfoParser.Parser = InfoParser.JsonParser{}
	parsechain := MessageTreatChain.ParseChain{Content:make([]byte, maxByteCount), Parser:tmp}
	transportchain := MessageTreatChain.TransportChain{GlobalSocketMap:mp, Parser:tmp}
	parsechain.SetNextChain(&transportchain)
	transportchain.SetNextChain(nil)

	fmt.Println("Connected from " + con.RemoteAddr().String())
	rev := make([]byte, maxByteCount)

	for {
		cnt, err := con.Read(rev)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Printf(err.Error())
		} else {
			info := &Info.MessageInfo{}
			parsechain.Content = rev[:cnt]
			parsechain.Treat(info)
		}
	}
	fmt.Println(con.RemoteAddr().String() + " Connection over!")
}

func main() {
	mp := &MessageTreatChain.GlobalSocketMap{Socketmap:make(map[string]net.Conn)}
	listner, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Printf(err.Error())
	} else {
		cnt := 0
		for {
			con, err := listner.Accept()
			if err != nil {
				fmt.Printf(err.Error())
			} else {
				mp.Set(string(cnt), con)
				cnt ++
				go handleConnection(con, mp)
			}
		}
	}
}

