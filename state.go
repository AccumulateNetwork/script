package script

import "strconv"

const (
	Invalid       int = iota // 0 is invalid.
	TDictionary              // Dictionaries
	TArray                   // Data
	TCArray                  // Compiled Arrays (smaller footprint arrays for execution)
	TInt                     // Integer
	TFloat                   // Floating point
	TCount                   // Number of object types
)

const (
	opAdd                = 0x01
	opSub                = 0x02
	opMul                = 0x03
	opDiv                = 0x04
	opLitInt             = 0x05
	opLitFloat           = 0x06
	opLitArray           = 0x07
	opLitExecutableArray = 0x08
)

// All objects must implement the Object Interface
type Object interface {
	SetExecuteable(executable bool) // Sets the object's executable flag, if it has one
	Execute(state *State)           // The Execution Behavior of  an object
	Type() int                      // Returns the type of the object
}

// State
// Holds the entire execution state of the interpreter.
type State struct {
	Dictionary []Object
	DataStack  []Object
}

// GetToken
// Parse out the next token in a list of bytes.  Returns nil  if no more tokens remain
func GetToken(input []byte) (remaining []byte, token []byte) {

WhiteSpaceLoop: //                            First of all, remove all whitespace
	for { //                                  Loop through the input

		if len(input) == 0 { //               If we have no more input, then no tokens remain in the input
			return nil, nil
		}

		switch input[0] { //                  Look at the next character
		case '\r', '\n', ' ', '\t': //          If whitespace, then just ignore it by removing it
			input = input[1:]
		default: //                           If anything else, then we are done looking at whitespace
			break WhiteSpaceLoop
		}
	}

	//                                        Now pull out the token
	for len(input) > 0 { //                   Collect charaacters unil end of input or whitespace
		switch input[0] {
		case '\r', '\n', ' ', '\t': //        If we find whitespace, we are done
			return token, input[1:] //        Don't keep the whitespace in the input; we don't have to look at it again
		default:
			token = append(token, input[0]) // Otherwise add the character to the token
			input = input[1:]               // Remove the character from input, and loop.
		}
	}
	return token, input //                     Return the token we extracted, and the zero length input
}

// Compile
// Compiles an input string into an array of objects
// If the input stream is invalid or only whitespace, an empty array is returned
// If the end of input was encountered, then the delimiter is nil
// If a close } or ] is encountered, then we return with more input to parse, and we return } or ] as the delimiter
// When an open { or [ is encountered, we recurse compile to get the array object
func (s *State) Compile(input []byte) (*CArray, delimiter []byte) {

	var token []byte //                            Place to keep parsed tokens

	for { //                                       Loop through the tokens in the input
		token, input = GetToken(input) //          Get the next token
		if token == nil {              //          If the token is nil, then return array and nil
			return array, nil //
		}
		iv, err := strconv.ParseInt(token,0,64)
		if err == nil {
			array =
		}
	}
	return nil
}
