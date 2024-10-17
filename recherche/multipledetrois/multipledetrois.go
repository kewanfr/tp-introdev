package multipleDeTrois

/*
Étant donné un tableau d'entiers positifs, la fonction unMultipleDeTrois retourne un multiple de trois présent dans ce tableau ou -1 s'il n'en existe pas.

# Entrée
- t : un tableau d'entiers positifs

# Sortie
- multiple : un entier de t qui est multiple de 3 ou -1 s'il n'en existe pas

# Info
2022-2023, test 1, exercice 2
*/

// RÉSOLU - 15/10/2024
// TEST 1 - 2022-2023

func unMultipleDeTrois(t []int) (multiple int) {

	multiple = -1

	for i :=0; i < len(t); i++ {
		if t[i] % 3 == 0 {
			multiple = t[i]
		}
	}

	return multiple
}
