package main

import "fmt"

func mult(mat1[4][4] int, mat2[4][4] int) int {
    var somme int
    var mat3[][] int
    for i:=0; i<=4; i++ {
        somme =0
        fmt.Println(" \n ")
        for j:=0; i<=4; j++ {
            somme = somme + mat1[i][i]*mat2[i][j]
            mat3[i][j] == somme
            
        }
    }
}

func main() {

    mat1 :=[][]int{{1, 2},{4,8},{7, 10},{9,7}}
    mat2 := [][]int{{5, 3},{0,5},{1,8},{4,3}}
    fmt.Println("Produit : \n", mult(mat1,mat2)) 
}