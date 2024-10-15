package classement

/*
La fonction classer doit trier un tableau de booléen en séparant d'un côté les booléens qui valent true et de l'autre les booléens qui valent false

# Entrée
- t : le tableau de booléens à trier

# Sortie
- tri : un tableau de booléens qui contient exactement les mêmes éléments que t mais avec tous les booléens qui valent true au début et tous les booléens qui valent false à la fin

# Info
- 2023-2024, test 1, exercice 4
*/

func classer(t []bool) (tri []bool) {

	var nbTrue int = 0
	var nbFalse int = 0

	for i := 0; i < len(t); i ++ {
		if t[i] {
			nbTrue++
		} else {
			nbFalse++
		}
	}

	for i := 0; i < nbTrue; i ++ {
		tri = append(tri, true)
	}

	for i := 0; i < nbFalse; i ++ {
		tri = append(tri, false)
	}

	return tri
}
