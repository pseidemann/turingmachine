package tape_test

import (
	"testing"

	"github.com/pseidemann/turingmachine/machine/tape"
)

func Test(t *testing.T) {
	const blank = '□'

	tp := tape.New(blank, "Hello")

	if tp.GetContent() != "Hello" {
		t.Error("tape content not Hello")
	}

	if tp.GetHead() != 'H' {
		t.Error("head not H")
	}

	tp.MoveRight() // e

	if tp.GetHead() != 'e' {
		t.Error("head not e")
	}

	tp.MoveRight() // l
	tp.SetHead('x')
	tp.MoveRight() // l
	tp.SetHead('y')
	tp.MoveRight() // o

	if tp.GetHead() != 'o' {
		t.Error("head not o")
	}

	tp.MoveRight() // blank right 1

	if tp.GetHead() != blank {
		t.Error("head not blank")
	}

	tp.MoveRight() // blank right 2
	tp.MoveRight() // blank right 3

	if tp.GetHead() != blank {
		t.Error("head not blank")
	}

	tp.MoveLeft() // blank right 2
	tp.MoveLeft() // blank right 1
	tp.MoveLeft() // o
	tp.MoveLeft() // y

	if tp.GetHead() != 'y' {
		t.Error("head not y")
	}

	tp.MoveLeft() // x

	if tp.GetHead() != 'x' {
		t.Error("head not x")
	}

	tp.MoveLeft() // e
	tp.MoveLeft() // H

	if tp.GetHead() != 'H' {
		t.Error("head not H")
	}

	tp.MoveLeft() // blank left 1

	if tp.GetHead() != blank {
		t.Error("head not blank")
	}

	tp.MoveLeft() // blank left 2
	tp.MoveLeft() // blank left 3

	if tp.GetHead() != blank {
		t.Error("head not blank")
	}

	if tp.GetContent() != "Hexyo" {
		t.Error("tape content not Hexyo")
	}

}
