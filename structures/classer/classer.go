package classer

/*
La fonction classer doit trier un tableau d'étudiants (tels que définis) ci-dessous
de celui qui a la meilleur moyenne (la plus élevée) à celui qui a la plus mauvaise
moyenne. En cas d'égalité de moyenne entre deux étudiants, celui qui a le nom de
famille qui arrive en premier dans l'ordre alphabétique doit être classer avant
l'autre (on peut utiliser <, >, <=, >=, == pour comparer les chaînes de caractères
par ordre alphabétique). Si deux étudiants ont la même moyenne et le même nom, on
met en premier celui qui a le prénom qui est en premier dans l'ordre alphabétique.

# Entrée
- promo : le tableau d'étudiants à trier

# Sortie
- classement : un tableau contenant les mêmes étudiants mais trié

# Info
2021-2022, test3, exercice 5
*/

type etudiant struct {
	nom, prenom string
	moyenne     float64
}

func aAprèsB(a, b etudiant) bool {
	if a.moyenne == b.moyenne {

		if a.nom == b.nom{
			return a.prenom > b.prenom
		}

		return a.nom > b.nom

	}
	return a.moyenne < b.moyenne
}

func classer(promo []etudiant) (classement []etudiant) {

	classement = append(classement, promo...)

	var i, j int
	for i = 0; i < len(classement); i ++{
		v := classement[i]
		
		for j = i - 1; j >= 0 && aAprèsB(classement[j], v); j -- {
			classement[j + 1] = classement[j]
		}

		classement[j + 1] = v
	}


	return classement
}
