package positifs

/*
La fonction positifs doit indiquer combien il y a de nombres strictement positifs dans un tableau donné en paramètre. On peut remarquer que le nombre de nombres strictement positifs dans un tableau t est la somme du nombre de nombres strictement positifs dans la première moitié de t et du nombre de nombres strictement positifs dans la deuxième moitié de t.

Vous ne devez pas utiliser de boucles for dans cet exercice.

# Entrée
- t, le tableau dans lequel compter les nombres strictement positifs

# Sortie
- num, le nombre de nombres strictement positifs dans t

# Info
2022-2023, test 4, exercice 5
*/

func positifs(t []int) (num int) {

	if len(t) == 0 {
		return 0
	} else if len(t) == 1 {
		if t[0] > 0 {
			return 1
		} else {
			return 0
		}
	}

	return positifs(t[:len(t) / 2]) + positifs(t[len(t)/2:])
}
