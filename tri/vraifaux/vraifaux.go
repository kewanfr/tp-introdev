package vraifaux

/*
On souhaite trier un tableau de booléens en mettant ceux dont la valeur est true au début et ceux dont la valeur est false à la fin. La fonction trier doit réaliser ce tri en place.

# Entrée
- tab : le tableau de booléens à trier

# Info
2022-2023, test3, exercice 0
*/

func trier(tab []bool) {
	
	var i, j int
	for i = 1; i < len(tab); i ++ {
		val := tab[i]

		for j = i; j > 0 && tab[j - 1] == false; j -- {
			tab[j] = tab[j-1]
		}
		
		tab[j] = val
	}
}
