package factorielle

/*
La fonction factorielle calcule la factorielle d'un nombre entier positif ou
nul. Pour rappel : 0! = 1 et pour n > 0, n! = n*(n-1)*(n-2)*...*1.

Vous ne devez pas utiliser de boucles, la fonction factorielle sera donc récursive

# Entrée
- n : l'entier dont on veut calculer la factorielle

# Sortie
- fact : la factorielle de n

# Exemple
factorielle(5) = 120

# Info
2021-2022, test3, exercice 1
*/

// RÉSOLU - 11/10/2024

func factorielle(n uint) (fact uint) {

	if n == 0 {
		return 1
	}

	fact =  n * factorielle(n-1)
	
	return fact
}
