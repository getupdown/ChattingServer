package main

import (
	"testing"
	"net"
	"fmt"
	"GoTcpServer/Info"
	"GoTcpServer/Info/InfoParser"
	"time"
)

func TestFunc(t *testing.T) {
	con, err := net.Dial("tcp", ":8080")
	if err != nil {

	} else {
		for i := 0; i < 10; i++ {
			str := "abcdefgasdhfuasdhfuasdjifasdjifasjdifjasdifjiasdfjiasdjfiasdjifjasdifjasdijif"
			con.Write([]byte(str))
		}
	}
	fmt.Printf("sleeping!")
	con.Close()
}


func receive(con net.Conn) {
	rev := make([]byte, 1000)

	for {
		_ , err := con.Read(rev)
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		fmt.Println("Received! : ")
		fmt.Println(string(rev))
	}

}

func TestSingleClient(t *testing.T) {
	con, err := net.Dial("tcp", ":8080")
	if err != nil {
		fmt.Println(err.Error())
	}
	go receive(con)
	parser := &InfoParser.JsonParser{}
	info := &Info.MessageInfo{FromId:"1234", TargetID:"1234", MessageType:"single"}

	//fmt.Println("adsfasdfasdfasd")

	var content string

	content = "abcdefghijk"
	//fmt.Printf("Please enter the content:")
	//fmt.Scanln(&content)
	info.MessageContent = content
	send , _ := parser.Encode(info)
	fmt.Println(string(send))
	con.Write(send)

	for {
		cnt := 1
		cnt = cnt + cnt
	}
	con.Close()
}


func TestHighCnt(t *testing.T) {

	maxh_cnn := 10000

	list := make([]net.Conn, 0)
	for cnt := 0;cnt < maxh_cnn; cnt ++ {
		con, _ := net.Dial("tcp", ":8080")
		list = append(list, con)
	}

	for i := range list {
		fmt.Println(i)
	}
	time.Sleep(1e10)
}

func TestSlice(t *testing.T) {
	myslice := make([]int, 10)

	for i,_ := range myslice {
		myslice = append(myslice, i)
	}
	fmt.Println(len(myslice))
}