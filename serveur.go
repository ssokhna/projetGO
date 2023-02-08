package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"strconv"
	"math/rand"
	"sync"
)

type res struct {
	ligne []int
	num_ligne  int
}

func worker_mult(ligne int, mat1, mat2 [][]int, resultat chan<- res) { //mat1 represente ici une forme de tableau car elle change d'indice pour chaque worker
	tab1 := mat1[ligne] 
	var res_ligne []int
	for i := 0; i < len(mat2[0]); i++ {
	//i->taille ligne mat2
		var somme int
		for j := 0; j < len(mat2); j++ {
			somme += tab1[j] * mat2[j][i]
		}
		res_ligne = append(res_ligne, somme)
	}
	resultat <- res{res_ligne, ligne}
}


func main() {

	args := os.Args
	if len(args) == 1 {
			fmt.Println("Veuiller entrer le port")
			return
	}

	port := ":" + args[1]
	l, err := net.Listen("tcp", port)
	if err != nil {
			fmt.Println(err)
			return
	}
	defer l.Close()

	c, err := l.Accept()
	if err != nil {
			fmt.Println(err)
			return
	}

	for {
			netData, err := bufio.NewReader(c).ReadString('\n')
			fmt.Print("-> ", string(netData))//afficher le msg recu du client
			if err != nil {
					fmt.Println(err)
			}
			if strings.TrimSpace(string(netData)) == "s" {
					fmt.Println("Exiting TCP server!")
					return
			}
			
			netData= strings.TrimSpace(netData) //enlever les spaces et retour en ligne dans le string
			taille, err := strconv.Atoi(netData)
			if err != nil {
				fmt.Println("Erreur de conversion :", err)
				return
			}
			fmt.Println("taille des matrices :", taille)
			lignes:=taille
			colonnes:=taille
			// Crée une matrice vide
			mat1 := make([][]int, lignes)
			for i := range mat1 {
				mat1[i] = make([]int, colonnes)
			}// Remplit la matrice avec des valeurs aléatoires
			for i := 0; i < lignes; i++ {
				for j := 0; j < colonnes; j++ {
					mat1[i][j] = rand.Intn(10) // Génère un entier aléatoire compris entre 0 et 9
				}
			}// Affiche la matrice générée aléatoirement

			mat2 := make([][]int, lignes)
			for i := range mat2 {
				mat2[i] = make([]int, colonnes)
			}
			for i := 0; i < lignes; i++ {
				for j := 0; j < colonnes; j++ {
					mat2[i][j] = rand.Intn(10) 
				}}
			fmt.Println(mat1 ," et ", mat2)

			var wg sync.WaitGroup
			resultat := make(chan res, len(mat1))

			for i := 0; i < len(mat1); i++ {
				wg.Add(1)
				go func(i int) {
					defer wg.Done()
					worker_mult(i, mat1, mat2, resultat)
				}(i) 
			}

			go func() {
				wg.Wait()
				close(resultat)
			}()

			var matC [][]int
			for r := range resultat {
				matC = append(matC, r.ligne)
			}
			fmt.Println("La multiplication donne: ")
			fmt.Println(matC) 

			monMessage :=  "Done!!\n"
			c.Write([]byte(monMessage))

		}

	}