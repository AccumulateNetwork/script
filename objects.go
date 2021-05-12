package script

type IObject interface {
	Type() int // Returns the type of the object
	Execute()  // Defines the execution behavior of an object.
	String()
}

const (
	TInt        = 1 // Tags for basic types
	TFloat      = 2
	TString     = 3
	TName       = 4
	TArray      = 5
	TOperator   = 6
	TDictionary = 7
)

var _ = TInt + TFloat + TString + TName + TArray + TOperator + TDictionary
