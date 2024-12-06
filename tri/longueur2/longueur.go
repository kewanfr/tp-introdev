package longueur

/*
On souhaite trier un tableau de chaînes de caractères (minuscules sans accents) de la plus longue à la plus courte (et par ordre alphabétique dans le cas où plusieurs chaînes ont la même longueur). La fonction trier doit réaliser ce tri en place. On rappelle que l'opérateur < sur les chaînes de caractères définit l'ordre alphabétique.

# Entrée
- tab : le tableau de chaînes de caractères à trier

# Info
2022-2023, test3, exercice 1
*/

func trier(tab []string) {
	var i, j int

	for i = 1; i < len(tab); i ++ {
		val := tab[i]

		for j = i; j > 0 && ((len(tab[j - 1]) == len(val) && tab[j - 1] > val) || len(tab[j - 1]) < len(val)); j -- {
			tab[j] = tab[j-1]
		}

		tab[j] = val
	}
}
