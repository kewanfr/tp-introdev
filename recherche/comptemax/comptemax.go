package comptemax

/*
La fonction compteMax doit trouver le plus grand entier dans un tableau et indiquer combien de fois cet entier apparaît dans le tableau.

# Entrée
- t : un tableau d'entiers

# Sortie
- val : le plus grand entier dans t
- nombre : le nombre de fois que val apparaît dans t

# Info
- 2023-2024, test 1, exercice 7
*/

// RÉSOLU - 17/10/2024
// TEST 1 - 2023-2024

func compteMax(t []int) (val, nombre int) {

	var max int = 0
	nombre = 0

	for i := 0; i < len(t); i++ {
		if t[i] > max{
			if max != t[i] {
				nombre = 0
			}
			max = t[i]
		}

		if t[i] == max {
			nombre += 1
		}
	}

	return max, nombre
}
