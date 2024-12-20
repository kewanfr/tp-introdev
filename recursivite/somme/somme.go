package somme

/*
On souhaite calculer la somme des valeurs contenues dans un tableau d'entiers. C'est le rôle de la fonction somme.

Pour cet exercice les boucles for sont interdites.

# Entrée
- t : un tableau d'entiers

# Sortie
- s : la somme des éléments contenus dans le tableau

# Info
2023-2024, test 2, exercice 2
*/

func somme(t []int) (s int) {

	if len(t) == 0 {
		return 0
	} else if len(t) == 1 {
		return t[0]
	}

	return somme(t[:len(t) / 2]) + somme(t[len(t) / 2:])
}
