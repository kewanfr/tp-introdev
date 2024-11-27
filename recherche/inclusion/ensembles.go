package ensembles

/*
	On souhaite comparer des ensembles de nombres qui seront représentés par des tableaux. On dira qu'un ensemble A est inclus dans un ensemble B si tous les nombres de A sont présents dans B (on rappelle que, dans un ensemble, un nombre ne peut pas apparaître plusieurs fois). Étant donnés deux ensembles, la fonction inclus vérifiera si le premier est inclus dans le deuxième.

	# Entrées
	- t1 : un tableau de nombres représentant un ensemble (donc, ne contenant aucun nombre en double)
	- t2 : un tableau de nombres représentant un ensemble (donc, ne contenant aucun nombre en double)

	# Sortie
	- estInclus : un booléen qui vaut true t1 est inclus dans t2 et false sinon

# Info
2024-2025, test 1, exercice 5
*/

func inclus(t1, t2 []int) (estInclus bool) {

	estInclus = true

	for i := 0; i < len(t1); i ++ {
		nbInclus := 0
		for j := 0; j < len(t2); j++ {
			if t2[j] == t1[i] {
				nbInclus++
			}
		}
		if nbInclus != 1{
			estInclus = false
		}
	}

	return estInclus
}
