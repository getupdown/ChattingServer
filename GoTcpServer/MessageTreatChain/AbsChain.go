package MessageTreatChain

type AbsChain struct {
	IChain
	NextChain IChain
}

func (abc *AbsChain) SetNextChain(nextchain IChain) {
	abc.NextChain = nextchain
}
