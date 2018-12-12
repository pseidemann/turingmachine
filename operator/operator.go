// Package operator implements a controller for a Turing machine.
package operator

import "errors"

const maxSteps = 1000

// ErrNoHalting is returned when it seemed that the machine will never stop.
var ErrNoHalting = errors.New("machine probably never stops")

// Machine is a Turing machine implementation.
type Machine interface {
	Step() (notHalted bool)
}

// An Operator runs a Turing machine.
type Operator struct {
	m Machine
}

// New creates an Operator.
func New(m Machine) *Operator {
	return &Operator{
		m: m,
	}
}

// Run operates the machine until it is finished or the maximum number of steps allowed is exhausted.
func (o *Operator) Run() error {
	numSteps := 0
	for o.m.Step() {
		numSteps++
		if numSteps > maxSteps {
			return ErrNoHalting
		}
	}
	return nil
}
