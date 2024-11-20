package doublons8

/*
On dispose d'un tableau d'entiers de longueur n et on suppose qu'il contient exactement p fois chaque nombre compris entre k et k + (n/p) - 1 (pas forcément dans l'ordre) pour un certain k et un certain p non connus à l'avance (p différent de 0). On voudrait vérifier que c'est bien le cas. C'est le travail de la fonction doublons.

# Entrée
- tab : un tableau d'entiers

# Sortie
- ok : un booléen qui doit valoir true si tab contient exactement p fois chaque entier de k à k + len(tab)/p - 1 et qui vaut false sinon

# Info
2023-2024, test 2, exercice 8
*/

func doublons(tab []int) (ok bool) {

	var t map[int]int = make(map[int]int)
	
	for i:= 0; i < len(tab); i++ {
		if _, ok := t[tab[i]]; ok {
			t[tab[i]] += 1
		} else {
			t[tab[i]] = 1
		}
	}


	var n int = len(tab)
	var p int = 0
	ok = true

	var k, kMax int
	k = tab[0]
	kMax = tab[0]

	for i, v := range t {
		if p == 0 {
			p = v
		}

		if v != p {
			ok = false
		}

		if i < k {
			k = i
		}

		if i > kMax {
			kMax = i
		}
	}

	if !ok {
		return ok
	}

	if kMax != k + (n/p) - 1 {
		ok = false
		return ok
	}

	for nb := k; nb < kMax; nb++ {
		estDedans := false
		for _, val := range tab {
			if val == nb {
				estDedans = true
			}
		}
		if !estDedans {
			ok = false
		}
	}

	return ok
}
