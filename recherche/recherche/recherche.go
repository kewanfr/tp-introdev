package recherche

/*
La fonction recherche doit indiquer la position d'une valeur v dans un tableau
tab. Si la valeur v n'est pas présente, il faut l'indiquer en retournant false.

# Entrées
- a : la valeur à chercher
- t : le tableau dans lequel chercher

# Sorties
- pos : la position de v dans tab (si elle existe)
- existe : un booléen indiquant si v est présente dans tab ou pas

# Info
2021-2022, test 1, exercice 5
2023-2024, test 1, exercice 0
*/

// RÉSOLU - 15/10/2024
// TEST 1 - 2021-2022

func trouve(a int, t []int) (pos int, existe bool) {

	pos = -1
	existe = false

	for i := 0; i < len(t); i++ {
		if t[i] == a {
			existe = true
			pos = i
		}
	}

	return pos, existe
}
