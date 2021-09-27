package sousensembles2

import (
	"errors"
)

var errPasEnsemble error = errors.New("ceci n'est pas un ensemble")

/*
Étant donné un ensemble E on peut construire tous ces sous-ensembles de manière récursive en remarquant que si E' est E\{x} pour un élément x de E, alors les sous-ensembles de E sont les sous-ensembles de E' auxquels s'ajoutent ces mêmes sous-ensembles augmentés de x.
La fonction sousEnsembles doit mettre en œuvre cette construction pour construire uniquement des sous-ensembles d'une longueur donnée.

# Entrée
- E : un tableau représentant un ensemble d'entiers
- k : un entier indiquant la taille des sous-ensembles à construire

# Sortie
- PE : un tableau de tableaux contenant tous les sous-ensembles de taille k de E (sans doublons)
- err : une erreur qui vaut errPasEnsemble si et seulement si E n'est pas un ensemble (il contient plusieurs fois la même valeur ou il vaut nil)

# Exemple
sousEnsembles([]int{1, 2}, 1) = [[1] [2]] (l'ordre des ensembles et les ordres des valeurs dans les ensembles n'ont pas d'importance)
*/
func sousEnsembles(E []int, k int) (PE [][]int, err error) {
	return PE, err
}
