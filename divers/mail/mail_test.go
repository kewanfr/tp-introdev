package mail

import "testing"

func TestOk(t *testing.T) {
	if !estAdresseValide("loig.jezequel@univ-nantes.fr") {
		t.Fail()
	}
}

func TestNOk(t *testing.T) {
	if estAdresseValide("loig.jezequel@univ-nantes@fr") {
		t.Fail()
	}
}

// Ajoutés après le DS

func TestOk2(t *testing.T) {
	if !estAdresseValide("x@y.zz") {
		t.Fail()
	}
}

func TestNOk2(t *testing.T) {
	if estAdresseValide("x@y.z") {
		t.Fail()
	}
}

func TestOk3(t *testing.T) {
	if !estAdresseValide("loig@univ-nantes.fr") {
		t.Fail()
	}
}

func TestNOk3(t *testing.T) {
	if estAdresseValide("x@y.z.fr") {
		t.Fail()
	}
}

func TestNOk4(t *testing.T) {
	if estAdresseValide("x@fr") {
		t.Fail()
	}
}

func TestNOk5(t *testing.T) {
	if estAdresseValide("@univ-nantes.fr") {
		t.Fail()
	}
}

func TestNOk6(t *testing.T) {
	if estAdresseValide("je@n-fr@ncois@univ-nantes.fr") {
		t.Fail()
	}
}
