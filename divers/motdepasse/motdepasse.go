package motdepasse

/*
On considère qu'un mot de passe est assez solide s'il contient au moins une lettre minuscule, une lettre majuscule, un chiffre et si, de plus, il fait au moins 8 caractères. La fonction estSolide doit indiquer si un mot de passe est solide ou pas.

#Entrée
- mdp : le mot de passe à tester

# Sortie
- solide : true si mdp est un mot de passe assez solide, false sinon

# Info
2023-2024, test 3, exercice 1
*/

// RÉSOLU - 02/10/2024

var minuscules string = "abcdefghijklmnopqrstuvwxyz"
var majuscules string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
var chiffres string = "0123456789"

func estSolide(mdp string) (solide bool) {

	var uneLettreMin bool = false
	var uneLettreMaj bool = false
	var unChiffre bool = false

	for i := 0; i < len(mdp); i ++ {
		for j := 0; j < len(majuscules); j ++ {
			if string(mdp[i]) == string(majuscules[j]) {
				uneLettreMaj = true
			}
			if string(mdp[i]) == string(minuscules[j]) {
				uneLettreMin = true
			}
		}


		for l := 0; l < len(chiffres); l ++ {
			if string(mdp[i]) == string(chiffres[l]) {
				unChiffre = true
			}
		}
	}

	solide = uneLettreMaj && uneLettreMin && unChiffre && len(mdp) >= 8

	return solide
}
