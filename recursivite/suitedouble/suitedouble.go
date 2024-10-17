package suitedouble

/*
Soient (Un) et (Vn) les suites définies ainsi :
- U(0) = 5
- V(0) = 3
- U(n) = 3 * V(n) + 6
- V(n) = U(n-1) + 7
La fonction terme doit permettre d'obtenir le terme numéro n de la suite (Un).

Pour cet exercice, les boucles for sont interdites.

# Entrée
- n, un entier positif ou nul

# Sortie
- un, le terme U(n) de la suite (Un)

# Info
- 2023-2024, test 1, exercice 8
*/

// RÉSOLU - 15/10/2024
// TEST 1 - 2023-2024

func terme(n uint) (un int) {

	if n == 0 {
		return 5
	}

	return 3 * (terme(n-1) + 7) + 6
	
}
