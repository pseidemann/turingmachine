package operator_test

import (
	"testing"

	"github.com/pseidemann/turingmachine/operator"
)

type tmWorking struct {
	steps int
}

func (t *tmWorking) Step() (notHalted bool) {
	t.steps++
	if t.steps == 42 {
		return false
	}
	return true
}

type tmNotHalting struct{}

func (n *tmNotHalting) Step() (notHalted bool) {
	return true
}

func TestWorking(t *testing.T) {
	m := &tmWorking{}
	o := operator.New(m)
	err := o.Run()
	if err != nil {
		t.Error("expected no error")
	}
	if m.steps != 42 {
		t.Error("expected steps to be executed until halt")
	}
}

func TestNotHalting(t *testing.T) {
	m := &tmNotHalting{}
	o := operator.New(m)
	err := o.Run()
	if err != operator.ErrNoHalting {
		t.Error("expected error because machine was never halting")
	}
}
