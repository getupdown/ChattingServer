package MessageTreatChain

import (
	"GoTcpServer/Info/InfoParser"
	"GoTcpServer/Info"
	"fmt"
)

type ParseChain struct {
	AbsChain
	Content []byte
	Parser InfoParser.Parser
}

func (p *ParseChain) Treat(info *Info.MessageInfo) (err error) {
	err = (p.Parser).Parse(p.Content, info)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return p.NextChain.Treat(info)
}





