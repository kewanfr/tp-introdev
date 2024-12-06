package tri

/*
On dispose d'un ensemble de rectangle, chacun portant un nom (représentés par une structure "rectangle"). On souhaite trier ces rectangles de la manière suivante :
- de la plus petite surface à la plus grande surface (on rappelle que la surface d'un rectangle est le produit de sa largeur par sa longueur)
- en cas d'égalité de surface, du nom le plus court au nom le plus long
- en cas d'égalité de surface et de longueur de nom, en ordre alphabétique des noms
C'est le rôle de la fonction ranger de réaliser ce tri.

# Entrée
- t : un tableau de rectangles

# Sortie
- enOrdre : un tableau contenant les mêmes rectangles mais trié selon la méthode décrite ci-dessus

# Info
2024-2025, test 2, exercice 7
*/

type rectangle struct {
	nom      string
	largeur  int
	longueur int
}

func estPlusGrand(a, b rectangle) bool{
	surfaceA := a.largeur*a.longueur
	surfaceB := b.largeur*b.longueur

	if surfaceA == surfaceB {
		if len(a.nom) == len(b.nom) {
			return a.nom > b.nom
		}

		return len(a.nom) > len(b.nom)
	}
	return surfaceA > surfaceB
}

func ranger(t []rectangle) (enOrdre []rectangle) {

	enOrdre = append(enOrdre, t...)

	var i, j int
	for i = 0; i < len(enOrdre); i ++ {
		v := enOrdre[i]

		for j = i - 1; j >= 0 && estPlusGrand(enOrdre[j], v); j -- {
			enOrdre[j + 1] = enOrdre[j]
		}

		enOrdre[j + 1] = v
	}

	return enOrdre
}
