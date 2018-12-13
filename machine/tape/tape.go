// Package tape implements the tape of a Turing machine.
package tape

import "strings"

// Tape implements a virtually infinite tape which expands both to left and right when needed.
type Tape struct {
	blank rune
	head  *symbol
}

type symbol struct {
	value  rune
	before *symbol
	after  *symbol
}

// New creates a Tape.
func New(blankSymbol rune, content string) Tape {
	head := &symbol{
		value: []rune(content)[0],
	}

	current := head
	for _, r := range content[1:] {
		after := &symbol{value: r}
		after.before = current
		current.after = after
		current = current.after
	}

	return Tape{
		blank: blankSymbol,
		head:  head,
	}
}

// GetHead returns the symbol at which the head is currently pointing at.
func (t *Tape) GetHead() rune {
	return t.head.value
}

// SetHead changes the symbol at which the head is currently pointing at.
func (t *Tape) SetHead(sym rune) {
	t.head.value = sym
}

// GetContent returns the tape content with leading and trailing blank symbols removed.
func (t *Tape) GetContent() string {
	result := string(t.head.value)
	for before := t.head.before; before != nil; before = before.before {
		result = string(before.value) + result
	}
	for after := t.head.after; after != nil; after = after.after {
		result += string(after.value)
	}
	left := strings.TrimLeft(result, string(t.blank))
	right := strings.TrimRight(left, string(t.blank))
	return right
}

// MoveLeft moves the head one step to the left.
func (t *Tape) MoveLeft() {
	before := t.head.getBefore(t.blank)
	before.after = t.head
	t.head.before = before
	t.head = before
}

// MoveRight moves the head one step to the right.
func (t *Tape) MoveRight() {
	after := t.head.getAfter(t.blank)
	after.before = t.head
	t.head.after = after
	t.head = after
}

func (s symbol) getBefore(blank rune) *symbol {
	if s.before == nil {
		return &symbol{value: blank}
	}
	return s.before
}

func (s symbol) getAfter(blank rune) *symbol {
	if s.after == nil {
		return &symbol{value: blank}
	}
	return s.after
}
