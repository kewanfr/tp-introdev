package doublons2

/*
On dispose d'un tableau d'entiers de longueur n et on suppose qu'il contient
exactement une fois chaque nombre compris entre 1 et n. On voudrait vérifier
cela. C'est le travail de la fonction doublons.

Coder la fonction doublons de manière à ne parcourir le tableau tab qu'une seule
fois rapportera plus de points.

# Entrée
- tab : un tableau d'entiers

# Sortie
- ok : un booléen qui vaut true si tab contient bien exactement une
fois chaque entier entre 1 et len(tab) et false sinon

# Info
2021-2022, test 1, exercice 8
*/

func doublons(tab []int) (ok bool) {
	var verif map[int]int = make(map[int]int)

	ok = true
	for i := 0; i < len(tab); i++ {

		val := tab[i]

		_, existVerif := verif[val]
		if existVerif {
			verif[val] += 1
			ok = false
		}else {
			verif[val] = 1
		}

		if val > len(tab) || val < 1{
			ok = false
		}

	}

	return ok
}
