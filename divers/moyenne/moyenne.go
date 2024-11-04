package moyenne

/*
Étant donné un ensemble de notes, on souhaite calculer leur moyenne. C'est le rôle de la fonction moyenne.

# Entrée
- t : un tableau (non vide) de notes

# Sortie
- m : la moyenne des notes du tableau

# Info
2024-2025, test 1, exercice 2
*/

func moyenne(t []float64) (m float64) {

	var sum float64 = 0

	for i := 0; i < len(t); i ++ {
		sum += t[i]
	}

	return sum/float64(len(t))
}
