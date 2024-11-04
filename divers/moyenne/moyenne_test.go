package moyenne

import "testing"

func TestMoyenne(t *testing.T) {
	notes := []float64{5, 10, 15}

	if moyenne(notes) != 10 {
		t.Fail()
	}
}

// Ajoutés après le test

func TestMoyenne2(t *testing.T) {
	notes := []float64{5}

	if moyenne(notes) != 5 {
		t.Fail()
	}
}

func TestMoyenne3(t *testing.T) {
	notes := []float64{10, 10, 5, 15, 5, 15, 5, 10, 15}

	if moyenne(notes) != 10 {
		t.Fail()
	}
}

func TestMoyenne4(t *testing.T) {
	notes := []float64{10, 10, 20, 15, 20, 15, 20, 10, 15}

	if moyenne(notes) != 15 {
		t.Fail()
	}
}
	
