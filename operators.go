// Copyright 2016 Steven Oud. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be found
// in the LICENSE file.

package mathcat

type association int

const (
	ASSOC_NONE association = iota
	ASSOC_LEFT
	ASSOC_RIGHT
)

type operator struct {
	prec  int
	assoc association
	unary bool
}

var ops = map[tokenType]operator{
	// Assignment operators
	EQ:     {0, ASSOC_RIGHT, false}, // =
	ADD_EQ: {0, ASSOC_RIGHT, false}, // +=
	SUB_EQ: {0, ASSOC_RIGHT, false}, // -=
	DIV_EQ: {0, ASSOC_RIGHT, false}, // /=
	MUL_EQ: {0, ASSOC_RIGHT, false}, // *=
	POW_EQ: {0, ASSOC_RIGHT, false}, // **=
	REM_EQ: {0, ASSOC_RIGHT, false}, // %=
	AND_EQ: {0, ASSOC_RIGHT, false}, // &=
	OR_EQ:  {0, ASSOC_RIGHT, false}, // |=
	XOR_EQ: {0, ASSOC_RIGHT, false}, // ^=
	LSH_EQ: {0, ASSOC_RIGHT, false}, // <<=
	RSH_EQ: {0, ASSOC_RIGHT, false}, // >>=

	// Relational operators
	EQ_EQ: {0, ASSOC_RIGHT, false}, // ==
	GT:    {0, ASSOC_RIGHT, false}, // >
	GT_EQ: {0, ASSOC_RIGHT, false}, // >=
	LT:    {0, ASSOC_RIGHT, false}, // <
	LT_EQ: {0, ASSOC_RIGHT, false}, // <=

	// Bitwise operators
	OR:  {1, ASSOC_RIGHT, false}, // |
	XOR: {2, ASSOC_RIGHT, false}, // ^
	AND: {3, ASSOC_RIGHT, false}, // &
	LSH: {4, ASSOC_RIGHT, false}, // <<
	RSH: {4, ASSOC_RIGHT, false}, // >>
	NOT: {8, ASSOC_LEFT, true},   // ~

	// Mathematical operators
	ADD:       {5, ASSOC_LEFT, false}, // +
	SUB:       {5, ASSOC_LEFT, false}, // -
	MUL:       {6, ASSOC_LEFT, false}, // *
	DIV:       {6, ASSOC_LEFT, false}, // /
	POW:       {7, ASSOC_LEFT, false}, // **
	REM:       {6, ASSOC_LEFT, false}, // %
	UNARY_MIN: {9, ASSOC_LEFT, true},  // -
}
