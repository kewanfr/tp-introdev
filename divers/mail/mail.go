package mail

/*
Une adresse mail valide est une chaîne de caractères de la forme x@y.z avec x, y et z trois chaînes de caractères ayant les caractéristiques suivantes :
- x est non-vide et ne contient pas le caractère "@"
- y est non-vide et ne contient pas le caractère "@" ni le caratère "."
- z contient exactement deux caractères qui sont des lettres minuscules non accentuées
La fonction estAdresseValide doit vérifier si une chaîne de caractères est bien une adresse mail valide.

# Entrée
- s : une chaîne de caractères

# Sortie
- ok : un booléen qui vaut "true" si s est une adresse mail valide et "false" sinon

# Info
2024-2025, test 2, exercice 5
*/

const minuscules string = "abcdefghijklmnopqrstuvwxyz"
const arobase byte = '@'
const point byte = '.'

func estAdresseValide(s string) (ok bool) {
	return
}
