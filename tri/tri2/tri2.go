package tri2

/*
La fonction triabs doit trier un tableau d'entiers de la plus grande valeure
absolue à la plus petite valeure absolue. Cette fonction ne doit pas modifier
le tableau donné en entrée.

# Entrée
- tinit : un tableau d'entiers qui ne doit pas être modifié.

# Sortie
- tfin : un tableau contenant les mêmes entiers que tinit mais triés du plus
         grand (en valeure absolue) au plus petit (en valeure absolue).

# Info
2021-2022, test2, exercice 7
*/

func triabs(tinit []int) (tfin []int) {
	tfin = append(tfin, tinit...)

	var i, j int
	for i = 1; i < len(tfin); i ++ {
		v := tfin[i]
		for j = i; j > 0 && abs(tfin[j - 1]) < abs(v); j-- {
			tfin[j] = tfin[j-1]
		}
		tfin[j] = v

	}


	return tfin
}

func abs(x int) int {
	if x < 0 {
		return (-1)*x
	}
	return x
}