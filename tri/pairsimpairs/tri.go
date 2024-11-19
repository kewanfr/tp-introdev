package tri

/*
La fonction tri doit trier un tableau d'entiers de la manière suivante : on trouvera d'abord tous les nombres pairs du tableau, dans l'ordre croissant, puis tous les nombres impairs, dans l'ordre décroissant.

# Entrée
- t, le tableau à trier (en place)

# Info
2022-2023, test 2, exercice 4
*/

func sort(tab []int, min, max int){
	var k, l int

	for k = min; k < max; k++ {
		val := tab[k]
		for l = k; l > min && ((min == 0 && tab[l-1] > val) || (max == len(tab) && tab[l-1] < val)); l-- {
			tab[l] = tab[l-1]
		}

		tab[l] = val
	}
}

func tri(t []int) {

	// var tpairs []int
	// var timpairs []int

	milieu := 0
	var i, j int
	for i = 1; i < len(t); i ++ {
		if t[i] % 2 == 0 {
			// val := t[i]
			milieu += 1
			for j = i; j > 0 && t[j-1] % 2 != 0; j-- {
				temp := t[j]
				t[j] = t[j-1]
				t[j - 1] = temp
			}
		}
	}

	sort(t, 0, milieu)
	sort(t, milieu, len(t))
}
