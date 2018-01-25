package MessageTreatChain

import (
	"GoTcpServer/Info"
	"GoTcpServer/Info/InfoParser"
	"fmt"
)

type TransportChain struct {
	AbsChain
	GlobalSocketMap *GlobalSocketMap
	Parser InfoParser.Parser
}


func (t *TransportChain) Treat(info *Info.MessageInfo, connection *Info.UserConnection) (err error) {
	//get the conn of the receiver
	con, err := t.GlobalSocketMap.Get(info.TargetID)
	if err != nil {
		fmt.Print(err.Error())
		return err
	}

	//encode the message into the format decided by the parser
	content , err := (t.Parser).Encode(*info)
	if err != nil {
		fmt.Print(err.Error())
		return err
	}

	con.Write(content)
	return nil
}

