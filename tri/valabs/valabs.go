package valabs

/*
La fonction valabs doit trier un tableau d'entiers de la plus petite valeur absolue
à la plus grande valeur absolue. En cas
d'égalité de valeur absolue, les nombres
négatifs doivent être placés avant les
nombres positifs.

# Entrée
- tab : un tableau d'entiers
*/

func valabs(tab []int) {

	var i, j int
	for i = 1; i < len(tab); i ++ {
		val := tab[i]
		for j = i; j > 0 && estPlusGrand(tab[j-1], val); j-- {
			tab[j] = tab[j-1]
		}

		tab[j] = val
	}
	
}

func estPlusGrand(a, b int) bool {
	if abs(a) == abs(b) {
		return a > b
	}
	return abs(a) > abs(b)
}


func abs(val int) int{
	if val < 0 {
		return (-1)*val
	}
	return val
}