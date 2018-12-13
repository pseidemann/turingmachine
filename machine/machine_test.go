package machine_test

import (
	"testing"

	"github.com/pseidemann/turingmachine/machine"
)

func TestBinaryIncrement(t *testing.T) {
	const (
		state0 = "s_0"
		state1 = "s_1"
		state2 = "s_2"
		stateE = "s_e"
	)

	const (
		sym0 = '0'
		sym1 = '1'
		symB = '□'
	)

	trans := machine.TransitionFunction{
		machine.TransIn{state0, sym0}: machine.TransOut{state0, sym0, machine.MoveRight},
		machine.TransIn{state0, sym1}: machine.TransOut{state0, sym1, machine.MoveRight},
		machine.TransIn{state0, symB}: machine.TransOut{state1, symB, machine.MoveLeft},
		machine.TransIn{state1, sym0}: machine.TransOut{state2, sym1, machine.MoveLeft},
		machine.TransIn{state1, sym1}: machine.TransOut{state1, sym0, machine.MoveLeft},
		machine.TransIn{state1, symB}: machine.TransOut{stateE, sym1, machine.MoveNone},
		machine.TransIn{state2, sym0}: machine.TransOut{state2, sym0, machine.MoveLeft},
		machine.TransIn{state2, sym1}: machine.TransOut{state2, sym1, machine.MoveLeft},
		machine.TransIn{state2, symB}: machine.TransOut{stateE, symB, machine.MoveRight},
	}

	m := machine.Machine{
		States:       []string{state0, state1, state2, stateE},
		TapeAlphabet: []rune{sym0, sym1, symB},
		BlankSymbol:  symB,
		InputSymbols: []rune{sym0, sym1},
		InitialState: state0,
		FinalStates:  []string{stateE},
		TransFunc:    trans,
	}

	m.ResetWithTape("101")

	t.Log("initial:", m.GetConfiguration())

	for m.Step() {
		t.Log("step:   ", m.GetConfiguration())
	}

	t.Log("halted: ", m.GetConfiguration())

	if !m.Accepted() {
		t.Error("expected machine to accept the input")
	}

	if m.GetTape() != "110" {
		t.Error("wrong output")
	}
}

func TestWordWrap(t *testing.T) {
	const (
		state0 = "s_0"
		state1 = "s_1"
		state2 = "s_2"
		state3 = "s_3"
		state4 = "s_4"
		state5 = "s_5"
		state6 = "s_6"
		stateE = "s_e"
	)

	const (
		symX = 'X'
		symY = 'Y'
		symL = '<'
		symR = '>'
		symB = '□'
	)

	trans := machine.TransitionFunction{
		// move to the right edge
		machine.TransIn{state0, symX}: machine.TransOut{state0, symX, machine.MoveRight},
		machine.TransIn{state0, symY}: machine.TransOut{state0, symY, machine.MoveRight},
		// wrap right three times
		machine.TransIn{state0, symB}: machine.TransOut{state1, symR, machine.MoveRight},
		machine.TransIn{state1, symB}: machine.TransOut{state2, symR, machine.MoveRight},
		machine.TransIn{state2, symB}: machine.TransOut{state3, symR, machine.MoveLeft},
		// move to the left edge
		machine.TransIn{state3, symX}: machine.TransOut{state3, symX, machine.MoveLeft},
		machine.TransIn{state3, symY}: machine.TransOut{state3, symY, machine.MoveLeft},
		machine.TransIn{state3, symR}: machine.TransOut{state3, symR, machine.MoveLeft},
		// wrap left three times
		machine.TransIn{state3, symB}: machine.TransOut{state4, symL, machine.MoveLeft},
		machine.TransIn{state4, symB}: machine.TransOut{state5, symL, machine.MoveLeft},
		machine.TransIn{state5, symB}: machine.TransOut{state6, symL, machine.MoveRight},
		// move to beginning of word
		machine.TransIn{state6, symL}: machine.TransOut{state6, symL, machine.MoveRight},
		machine.TransIn{state6, symX}: machine.TransOut{stateE, symX, machine.MoveNone},
		machine.TransIn{state6, symY}: machine.TransOut{stateE, symY, machine.MoveNone},
	}

	m := machine.Machine{
		States:       []string{state0, state1, state2, state3, state4, state5, state6, stateE},
		TapeAlphabet: []rune{symX, symY, symL, symR, symB},
		BlankSymbol:  symB,
		InputSymbols: []rune{symX, symY},
		InitialState: state0,
		FinalStates:  []string{stateE},
		TransFunc:    trans,
	}

	// XY

	m.ResetWithTape("XY")

	t.Log("initial:", m.GetConfiguration())

	for m.Step() {
		t.Log("step:   ", m.GetConfiguration())
	}

	t.Log("halted: ", m.GetConfiguration())

	if !m.Accepted() {
		t.Error("expected machine to accept the input")
	}

	if m.GetTape() != "<<<XY>>>" {
		t.Error("wrong output")
	}

	// YYYYX

	m.ResetWithTape("YYYYX")

	t.Log("initial:", m.GetConfiguration())

	for m.Step() {
		t.Log("step:   ", m.GetConfiguration())
	}

	t.Log("halted: ", m.GetConfiguration())

	if !m.Accepted() {
		t.Error("expected machine to accept the input")
	}

	if m.GetTape() != "<<<YYYYX>>>" {
		t.Error("wrong output")
	}
}
