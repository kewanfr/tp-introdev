package tri

/*
On souhaite trier un ensemble de chaînes de caractères selon leur longueur : la plus courte en premier, la plus longue en dernier. C'est le rôle de la fonction classer.

# Entrée
- t : un tableau de chaînes de caractères

# Sortie
- enOrdre : un tableau de chaînes de caractères qui contient les mêmes chaînes que t mais triées en ordre croissant de leur longueur. Si deux chaînes ont la même longueur leur ordre relatif n'a pas d'importance.

# Info
2024-2025, test 2, exercice 3
*/

func classer(t []string) (enOrdre []string) {

	enOrdre = append(enOrdre, t...)

	var i, j int
	for i = 0; i < len(enOrdre); i ++ {
		v := enOrdre[i]

		for j = i - 1; j >= 0 && len(enOrdre[j]) > len(v); j -- {
			enOrdre[j + 1] = enOrdre[j]
		}
		enOrdre[j + 1] = v

	}

	return enOrdre
}
