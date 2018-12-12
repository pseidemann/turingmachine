package machine_test

import (
	"testing"

	"github.com/pseidemann/turingmachine/machine"
)

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

func Test(t *testing.T) {
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
