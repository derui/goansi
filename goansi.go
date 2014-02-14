package goansi

import (
	"fmt"
	"strconv"
	"sync"
)

// A form includeing fields that are having to display message.
type Forms struct {
	fields map[string]*Field
}

// The field that is used as field to be able to display message at
// any location in terminal.
type Field struct {
	msg string										// The message to display in this field.
	row int												// The row at this field.
	col int												// The col at this field.
	tag string										// used as index for Forms.fields
	inuse sync.Mutex							// The mutex field to be used to parallel operations with this.
}

// functions for Forms and Field

// Add field to the Form.
func (form *Forms) Add(field Field, index string) {
	if form.fields == nil {
		form.fields = make(map[string]*Field)
	}

	form.fields[index] = &field
}

// Delete field from the Form
func (form *Forms) Delete(index string) {
	if form.fields == nil {
		form.fields = make(map[string]*Field)
	}

	delete(form.fields, index)
}

// These functions are to set display attribute for given string.
// There can apply lapping over each others, Example Underscore(Red("test")), then
// the string of `test' apply attributes that are color of red and add underscore it.
var (
	Reset      = newAttrFunc(0)
	Bright     = newAttrFunc(1)
	Dim        = newAttrFunc(2)
	Underscore = newAttrFunc(4)
	Blink      = newAttrFunc(5)
	Reverse    = newAttrFunc(7)
	Hidden     = newAttrFunc(8)
	Black      = newAttrFunc(30)
	Red        = newAttrFunc(31)
	Green      = newAttrFunc(32)
	Yellow     = newAttrFunc(33)
	Blue       = newAttrFunc(34)
	Magenta    = newAttrFunc(35)
	Cyan       = newAttrFunc(36)
	White      = newAttrFunc(37)
	OnBlack    = newAttrFunc(40)
	OnRed      = newAttrFunc(41)
	OnGreen    = newAttrFunc(42)
	OnYellow   = newAttrFunc(43)
	OnBlue     = newAttrFunc(44)
	OnMagenta  = newAttrFunc(45)
	OnCyan     = newAttrFunc(46)
	OnWhite    = newAttrFunc(47)
)

// These functions are for setting resolution of screen and type to the mode.
var (
	MiniMono       = newResolutionFunc(0)
	MiniColor      = newResolutionFunc(1)
	MediumMono     = newResolutionFunc(2)
	MediumColor    = newResolutionFunc(3)
	Large4Color    = newResolutionFunc(4)
	LargeMono      = newResolutionFunc(5)
	LargeMonoTwice = newResolutionFunc(6)
	EnableWrap     = newResolutionFunc(7)

	LargeColor     = newResolutionFunc(13)
	Large16Color   = newResolutionFunc(14)
	VeryLargeMono  = newResolutionFunc(15)
	VeryLargeColor = newResolutionFunc(16)
	LargestMono    = newResolutionFunc(17)
	LargestColor   = newResolutionFunc(18)
	Medium256Color = newResolutionFunc(19)
)

// A private function to make new function is used to change resolution of the screen.
func newResolutionFunc(num int) func() {
	return func() {
		fmt.Print("=" + toEscSeq(strconv.Itoa(num)) + "h")
	}
}

// A private function to make new function is used to apply attribute to text.
func newAttrFunc(num int) func(string) string {
	return func(text string) string {
		ret := ""
		ret += "\x1b[" + strconv.Itoa(num) + "m"
		ret += text
		ret += "\x1b[0m"
		return ret
	}
}

func toEscSeq(str string) string {
	return "\x1b[" + str
}

// Cursor functions

// Move cursor to (x, y) coodinate on terminal.
func MoveTo(x, y int) {
	fmt.Print(makeMoveTo(x, y))
}

// Hide the cursor
func HideCursor() {
	fmt.Print(toEscSeq("?25l"))
}

// Shows the cursor
func ShowCursor() {
	fmt.Print(toEscSeq("?25h"))
}

func moveCursor(num int, direction string) string {
	ret := toEscSeq(strconv.Itoa(num) + direction)
	return ret
}

func MoveUp(num int) {
	fmt.Print(moveCursor(num, "A"))
}

func MoveDown(num int) {
	fmt.Print(moveCursor(num, "B"))
}

func MoveForward(num int) {
	fmt.Print(moveCursor(num, "C"))
}

func MoveBackward(num int) {
	fmt.Print(moveCursor(num, "D"))
}

func MoveNextLine(num int) {
	fmt.Print(moveCursor(num, "E"))
}

func MovePreviousLine(num int) {
	fmt.Print(moveCursor(num, "F"))
}

func MoveColumn(num int) {
	fmt.Print(moveCursor(num, "G"))
}

// A private function to make sequence to move at coodinate position
func makeMoveTo(x, y int) string {
	ret := toEscSeq(strconv.Itoa(y))
	ret += ";"
	ret += strconv.Itoa(x) + "H"
	return ret
}

// Erase function

// Clears the screen and moves the cursor to the home position
func Erase() {
	fmt.Print(toEscSeq("2J"))
}

// Clears all characters from the cursor position to the end of the line.
// Including the character at the cursor position.
func EraseLineToEnd() {
	fmt.Print(toEscSeq("0K"))
}

func EraseLineToBegin() {
	fmt.Print(toEscSeq("1K"))
}

func EraseWholeLine() {
	fmt.Print(toEscSeq("2K"))
}

// Clear character at the position. Restore cursor position before call this funciton.
func ErasePos(x,y int) {
	MoveTo(x,y)

	fmt.Print(" ")
}









