package script

type Name struct {
	ID [16]byte
}

func (n *Name) Type() int {
	return TName
}

func (n *Name) Execute(s *State) {

}
