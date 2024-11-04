package egalite

/*
On considère des ensembles de nombres représentés par des tableaux : on considère
que ces tableaux ne contiennent qu'une seule fois chaque nombre (puisqu'ils
représentent des ensembles) et les nombres ne sont pas nécessairement dans
l'ordre dans les tableaux.

On veut savoir si deux ensembles sont égaux ou pas, c'est-a-dire savoir si deux
tableaux contiennent les mêmes nombres ou pas. C'est à cette question que doit
répondre la fonction egalite.

# Entrées
- t1 : un tableau d'entiers (sans doublons) représentant un ensemble
- t2 : un tableau d'entiers (sans doublons) représentant un ensemble

# Sortie
- egaux : un booléen qui vaut true si t1 et t2 représentent le même ensemble et
          qui vaut false sinon

# Info
2021-2022, test 2, exercice 3
2023-2024, test 2, exercice 6
2024-2025, test 1, exercice 6
*/

// RÉSOLU - 11/10/2024

func egalite(t1, t2 []int) (egaux bool) {
	egaux = true
	for i := 0; i < len(t1); i++ {
		var est_present bool = false
		for j := 0; j < len(t2); j++ {
			if t2[j] == t1[i]{
				est_present = true
			}
		}
		if !est_present {
			egaux = false
		}
		est_present = false
	}

	if len(t2) > len(t1){
		egaux = false
	}

	return egaux
}
