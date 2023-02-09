Dans le fichier "matrice_final.go":
Ce programme calcule la multiplication de matrices en utilisant le modèle de programmation concurrent. 
Il y a une fonction principale nommée "main" et une fonction de travailleur nommée "workers".
La fonction "workers" prend en entrée l'id du travailleur, le marge de ligne à traiter, les matrices à multiplier, et le canal de résultats. 
La fonction "main" définit les matrices "mat1" et "mat2" et les initialise avec des valeurs. Ensuite, elle définit
une canal "resultat" pour stocker les résultats du travail des travailleurs.
La fonction de travailleur calcule la multiplication des matrices en parcourant chaque ligne et en multipliant
chaque élément de la première matrice par chaque élément de la seconde matrice. 
Les résultats sont ensuite stockés dans un canal "resultat" sous forme d'une structure "res".
Dans la fonction "main", il y a différents tests effectués avec un nombre différent de travailleurs pour calculer 
le temps d'exécution. Dans chaque test, un certain nombre de travailleurs sont lancés et le temps d'exécution est mesuré. 
Enfin, les résultats sont lus à partir du canal "resultat" et stockés dans la matrice "matC".
Le résultat final est affiché à l'aide de la commande "fmt.Println (matC)".

Dans les fichiers serveur.go et client.go:
serveur.go est un serveur TCP qui écoute un numéro de port donné et accepte une seule connexion client.
Le serveur effectue une multiplication matricielle entre deux matrices de taille N (où N est une valeur reçue du client).
La multiplication matricielle est effectuée à l'aide de Goroutines et WaitGroup pour synchroniser les travailleurs.
Le serveur reçoit la taille des matrices (N) du client, génère deux matrices de taille N remplies de valeurs aléatoires,
puis effectue la multiplication matricielle en divisant le travail entre plusieurs Goroutines. Chaque Goroutine calcule le
produit d'une seule ligne de la première matrice et de toutes les colonnes de la seconde matrice. Les résultats de 
chaque Goroutine sont collectés dans un canal puis ajoutés à une matrice résultante. Enfin, la matrice résultante est imprimée sur la console.
Si l'utilisateur entre "exit", la boucle se termine et le client se ferme.

