package doublons1

/*
On dispose d'un tableau d'entiers de longueur n et on suppose qu'il contient
exactement les nombres de 1 à n dans l'ordre. On voudrait vérifier cela. C'est
le travail de la fonction doublons.

# Entrée
- tab : un tableau d'entiers

# Sortie
- ok : un booléen qui doit valoir true si tab contient exactement les entiers de
1 à len(tab) dans l'ordre et false sinon
*/

// RÉSOLU - 15/10/2024

func doublons(tab []int) (ok bool) {

	ok = true
	
	for i := 0; i < len(tab); i++ {
		if (tab[i] - 1 != i) {
			ok = false
		}
	}

	return ok
}
