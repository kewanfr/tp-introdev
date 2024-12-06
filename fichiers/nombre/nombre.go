package nombre

import "os"

/*
La fonction nombre doit dire si le fichier dont le chemin est donné en paramètre
contient ou ne contient pas le chiffre 1. On suppose que le fichier en question
existe toujours (vous n'avez pas à vous préocupper des erreurs à l'ouverture du
fichier).

# Entrée
- chemin : le chemin pour accéder au fichier à traiter

# Sortie
- contient : un booléen qui vaut true si le chiffre 1 apparaît dans le fichier
             considéré et false sinon (le nom du fichier ne compte pas, seul
             son contenu est à prendre en compte).

# Info
2021-2022, test3, exercice 9
*/

func nombre(chemin string) (contient bool) {

    data, _ := os.ReadFile(chemin)
    contient = false

    for i := 0; i < len(data); i ++ {
        if string(data[i]) == "1"{
            contient = true
        }
    }

	return contient
}
