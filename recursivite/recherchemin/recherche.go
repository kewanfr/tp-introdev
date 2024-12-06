package recherche

/*
Étant donné un ensemble d'entiers, on souhaite connaître la valeur du plus petit d'entre-eux. C'est le rôle de la fonction trouvePlusPetit.

Attention, pour cet exercice les boucles for sont interdites.

# Entrée
- t : un tableau d'entier (qui contient au moins un élément)

# Sortie
- plusPetit : le plus petit entier présent dans t

# Info
2024-2025, test 2, exercice 6
*/

func trouvePlusPetit(t []int) (plusPetit int) {

	if len(t) == 1 {
		return t[0]
	}
	// if len(t) == 2 {
	// 	if t[0] < t[1]{
	// 		return t[0]
	// 	}
	// 	return t[1]
	// }

	p1 := trouvePlusPetit(t[:len(t)/2])
	p2 := trouvePlusPetit(t[len(t)/2:])

	if p1 < p2{
		return p1
	}
	return p2
}
