package goansi

import (
	"fmt"
	"testing"
)

type Mapping struct {
	Func  func(string) string
	Param int
}

func TestAddingAttributesToText(t *testing.T) {
	mapping := map[string]Mapping{
		"Reset":      {Reset, 0},
		"Bright":     {Bright, 1},
		"Dim":        {Dim, 2},
		"Underscore": {Underscore, 4},
		"Blink":      {Blink, 5},
		"Reverse":    {Reverse, 7},
		"Hidden":     {Hidden, 8},
		"Black":      {Black, 30},
		"Red":        {Red, 31},
		"Green":      {Green, 32},
		"Yellow":     {Yellow, 33},
		"Blue":       {Blue, 34},
		"Magenta":    {Magenta, 35},
		"Cyan":       {Cyan, 36},
		"White":      {White, 37},
		"OnBlack":    {OnBlack, 40},
		"OnRed":      {OnRed, 41},
		"OnGreen":    {OnGreen, 42},
		"OnYellow":   {OnYellow, 43},
		"OnBlue":     {OnBlue, 44},
		"OnMagenta":  {OnMagenta, 45},
		"OnCyan":     {OnCyan, 46},
		"OnWhite":    {OnWhite, 47},
	}

	for key, v := range mapping {

		ret := v.Func(key)

		if ret != fmt.Sprintf("\x1b[%dm%s\x1b[0m", v.Param, key) {
			t.Errorf("%s is not match attribute %d => %s", key, v.Param, ret)
		}
	}
}

func TestCursorPositionChanging(t *testing.T) {
	ret := makeMoveTo(10, 20)

	if ret != "\x1b[20;10H" {
		t.Errorf("Can not move to (%d, %d) on coodinate", 10, 20)
	}
}

func TestAppendEscapeSequenseToString(t *testing.T) {
	ret := toEscSeq("test")

	if ret != "\x1b[test" {
		t.Errorf(`toEscSeq function should add to prefix to given string. Returned string has not prefix`)
	}
}

func TestCsontructCursorMovementInstruction(t *testing.T) {
	ret := moveCursor(1, "A")

	if ret != "\x1b[1A" {
		t.Errorf(`moveCussor must have contain number of cursor moving and character for direction`)
	}

	ret = moveCursor(4, "A")

	if ret != "\x1b[4A" {
		t.Errorf(`moveCursor should be able to move multiple number of cursor position`)
	}
}
