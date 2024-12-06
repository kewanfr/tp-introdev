package suite

import (
	"testing"
)

func Test0(t *testing.T) {
	if u(0) != 5 {
		t.Fail()
	}
}

func Test1(t *testing.T) {
	if u(1) != 4 {
		t.Fail()
	}
}

func Test4(t *testing.T) {
	if u(4) != 11 {
		t.Fail()
	}
}

func Test5(t *testing.T) {
	if u(5) != 7 {
		t.Fail()
	}
}

// Ajoutés après le DS

func Test10(t *testing.T) {
	if u(10) != 67 {
		t.Fail()
	}
}

func Test11(t *testing.T) {
	if u(11) != 35 {
		t.Fail()
	}
}
