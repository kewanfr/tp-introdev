package suitecond

/*
On considère la suite définie par
- U(n) = 2*U(n-1) + 1 si U(n-1) < 100
- U(n) = U(n-1) - 100 si U(n-1) >= 100
- U(0) = 2.
La fonction terme doit permettre d'obtenir le n-ième terme de cette suite.

Pour cet exercice, les boucles for sont interdites

# Entrée
- n, un entier

# Sortie
- un, le terme U(n) de la suite

# Info
2022-2023, test 2, exercice 1
*/

// RÉSOLU - 11/10/2024

func terme(n uint) (un int) {
	if n == 0 {
		return 2
	}

	var un_1 int = terme(n-1)

	if un_1 >= 100 {
		un = un_1 - 100
	} else {
		un = 2*un_1 + 1
	}

	return un
}
