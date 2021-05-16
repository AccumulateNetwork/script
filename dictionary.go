package script

type Dictionary struct {
	Dict map[[16]byte]Object
}

func (d *Dictionary) Execute(state *State) {
	state.DataStack = append(state.Dictionary, d)
}

func (d *Dictionary) Type() int {
	return TDictionary
}
