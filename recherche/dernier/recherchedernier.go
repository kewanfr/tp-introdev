package recherchedernier

/*
Étant donné un tableau d'entiers, la fonction trouve doit dire si un entier donné appartient à ce tableau et, si c'est le cas, doit indiquer la position où trouver la dernière apparition de cet entier dans le tableau.

# Entrées
- a : un entier
- t : un tableau d'entier

# Sorties
- pos : la position la plus grande où on trouve l'entier a dans le tableau t, si une telle position existe, ou une valeur quelconque sinon
- existe : un booléen qui vaut true si l'entier a existe dans le tableau t et qui vaut false sinon

# Info
- 2023-2024, test 1, exercice 6
*/

// RÉSOLU - 17/10/2024
// TEST 1 - 2023-2024

func trouveDernier(a int, t []int) (pos int, existe bool) {

	existe = false

	for i :=0; i < len(t); i ++ {
		if t[i] == a {
			pos = i
			existe = true
		}
	}

	return pos, existe
}
