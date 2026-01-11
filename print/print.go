package print

import "fmt"


func Print(){
	var(
		nome string = "Mike"
        sobrenome string = "Marciano"
		salario float32 = 100.20
	)
	fmt.Print("O nome é: " + nome + 
	" sobrenome: " + sobrenome + 
	" salario: " +
	 fmt.Sprint(salario))
}


