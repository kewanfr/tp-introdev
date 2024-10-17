package alphabet

/*
La fonction alphabet doit indiquer si deux mots sont dans l'ordre alphabétique.
On considérera que les mots en questions ne sont constitués que de lettres
minuscules non accentuées, donc prises dans l'ensemble {a, b, c, d, e, f, g, h,
i, j, k, l, m, n, o, p, q, r, s, t, u, v, w, x, y, z}.

# Entrées
- m1 : un mot constitué de lettres minuscules non accentuées
- m2 : un mot constitué de lettres minuscules non accentuées

# Sorties
- m1avantm2 : un booléen qui vaut true si m1 est (strictement) avant m2 dans
l'ordre alphabétique et qui vaut false sinon

# Indication
On peut comparer les caractères d'une chaîne de caractères comme on compare des
entiers. Ainsi, m1[i] < m2[j] vaut true si et seulement si le i-ième caractère
de m1 est avant le j-ième caractère de m2 dans l'alphabet.

# Info
2021-2022, test 1, exercice 7
*/

// RÉSOLU - 02/10/2024
// TEST 1 - 2021-2022 

func alphabet(m1, m2 string) (m1avantm2 bool) {

 	// Var si chaque lettre analysée était égale
	var seulementEgal bool = false

	// Boucle dans chaque lettre, avec condition d'arrêt pour les deux
	for i := 0; i < len(m1) && i < len(m2); i++ {

		if m1[i] == m2[i] {
			// Si la lettre est la même
			// Pour le moment, m1 est avant m2 et chaque lettre était égale
			m1avantm2 = true
			if(seulementEgal) {
				seulementEgal = true
			}

		} else if m1[i] < m2[i] {
			// Si la lettre de m1 est avant celle de m2
			// Pour le moment, m1 est avant m2, mais tt les lettres ne sont pas égales
			m1avantm2 = true
			seulementEgal = false
			break
		} else if m1[i] > m2[i] {
			// Si la lettre de m2 est après
			// m1 est forcément après m2
			m1avantm2 = false
			seulementEgal = false
			break
		}
	}


	if len(m2) < len(m1) && seulementEgal{
		m1avantm2 = false
	}

	if m1 == "" {
		m1avantm2 = true
	}

	if m1 == m2 {
		m1avantm2 = false
	}


	return m1avantm2
}
