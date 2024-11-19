package alphabetique

/*
On souhaite classer un ensemble de mots par ordre alphabétique. Les mots sont donnés dans un tableau et doivent être triés en place dans ce même tableau. C'est le rôle de la fonction alphabetique.

On rappel que la comparaison "<" peut être utilisées sur des chaînes de caractères pour savoir la quelle est la première en ordre alphabétique.

# Entrée
- dico : un tableau de chaînes de caractères qui doit être trié en place en ordre alphabétique

# Info
2023-2024, test 3, exercice 2
*/

func alphabetique(dico []string) {

	var i, j int
	for i = 0; i < len(dico); i ++ {
		val := dico[i]

		for j = i; j > 0 && dico[j - 1] > val; j -- {
			dico[j] = dico[j - 1]
		}

		dico[j] = val
	}

}
