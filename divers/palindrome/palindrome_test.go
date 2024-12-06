package palindrome

import (
	"testing"
)

func TestBanane(t *testing.T) {
	if sousChaine("banane") != "ana" {
		t.Fail()
	}
}

func TestKayak(t *testing.T) {
	if sousChaine("kayak") != "kayak" {
		t.Fail()
	}
}

// Ajoutés après le DS

func TestVide(t *testing.T) {
	if sousChaine("") != "" {
		t.Fail()
	}
}

func TestPetit(t *testing.T) {
	if sousChaine("a") != "a" {
		t.Fail()
	}
}

func TestNuage(t *testing.T) {
	if sousChaine("grandnuage") != "ndn" {
		t.Fail()
	}
}

func TestLong(t *testing.T) {
	if sousChaine("bananekayakkayakkayakgrandnuage") != "kayakkayakkayak" {
		t.Fail()
	}
}
