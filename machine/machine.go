// Package machine implements a Turing machine.
package machine

import (
	"fmt"

	"github.com/pseidemann/turingmachine/machine/tape"
)

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
	state        string
	tape         tape.Tape
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

// ResetWithTape restarts the machine with a new tape input.
func (m *Machine) ResetWithTape(input string) {
	m.state = m.InitialState
	m.tape = tape.New(m.BlankSymbol, input)
}

// Step reads the symbol currently under the head and decides, while looking at the current state
//     - if it should replace the current symbol
//     - if it should move the head
//     - what the new state is
func (m *Machine) Step() (notHalted bool) {
	out := m.TransFunc[TransIn{m.state, m.tape.GetHead()}]
	if out.isUndefined() {
		return false
	}
	m.tape.SetHead(out.Symbol)
	m.move(out.Move)
	m.state = out.State
	return !m.isFinalState(m.state)
}

// GetConfiguration returns the current configuration of the machine by way of illustration.
func (m *Machine) GetConfiguration() string {
	return fmt.Sprintf(
		"state:%s head:%s tape:%s",
		m.state,
		string(m.tape.GetHead()),
		m.tape.GetContent(),
	)
}

// Accepted returns true if the machine is currently in an accepting state.
func (m *Machine) Accepted() bool {
	return m.isFinalState(m.state)
}

// GetTape returns the tape content with leading and trailing blank symbols removed.
func (m *Machine) GetTape() string {
	return m.tape.GetContent()
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
		m.tape.MoveLeft()
	case MoveRight:
		m.tape.MoveRight()
	}
}

func (t *TransOut) isUndefined() bool {
	return t.State == ""
}
