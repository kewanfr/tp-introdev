package plusQueZero

/*
La fonction tousPositifs indique si tous les nombres contenus dans un tableau sont strictement supérieurs à 0

# Entrée
- t : un tableau d'entiers

# Sortie
- reponse : un booléen qui vaut true tous les nombres de t sont strictement supérieurs à 0 et false sinon

# Info
2022-2023, test 1, exercice 3
*/

// RÉSOLU - 15/10/2024
// TEST 1 - 2022-2023

func tousPositifs(t []int) (reponse bool) {

	reponse = true

	for i := 0; i < len(t); i++ {
		if t[i] <= 0 {
			reponse = false
		}
	}

	return reponse
}
