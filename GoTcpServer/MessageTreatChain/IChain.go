package MessageTreatChain

import (
	"GoTcpServer/Info"
)

type IChain interface {
	Treat(*Info.MessageInfo, *Info.UserConnection) (error)
	SetNextChain(IChain) IChain
}
