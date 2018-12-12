// Package machine implements a Turing machine.
package machine

import "fmt"

// A Movement describes in which direction the machine's head will move.
type Movement int

// All possible movements of the machine's head.
const (
	MoveNone Movement = iota
	MoveLeft
	MoveRight
)

// Machine is a Turing machine.
type Machine struct {
	States       []string
	TapeAlphabet []rune
	BlankSymbol  rune
	InputSymbols []rune
	InitialState string
	FinalStates  []string
	TransFunc    TransitionFunction
	conf         configuration
}

// TransitionFunction is a map which describes what the machine should do next for a given state and symbol.
type TransitionFunction map[TransIn]TransOut

// TransIn is an input for the transition function.
type TransIn struct {
	State  string
	Symbol rune
}

// TransOut is an output from transition function.
type TransOut struct {
	State  string
	Symbol rune
	Move   Movement
}

// Configuration defines in what state the machine is, which symbols are on the
// tape and what the position of the head is.
// The position of the head is determined implicitly by defining which symbols
// are before the symbol at which the head is pointing at, and which symbols are
// after that, including the symbol the head is pointing at.
type configuration struct {
	BeforeHead []rune
	State      string
	Head       []rune
}

// ResetWithTape restarts the machine with a new tape input.
func (m *Machine) ResetWithTape(input string) {
	m.conf = configuration{
		BeforeHead: []rune{m.BlankSymbol},
		State:      m.InitialState,
		Head:       append([]rune(input), m.BlankSymbol),
	}
}

// Step reads the symbol currently under the head and decides, while looking at the current state
//     - if it should replace the current symbol
//     - if it should move the head
//     - what the new state is
func (m *Machine) Step() (notHalted bool) {
	out := m.TransFunc[TransIn{m.conf.State, m.conf.Head[0]}]
	if out.isUndefined() {
		return false
	}
	m.conf.State = out.State
	m.conf.Head[0] = out.Symbol
	m.move(out.Move)
	return !m.isFinalState(out.State)
}

// GetConfiguration returns the current configuration of the machine by way of illustration.
func (m *Machine) GetConfiguration() string {
	return fmt.Sprintf(
		"state:%s tape:%s[%s]%s",
		m.conf.State,
		string(m.conf.BeforeHead),
		string(m.conf.Head[0]),
		string(m.conf.Head[1:]),
	)
}

// Accepted returns true if the machine is currently in an accepting state.
func (m *Machine) Accepted() bool {
	return m.isFinalState(m.conf.State)
}

// GetTape returns what currently on the tape is, without leading nor trailing blank symbols.
func (m *Machine) GetTape() string {
	var before []rune
	leftBlankArea := false
	for _, r := range m.conf.BeforeHead {
		if r != m.BlankSymbol || leftBlankArea {
			leftBlankArea = true
			before = append(before, r)
		}
	}

	leftBlankArea = false
	var after []rune
	for i := len(m.conf.Head) - 1; i >= 0; i-- {
		r := m.conf.Head[i]
		if r != m.BlankSymbol || leftBlankArea {
			leftBlankArea = true
			after = append([]rune{r}, after...)
		}
	}

	return string(append(before, after...))
}

func (m *Machine) isFinalState(state string) bool {
	for _, s := range m.FinalStates {
		if state == s {
			return true
		}
	}
	return false
}

func (m *Machine) move(move Movement) {
	switch move {
	case MoveLeft:
		lastIdx := len(m.conf.BeforeHead) - 1
		m.conf.Head = append([]rune{m.conf.BeforeHead[lastIdx]}, m.conf.Head...)
		m.conf.BeforeHead = m.conf.BeforeHead[:lastIdx]
	case MoveRight:
		m.conf.BeforeHead = append(m.conf.BeforeHead, m.conf.Head[0])
		m.conf.Head = m.conf.Head[1:]
	}
}

func (t *TransOut) isUndefined() bool {
	return t.State == ""
}
