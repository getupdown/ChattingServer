package MessageTreatChain

import (
	"GoTcpServer/Info"
)

type IChain interface {
	Treat(*Info.MessageInfo) (error)
	SetNextChain(IChain)
}
