package decroissant

/*
La fonction decroissant doit trier un tableau d'entiers dans l'ordre décroissant.

# Entrée
- t, le tableau à trier (en place)

# Info
2022-2023, test 4, exercice 7
*/

func decroissant(tab []int) {
	var i, j int

	for i = 1; i < len(tab); i ++ {
		val := tab[i]

		for j = i; j > 0 && tab[j - 1] < val; j -- {
			tab[j] = tab[j - 1]
		}

		tab[j] = val
	}
}
