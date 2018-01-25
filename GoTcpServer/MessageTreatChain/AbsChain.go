package MessageTreatChain

type AbsChain struct {
	IChain
	NextChain IChain
}

func (abc *AbsChain) SetNextChain(nextchain IChain) IChain {
	abc.NextChain = nextchain
	return nextchain
}
