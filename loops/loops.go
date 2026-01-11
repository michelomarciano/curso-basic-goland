package loops

import (
	"fmt"
	"time"
)

func LoopFor() {
	for i := 0; i < 10; i++ {
		fmt.Println("Loop For: ", i)
		time.Sleep(time.Second * 1)
	}
}

func LoopWhile() {
	i := 0
	for i < 10 {
		fmt.Println("Loop While: ", i)
		i++
	}
}

func LoopDoWhile() {
	i := 0
	for {
		fmt.Println("Loop Do While: ", i)
		i++
		if i >= 10 {
			break
		}
	}
}

func LoopForRange() {
	slice := []string{"Golang", "Python", "Java", "JavaScript", "C#"}
	for indice, valor := range slice {
		fmt.Println("Loop For Range: O valor do indice é: ", indice)
		fmt.Println("Loop For Range: O valor do valor é: ", valor)
	}
}

func LoopForRangeString() {
	texto := "Golang"
	for indice, valor := range texto {
		fmt.Println("Loop For Range String: O valor do indice é: ", indice)
		fmt.Println("Loop For Range String: O valor do valor é: ", valor)
		fmt.Println("Loop For Range String: O valor do valor é: ", string(valor))
	}
}

func LoopForRangeMap() {
	mapa := map[string]string{"Golang": "2009", "Python": "1991", "Java": "1995"}
	for chave, valor := range mapa {
		fmt.Println("Loop For Range Map: O valor da chave é: ", chave)
		fmt.Println("Loop For Range Map: O valor do valor é: ", valor)
	}
}

// func LoopInfinito() {
// 	for {
// 		fmt.Println("Loop Infinito")
// 		time.Sleep(time.Second * 1)
// 	}
// }
