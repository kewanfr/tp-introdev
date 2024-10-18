package tri

/*
La fonction tri doit trier un tableau d'entiers du plus petit au plus grand.
Cette fonction ne doit pas modifier le tableau donné en entrée.

# Entrée
- tinit : un tableau d'entiers qui ne doit pas être modifié.

# Sortie
- tfin : un tableau contenant les mêmes entiers que tinit mais triés du plus
         petit au plus grand.

# Info
2021-2022, test2, exercice 6
*/

func tri(tinit []int) (tfin []int) {

	tfin = append(tfin, tinit...)

	// Tri par insertion

	for i := 0; i < len(tfin); i ++ {
		var v = tfin[i]
		var j = i - 1

		for j >= 0 && tfin[j] > v {
			tfin[j + 1] = tfin[j]
			j--
		}

		tfin[j + 1] = v
	}

	return tfin
}
