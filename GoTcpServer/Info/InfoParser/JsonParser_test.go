package InfoParser

import (
	"testing"
	"fmt"
	"GoTcpServer/Info"
)

func TestJsonParser_Parse(t *testing.T) {
	str := make([]string, 50)
	str = append(str, "{")

	str = append(str, "\"FromId\":\"userid\",")
	str = append(str, "\"MessageType\":\"personal\",")
	str = append(str, "\"TargetId\":\"targetid\",")
	str = append(str, "\"MessageContent\":\"what's wrong?\"")
	str = append(str, "}")

	//{
	//	"fromId":"userid",
	//	"messageType" : "personal",
	//	"targetId":"targetid",
	//	"encryptedAlgorithm":"RSA",
	//	"messageContent":"what's wrong?"
	//}



	tmpstr := " "
	for _, v := range str {
		tmpstr = tmpstr + v
	}

	ps := &JsonParser{}
	info := Info.MessageInfo{}
	err := ps.Parse([]byte(tmpstr), &info)

	if err != nil {
		fmt.Printf(err.Error())
	}

	fmt.Print(info.FromId)
	fmt.Print(info.MessageContent)
	fmt.Print(info.MessageType)
	fmt.Print(info.TargetID)
}
