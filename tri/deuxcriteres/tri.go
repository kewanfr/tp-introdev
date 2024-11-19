package tri

/*
La fonction tri doit trier un tableau d'entiers de la manière suivante : on trouvera d'abord tous les nombres strictement négatifs du tableau, dans l'ordre décroissant, puis tous les nombres positifs ou nuls, dans l'ordre croissant.

# Entrée
- t, le tableau à trier (en place)

# Info
2022-2023, test 4, exercice 8
*/

func sort(tab []int, min, max int, croissant bool) {
	var i, j int

	for i = min; i < max; i ++ {
		val := tab[i]

		operande := func(a, b int) bool {
			if croissant {
				return a > b
			} else {
				return a < b
			}
		}

		for j = i; j > min && operande(tab[j - 1], val); j -- {
			tab[j] = tab[j-1]
		}
		tab[j] = val
	}
}

func tri(t []int) {

	milieu := 0

	var i, j int
	for i = 0; i < len(t); i ++ {
		if t[i] < 0{
			milieu += 1

			val := t[i]
			for j = i; j > 0 && t[j - 1] >= 0; j --{
				t[j] = t[j - 1]
			} 

			t[j] = val

		}
	}

	sort(t, 0, milieu, false)
	sort(t, milieu, len(t), true)

}
