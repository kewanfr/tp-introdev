package moitie

/*
La fonction moitieDePositifs indique si au moins de la moitié des nombres contenus dans un tableau sont strictement supérieurs à 0

# Entrée
- t : un tableau d'entiers

# Sortie
- reponse : un booléen qui vaut true si au moins la moitié des nombres de t sont strictement supérieurs à 0 et false sinon

# Info
2022-2023, test 1, exercice 6
*/

// RÉSOLU - 15/10/2024
// TEST 1 - 2022-2023

func moitieDePositifs(t []int) (reponse bool) {

	var nbSup0 float32 = 0

	for i := 0; i < len(t); i++ {
		if t[i] > 0 {
			nbSup0 += 1
		}
	}

	reponse = nbSup0 >= (float32(len(t)) / float32(2))

	return reponse
}
