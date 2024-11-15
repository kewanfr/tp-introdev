package doublons6

/*
On dispose d'un tableau d'entiers de longueur n et on suppose qu'il contient exactement les nombres de k à k + n - 1 dans l'ordre pour un certain k non connu à l'avance. On voudrait vérifier que c'est bien le cas. C'est le travail de la fonction doublons.

# Entrée
- tab : un tableau d'entiers

# Sortie
- ok : un booléen qui doit valoir true si tab contient exactement les entiers de k à k + len(tab) - 1 dans l'ordre et qui vaut false sinon

# Info
- 2023-2024, test 2, exercice 0
*/

func doublons(tab []int) (ok bool) {
	min := tab[0]
	max := tab[0]

	for i := 0; i < len(tab); i ++ {
		if tab[i] <= min{
			min = tab[i]
		}

		if tab[i] >= max {
			max = tab[i]
		}
	}

	ok = tab[0] == min
	prec := tab[0]
	for j := 1; j < len(tab); j ++ {
		if tab[j] != prec + 1 {
			ok = false
		}
		prec = tab[j]
	}


	return ok
}
