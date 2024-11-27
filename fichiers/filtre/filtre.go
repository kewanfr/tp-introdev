package filtre

import (
	"bufio"
	"errors"
	"os"
)

/*
Étant donné un nom de fichier, on souhaite obtenir une copie de ce fichier dont les lignes de longueur impaire (sans compter le caractère de fin de ligne \n) ont été retirées.

# Entrée
- fName : le nom du fichier à traiter

# Sorties
- pairs : une chaîne de caractère qui est une copie du contenu du fichier dont les lignes de longueur impaire ont été retirées
- err : nil si le fichier fName existe, errImpossible sinon

# Info
2022-2023, test3, exercice 6
*/

var errImpossible error = errors.New("le fichier n'existe pas")

func filtrePairs(fName string) (pairs string, err error) {

	file, err := os.Open(fName)

	if err != nil {
		return "", errImpossible
	}

	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan(){
		line := scanner.Text()

		lines = append(lines,line)
	}

	pairs = ""
	for i, v := range lines{
		if len(v) %2 == 0{
			pairs += string(v)
			if i < len(lines) - 2 {
				pairs += "\n"
			}
		}
	}

	return pairs, err
}
