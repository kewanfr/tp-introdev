package bougies

/*
La fonction bougies doit indiquer, pour une date donnée, combien de bougies il y avait sur le dernier gateau d'anniversaire d'une personne née le 17 avril 1986. Si jamais la date en question est un 17 avril, on donnera le nombre de bougies sur le gateau à cette date. Avant son premier anniversaire, une personne n'a jamais eu de gateau d'anniversaire et, donc, a toujours eu 0 bougies sur ses gateaux d'anniversaire.

# Entrées
- année : un entier représentant l'année
- mois : un entier représentant le numéro d'un mois dans l'année
- jour : un entier représentant le numéro d'un jour dans le mois

# Sortie
- numBougies : le nombre de bougies sur le dernier gateau d'anniversaire d'une personne née le 17 avril 1986 si on est à la date définie par jour, mois et année

# Remarque
On considèrera que la date donnée en entrée est toujours valide (pas de jour 35 ou de mois 52)

# Exemples
bougies(23, 5, 1990) = 4
bougies(17, 4, 1990) = 4
bougies(25, 2, 1990) = 3

# Info
2022-2023, test 1, exercice 5
*/

// RÉSOLU - 15/10/2024
// TEST 1 - 2022-2023

func bougies(jour, mois, annee uint) (numBougies uint) {

	numBougies = annee - 1986

	if (mois < 4 || (mois == 4 && jour < 17)) {
		numBougies -= 1
	}

	if annee < 1986 || (annee == 1986 && mois < 4) || (annee == 1986 && mois == 4 && jour < 17){
		numBougies = 0
	}

	return numBougies
}
