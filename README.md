# Exercices introdev

Collection d'exercices pour le cours d'introduction au développement au département informatique de l'IUT de Nantes.

## Instructions pour récupérer les exercices

Vous pouvez récupérer tout le contenu de ce dépôt avec la commande suivante : git clone https://gitlab.univ-nantes.fr/jezequel-l/exercices-introdev.git

Les semaines suivantes vous pourrez le mettre à jour avec les nouveaux exercices simplement en vous plaçant dans le dossier et en utilisant la commande suivante : git pull

## Instructions pour faire les exercices

Le dossier que vous aurez récupéré contient plusieurs dossiers (basiques, recherche, etc) qui eux-même contiennent des sous-dossiers (dans basiques il y a par exemple facteurspremiers, monnaie, etc). Chacun de ces sous-dossiers est un exercice.

Chaque exercice contient deux fichiers :
- xxx_test.go, décrivant un jeu de tests
- xxx.go, l'exercice en lui même

Pour résoudre un exercice vous devez coder la fonction Go qui est décrite dans le fichier de l'exercice et vérifier qu'elle passe bien tous les tests. Pour cela vous pouvez utiliser la commande **go test** en vous plaçant dans le dossier de l'exercice.

## Instructions pour choisir les exercices

Vous pouvez faire les exercices dans l'ordre que vous voulez. Dans chaque dossier vous trouverez des exercices faciles et d'autres plus difficiles, ne restez pas bloqués sur un exercice : passez à un autre et revenez-y plus tard.

Vous disposez aussi d'une page Web pour voir la difficulté (ressentie par votre promo et les anciens étudiants) des exercices. Cette page est accessible à l'adresse http://172.26.82.22 **depuis le réseau de l'IUT uniquement**. Votre identifiant est votre nom de famille en majuscules suivi de votre prénom (ex : JEZEQUEL Loïg) et votre mot de passe est votre nom de famille en majuscules (ex : JEZEQUEL).

N'oubliez pas, à chaque fois que vous résolvez un exercice, d'évaluer sa difficulté sur cette page pour aider les autres à repérer les exercices plus ou moins difficiles.

## Historique

### 17 septembre 2021

#### basiques
- facteurspremiers
- factorielle
- monnaie
- monnaie2
- monnaie3
- nombreparfait
- nombrepremier
- nombrespairs
- nombrespremiers
- stevej

#### recherche
- matrice
- maxoccurrences
- occurrences
- occurrencesmax
- prefixes
- tabtab

### 27 septembre 2021

#### basiques
- ensemble

#### recherche
- nombrespremiers
- nombrespremiers2

#### recursivite
- chainesbinaires
- chocolats
- palindrome
- sousensembles
- sousensembles2

### 1 octobre 2021

#### programmation dynamique
- piecesjaunes
- souschaine

#### recherche
- classement

#### recursivite
- huitreines
- huitreines2
- piecesjaunes
- racaman
- souschaine

### 18 octobre 2021

#### recherche
- doublons1
- doublons3
- doublons4
- doublons5

#### tri
- alphabetique
- bienrange
- decroissant
- valabs

### 2021-2022 Test machine 1

#### basiques
- alphabet (exercice 7)
- calcul (exercice 0)
- somme (exercice 4)

#### pointeurs
- double (exercice 1)
- varswitch (exercice 2)

#### recherche
- compte (exercice 6, remplace occurrences, qui est similaire)
- doublons2 (exercice 8)
- puissant (exercice 9)
- recherche (exercice 5)

#### recursivite
- suite (exercice 3)

### 2021-2022 Test machine 2

basiques devient divers

#### divers
- chiffres (exercice 9)

#### pointeurs
- add (exercice 0)
- copy (exercice 1)

#### recherche
- egalite (exercice 3)
- recherche2 (exercice 2)

#### recursivite
- conway (exercice 8)
- ppcm (exercice 5)
- syracuse (exercice 4)

#### tri
- tri (exercice 6)
- tri2 (exercice 7)

### 23 novembre 2021

#### fichiers
- acrostiche
- existe
- lignes
- lignes2

#### recherche
- egalite2

#### recursivite
- pgcd
- pgcd2
- ppcm2

### 2021-2022 Test machine 3

#### divers
- envers (exercice 0)

#### fichiers
- nombre (exercice 9)

#### pointeurs
- doubleptr (exercice 7)

#### recursivite
- compte (exercice 3)
- compter (exercice 4)
- factorielle (exercice 1)
- recherche (exercice 2)

#### structures
- classement (exercice 6)
- classer (exercice 5)
- longueur (exercice 8)

### 2022-2023 Test machine 1

#### divers
- bougies (exercice 5)
- decimales (exercice 9)
- divisible (exercice 4)

#### pointeurs
- setval (exercice 0)

#### recherche
- deuxgrands (exercice 8)
- moitie (exercice 6)
- multipledetrois (exercice 2)
- plusquezero (exercice 3)

#### recursivite
- suitemoinssimple (exercice 7)
- suitesimple (exercice 1)

### 2022-2023 Test machine 2

#### pointeurs
- somme (exercice 2)

#### recursivite
- combinaisons (exercice 8)
- desert (exercice 5)
- pairs (exercice 3)
- parentheses (exercice 9)
- suite2 (exercice 0)
- suiteconditionnelle (exercice 1)
- suitehistorique (exercice 7)
- suiteparametree (exercice 6)

#### tri
- pairsimpairs (exercice 4)

### 2022-2023 Test machine 3

#### fichiers
- filtre (exercice 6)
- unmot (exercice 5)

#### recursivite
- encoreunesuite (exercice 7)

#### structures
- livres (exercice 4)
- pages (exercice 3)
- queue (exercice 8)
- queue2 (exercice 9)
- titre (exercice 2)

#### tri
- longueur (exercice 1)
- vraifaux (exercice 0)

### 2022-2023 Test machine 4

#### pointeurs
- produit (exercice 1)

#### recherche
- deuxpetits (exercice 3)
- interval (exercice 2)
- moitie2 (exercice 4)
- positifs (exercice 5)

#### recursivite
- suitesimple (exercice 0, mise à jour)
- suitemoinssimple2 (exercice 6)
- suiteparametree2 (exercice 9)

#### tri
- decroissant (exercice 7, mise à jour)
- deuxcriteres (exercice 8)

### 2023-2024 Test machine 1

#### recherche
- comptemax (exercice 7)
- comptevrai (exercice 2)
- dernier (exercice 6)
- interval2 (exercice 5)
- multiplesdedeux (exercice 3)
- recherche (exercice 0, mise à jour)

#### recursivite
- ateddybearpicnic (exercice 9)
- suitedouble (exercice 8)
- suitesimple2 (exercice 1)

#### tri
- classement (exercice 4)

### 2023-2024 Test machine 2

#### divers
- alphanum (exercice 3)

#### recherche
- doublons6 (exercice 0)
- doublons7 (exercice 7)
- doublons8 (exercice 8)
- egalite (exercice 6, mise à jour)

#### recursivite
- sequences (exercice 9)
- somme (exercice 2)
- yetanothersuite (exercice 1)

#### structures
- contient (exercice 5)
- majeur (exercice 4)

### 2023-2024 Test machine 3

#### divers
- fibonacci (exercice 9)
- motdepasse (exercice 1)
- riz (exercice 5)

#### fichiers
- file2struct (exercice 6)
- file2tab (exercice 7)

#### structures
- fieldval (exercice 0)
- sortlist (exercice 8)
- sortstruct (exercice 3)
- sortstruct2 (exercice 4)

#### tri
- alphabetique (exercice 2, mise à jour)

### 2024-2025 Test machine 1

#### divers
- moyenne (exercice 2)

#### recherche
- egalite (exercice 6, déjà présent)
- egalite2 (exercice 7, déjà présent)
- inclusion (exercice 5)
- manquant (exercice 8)
- plusgrand (exercice 3)
- recherche (exercice 0, déjà présent)

#### recursivite
- hanoi (exercice 9)
- suite3 (exercice 1)
- suiteconditionnelle2 (exercice 4)

### 2024-2025 Test machine 2

#### divers
- mail (exercice 5)
- palindrome (exercice 9)

#### recherche
- trouvevrai (exercice 1)

#### recursivite
- recherchemin (exercice 6)
- suite4 (exercice 0)
- suiteconditionnelle3 (exercice 2)

#### structures
- cycle (exercice 8)

#### tri
- longueur (exercice 3, l'ancien longueur devient longueur2)
- rectangles (exercice 4)
- rectangles2 (exercice 7)