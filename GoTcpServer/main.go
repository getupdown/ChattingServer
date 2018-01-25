package main

import (
	"net"
	"fmt"
	"GoTcpServer/MessageTreatChain"
	"GoTcpServer/Info/InfoParser"
	"GoTcpServer/Info"
)

const (
	maxByteCount = 1000
)


//func handleConnection(con net.Conn, mp *MessageTreatChain.GlobalSocketMap) {
//	var tmp InfoParser.Parser = InfoParser.JsonParser{}
//	parsechain := MessageTreatChain.ParseChain{Content:make([]byte, maxByteCount), Parser:tmp}
//	transportchain := MessageTreatChain.TransportChain{GlobalSocketMap:mp, Parser:tmp}
//	parsechain.SetNextChain(&transportchain)
//	transportchain.SetNextChain(nil)
//
//	fmt.Println("Connected from " + con.RemoteAddr().String())
//	rev := make([]byte, maxByteCount)
//
//	for {
//		cnt, err := con.Read(rev)
//		if err != nil {
//			if err == io.EOF {
//				break
//			}
//			fmt.Printf(err.Error())
//		} else {
//			info := &Info.MessageInfo{}
//			parsechain.Content = rev[:cnt]
//			parsechain.Treat(info)
//		}
//	}
//	fmt.Println(con.RemoteAddr().String() + " Connection over!")
//}


func handleConnection(conn Info.UserConnection, mp *MessageTreatChain.GlobalSocketMap) {
	rev := make([]byte, maxByteCount)

	var jsonparser InfoParser.Parser = &InfoParser.JsonParser{}

	transporter := MessageTreatChain.TransportChain{GlobalSocketMap:mp, Parser:jsonparser}


	for {
		if _, err := conn.Conn.Read(rev); err != nil {
			fmt.Println(err.Error())
			break
		} else {
			info := &Info.MessageInfo{}
			if _, err = jsonparser.Encode(info) ; err != nil {
				fmt.Println(err.Error())
				break
			}

			if conn.HasLogged() {
				transporter.Treat(info, &conn)
			} else {

			}
		}
	}
}

func main() {
	MAP := &MessageTreatChain.GlobalSocketMap{Socketmap:make(map[string]net.Conn)}

	listner, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Printf(err.Error())
	} else {
		for {
			con, err := listner.Accept()
			if err != nil {
				fmt.Printf(err.Error())
			} else {
				go handleConnection(Info.NewUserConnection(con, Info.User{}), MAP)
			}
		}
	}
}

