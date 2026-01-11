package funcoes

import "fmt"

func recuperarExecucao(){
	if r := recover(); r != nil {
		fmt.Println("Recuperado de panic:", r)
	}
}

func FuncaoPanic(n1, n2 int8) {
	defer recuperarExecucao()
   if media := (n1 + n2) / 2; media < 6 {
	panic("Media menor que 6")
   }
   fmt.Println("Media calculada. Resultado será retornado")
}
