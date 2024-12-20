package alphanum

/*
On appelle caratères alphanumériques les lettres minuscules, les lettres majuscules et les chiffres. On souhaite vérifier qu'une chaîne de caractères est constituée uniquement de caractères alphanumériques. C'est le travail de la fonction alphanum.

# Entrée
- s : une chaîne de caractères

# Sortie
- isAlphanum : un booléen qui vaut true si s est entièrement constituée de caractères alphanumériques et qui vaut false sinon

# Info
2023-2024, test 2, exercice 3
*/

// RÉSOLU - 02/10/2024

var alphaList []rune = []rune{
	'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j',
	'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't',
	'u', 'v', 'w', 'x', 'y', 'z',
	'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J',
	'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T',
	'U', 'V', 'W', 'X', 'Y', 'Z',
}

var numList []rune = []rune{
	'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
}


func alphanum(s string) (isAlphanum bool) {
	isAlphanum = true
	
	for i := 0; i < len(s); i ++ {

		var isAlpha bool = false
		for j := 0; j < len(alphaList); j ++ {
			if string(s[i]) == string(alphaList[j]) {
				isAlpha = true
			}
		}

		var isNum bool = false
		for j := 0; j < len(numList); j ++ {
			if string(s[i]) == string(numList[j]) {
				isAlpha = true
			}
		}
		if !isNum && !isAlpha {
			isAlphanum = false
		}

	}

	return isAlphanum
}
