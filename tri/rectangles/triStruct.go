package triStruct

/*
On souhaite trie un ensemble de rectangle en fonction de leur surface : de la plus petite surface à la plus grande surface. Les rectangles sont représentés par une structure rectangle qui comprend un champs largeur et un champs longueur (on rappelle que la surface d'un rectangle est le produit de sa largeur par sa longueur). C'est la fonction ranger qui doit effectuer ce tri.

# Entrée
- t : un tableau de rectangles

# Sortie
- enOrdre : un tableau de rectangles ayant exactement les mêmes éléments que t mais rangés de celui qui a la plus petite surface à celui qui a la plus grande surface. L'ordre relatif des rectangles ayant la même surface n'a pas d'importance.

# Info
2024-2025, test 2, exercice 4
*/

type rectangle struct {
	largeur  int
	longueur int
}

func estPlusGrand(a, b rectangle) bool {
	return a.largeur*a.longueur > b.largeur*b.longueur
}

func ranger(t []rectangle) (enOrdre []rectangle) {

	enOrdre = append(enOrdre, t...)

	var i, j int
	for i = 0; i < len(enOrdre); i ++{
		v := enOrdre[i]

		for j = i - 1; j >= 0 && estPlusGrand(enOrdre[j], v); j -- {
			enOrdre[j + 1] = enOrdre[j]
		}

		enOrdre[j + 1] = v
	}

	return enOrdre
}
