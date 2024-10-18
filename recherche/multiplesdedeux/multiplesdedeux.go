package multiplesdedeux

/*
La fonction sontMultiples doit dire si oui ou non tous les nombres contenus dans un tableau sont des multiples de 2.

# Entrée
- t : un tableau d'entiers

# Sortie
- ok : un booléen qui vaut true si tous les entiers de t sont des multiples de 2 et false sinon

# Info
- 2023-2024, test 1, exercice 3
*/

// RÉSOLU - 17/10/2024
// TEST 1 - 2023-2024

func sontMultiples(t []int) (ok bool) {

	ok = true

	for i := 0; i < len(t); i ++ {
		if t[i] % 2 != 0 {
			ok = false
		}
	}

	return ok
}
