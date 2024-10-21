package existe

/*
La fonction existe doit dire si un fichier dont le nom est donné en paramètre
existe ou pas.

# Entrée
- fName : un nom de fichier

# Sortie
- ok : un booléen qui vaut true si le fichier de nom fName existe et false sinon
*/

import (
	"errors"
	"os"
)

func existe(fName string) (ok bool) {

	_, err := os.Stat(fName)
	
	return !errors.Is(err, os.ErrNotExist)
}
