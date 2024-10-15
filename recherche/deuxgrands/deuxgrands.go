package deuxGrands

/*
La fonction lesDeuxPlusGrands retourne les deux plus grands entiers présents dans un tableau

# Entrée
- t : un tableau d'entiers qui en contient au moins 2

# Sorties
- v1 : le plus grand entier dans t
- v2 : le deuxième plus grand entier dans t

# Info
2022-2023, test 1, exercice 8
*/

func lesDeuxPlusGrands(t []int) (v1, v2 int) {


	// On prend les deux premiers et on attribue selon le plus grand et le second plus grand
	if t[0] > t[1] {

		v1 = t[0]
		v2 = t[1]
	}else {
		v1 = t[1]
		v2 = t[0]
	}

	// On parcourt à partir du 3ème élément
	for i := 2; i < len(t); i++ {
		// Si l'élément est + grand ou égal à v1, alors le second plus grand est à v1 (l'ancien plus grand) et v1 est au nouveau
		if t[i] >= v1 {
			v2 = v1
			v1 = t[i]
		}
		
	}

	return v1, v2
}
