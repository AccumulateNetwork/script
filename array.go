package script

type Array struct {
	Objects []Object
}

func (a *Array) Execute(state *State) {
	for _, v := range a.Objects {
		v.Execute(state)
	}
}

func (a *Array) Type() int {
	return TArray
}
