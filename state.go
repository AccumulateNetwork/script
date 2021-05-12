package script

type State struct {
	Compile      bool         // Are we compiling or executing?
	ReturnStack  []IObject    // When executing, where we go on a return
	CompileStack []IObject    // Where we put stuff to recurse while compiling
	DataStack    []IObject    // The Data Stack used to pass parameters to operators
	Dictionary   []Dictionary // Names are given values by the context defined by a stack of Dictionaries
	Input        []byte       // Input from the input stream
	Output       []byte       // Where compiled code goes
}
