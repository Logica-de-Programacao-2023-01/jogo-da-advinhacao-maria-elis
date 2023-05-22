package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var tentativasTotais = [][]int{}
	repetir := true

	for repetir == true {

		x := rand.Intn(100) + 1
		var y int
		var tentativas = []int{}
		fmt.Println("Bem vindo ao meu jogo de adivinhação, tente acertar o número que estou pensando")
		fmt.Print("Digite um número entre 1 e 100: ")
		fmt.Scan(&y)

		for {
			if x < y {
				tentativas = append(tentativas, y)
				fmt.Printf("O número misterioso é menor que %d, tente novamente: ", y)
				fmt.Scan(&y)
			}
			if x > y {
				tentativas = append(tentativas, y)
				fmt.Printf("O número misterioso é maior que %d, tente novamente: ", y)
				fmt.Scan(&y)
			}
			if x == y {
				tentativas = append(tentativas, y)
				fmt.Println("Parabéns, você é um gênio, acertou o número.")
				fmt.Println("Você utilizou", len(tentativas), "tentativas.")
				break
			}
		}

		tentativasTotais = append(tentativasTotais, tentativas)

		var res string
		fmt.Print("Você quer jogar novamente? (s/n): ")
		fmt.Scan(&res)
		if res != "s" && res != "S" {
			repetir = false
		}

	}

	for i := range tentativasTotais {
		fmt.Println("Você usou", len(tentativasTotais[i]), "tentativas no jogo", i+1)
	}
	fmt.Println("Obrigado por jogar!")

}
