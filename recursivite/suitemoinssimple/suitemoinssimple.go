package suiteMoinsSimple

/*
On considère la suite dont le terme u(n) est défini par
- u(n) = 3 * u(n-1) + 3 si u(n-1) est divisible par 2
- u(n) = u(n-1) - 1 si u(n-1) n'est pas divisible par 2 et
- u(0)=1
La fonction terme calcule les termes de cette suite.

# Entrée
- n : un entier indiquant le numéro du terme à calculer

# Sortie
- un : le terme u(n) de la suite

# Remarque
Les boucles ne sont pas autorisées pour résoudre cet exercice

# Info
2022-2023, test 1, exercice 7
*/

// RÉSOLU - 15/10/2024
// TEST 1 - 2022-2023

func terme(n int) (un int) {

	// Cas de base
	if n == 0 {
		return 1
	}

	var uNm1 int = terme(n - 1) // Récup de u(n-1)

	// Si u(n-1) est divisible par 2
	if uNm1 % 2 == 0 {
		return 3*uNm1 + 3
	}

	// Sinon
	return uNm1 - 1
	
}
