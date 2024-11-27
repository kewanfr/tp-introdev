package classement

/*
Étant donné un tableau d'entiers, la fonction classer doit les ranger dans trois tableaux : ceux qui sont pairs, ceux qui sont impairs et plus petits que 100, ceux qui sont impairs et plus grands que 100.

# Entrée
- t : un tableau d'entiers

# Sorties
- pairs : tous les entiers pairs de t
- petits : tous les entiers impairs plus petits que 100 de t
- grands : tous les entiers impairs plus grands que 100 de t

# Exemple
classer([]int{99, 100, 101}) = [100], [99], [101]
*/
func classer(t []int) (pairs, petits, grands []int) {

	for i := 0; i < len(t); i ++ {
		v := t[i]
		if v % 2 == 0 {
			pairs = append(pairs, v)
		} else if v < 100 {
			petits = append(petits, v)
		} else {
			grands = append(grands, v)
		}
	}

	return pairs, petits, grands
}
