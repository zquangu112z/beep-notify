package beep

import (
	"testing"
)

func TestBeep(t *testing.T) {
	Beep(1)
	Beep(2)
	Beep(3, 4)
}

func TestBeepNoID(t *testing.T) {
	Beep()
}
