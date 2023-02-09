package main
import "fmt"
//import "sync"
import "time"

type res struct {
    debut int
    fin int
    M [][]int
}

func workers(id int,l1 int,l2 int,matA[][] int, matB[][] int, resultat chan<- res){
	//var wg sync.WaitGroup
    //defer wg.Done()
    var res_lignes [][]int
    for i :=l1; i<l2; i++{
        var res_ligne []int
        var somme int
        for j:=0; j<=(len(matA[i])-1); j++ {
            somme =0
            for k:=0; k<=(len(matB)-1); k++ {
                somme = somme + matA[i][k]*matB[k][j]
            }
            res_ligne =append(res_ligne,somme)
        }
        res_lignes = append(res_lignes, res_ligne)
}
    resultat <- res{l1,l2,res_lignes}
}


func main(){
    mat1 := [][]int{{4,2,8,6},{4,3,4,3},{7,0,7,1},{5,7,3,5}}
    mat2 := [][]int{{5,3,0,6},{1,5,2,3},{1,8,2,3},{4,3,8,5}}

   // var wg sync.WaitGroup
	resultat := make(chan res, 4)
    //wg.Add(1)
    start := time.Now()
	go workers (1,0,1,mat1,mat2,resultat)
    go workers (2,1,2,mat1,mat2,resultat)
    go workers (3,2,3,mat1,mat2,resultat)
    go workers (4,3,4,mat1,mat2,resultat)
    fmt.Println(" 4 Worker Temps d'execution:", time.Since(start))
    ///wg.Wait()
    

    fmt.Println(" ---- ")

    start1 := time.Now()
	go workers (1,0,2,mat1,mat2,resultat)
    go workers (2,2,3,mat1,mat2,resultat)
    fmt.Println(" 2 Worker Temps d'execution:", time.Since(start1))
    


    fmt.Println(" ---- ")

    start2 := time.Now()
	go workers (1,0,3,mat1,mat2,resultat)
    fmt.Println(" 1 Worker Temps d'execution:", time.Since(start2))
   
    fmt.Println(" ---- ")
    var matC [][]int
    r := <- resultat
    matC = r.M
    for i := 1; i < 4; i++ {
        r := <- resultat
        matC = append(matC, r.M...)
    }
    fmt.Println (matC)
    
}


