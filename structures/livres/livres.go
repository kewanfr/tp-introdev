package livres

/*
On dispose d'une structure de données permettant de réprésenter des livres avec un titre, un prix et un nombre de page.

On souhaite trier un tableau de livres soit par prix (du moins cher au plus cher) soit par nombre de pages (du plus grand nombre de pages au plus petit nombre de pages). Dans les deux cas, si deux livres ne sont pas distinguables, on les triera par ordre alphabétique de leur titre (en utilisant l'opérateur < sur les chaînes de caractères). La fonction trier doit réaliser ce tri en place.

# Entrée
- tab : le tableau de livres à trier
- parPrix : un booléen indiquant le tri à réaliser, s'il vaut true on triera les livres du moins cher au plus cher, s'il vaut false on triera les livres de celui qui a le plus de pages à celui qui en a le moins.

# Info
2022-2023, test3, exercice 4
*/

type livre struct {
	titre    string
	prix     float64
	numPages int
}

func obtenirFonction(parPrix bool) func(a, b livre) bool {
	if parPrix {
		return func(a, b livre) bool { return a.prix > b.prix}
	}
	return func(a, b livre) bool { return a.numPages < b.numPages }
}

func trier(tab []livre, parPrix bool) {

	condition := obtenirFonction(parPrix)

	var i, j int
	for i = 1; i < len(tab); i ++ {
		val := tab[i]
		for j = i; j > 0 && condition(tab[j-1], val); j-- {
			tab[j] = tab[j-1]
		}

		tab[j] = val
	}

}
