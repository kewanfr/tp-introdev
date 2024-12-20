package interval

/*
La fonction dansInterval doit indiquer la position dans un tableau d'un flottant situé dans un interval spécifié par ses bornes.

# Entrées
- inf : la borne inférieure d'un interval
- sup : la borne supérieure d'un interval
- t : un tableau de flottants

# Sorties
- pos : la position dans t d'un nombre strictement plus grand que inf et strictement plus petit que sup s'il en existe, on un nombre quelconque sinon
- existe : un booléen qui vaut true s'il existe un nombre strictement plus grand que inf et strictement plus petit que sup dans t et qui vaut false s'il n'en existe pas

# Info
- 2023-2024, test 1, exercice 5
*/

// RÉSOLU - 17/10/2024
// TEST 1 - 2023-2024

func dansInterval(inf, sup float64, t []float64) (pos int, existe bool) {

	existe = false
	pos = -1

	for i :=0; i < len(t); i++ {
		if inf < t[i] && t[i] < sup {
			pos = i
			existe = true
		}
	}

	return pos, existe
}
