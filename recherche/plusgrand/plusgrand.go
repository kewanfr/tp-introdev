package plusgrand

/*
Étant donné un tableau d'entiers, on souhaite trouver le plus grand entier présent dans ce tableau. C'est le rôle de la fonction plusgrand.

# Entrée
- t : un tableau d'entiers de longueur supérieure ou égale à 1

# Sortie
- x : le plus grand entier présent dans t

# Info
2024-2025, test 1, exercice 3
*/

func plusgrand(t []int) (x int) {

	x = t[0]

	for i := 1; i < len(t); i ++ {
		if t[i] > x{
			x = t[i]
		}
	}

	return x
}
