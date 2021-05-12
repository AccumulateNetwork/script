package script

type Dictionary struct {
	dictionary map[[16]byte]IObject
}

func (d *Dictionary) Type() int {
	return TDictionary
}
