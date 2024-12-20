package divisible

/*
La fonction estDivisible doit indiquer, étant donnés deux entiers positifs ou nuls, si l'un est divisible par l'autre ou pas.

# Entrées
- n : un entier positif ou nul
- m : un entier positif ou nul

# Sortie
- reponse : un booléen qui vaut true si n est divisible par m ou si m est divisible par n et false sinon

# Info
2022-2023, test 1, exercice 4
*/

// RÉSOLU - 15/10/2024
// TEST 1 - 2022-2023

func estDivisible(n, m uint) (reponse bool) {

	if m > n {
		reponse = m % n == 0
	}else {
		reponse = n % m == 0
	}

	return reponse
}
