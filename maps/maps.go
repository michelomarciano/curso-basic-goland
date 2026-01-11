package maps

import "fmt"

func Maps() {
	dados := map[string]string{
		"nome":  "Mike",
		"idade": "30",
		"email": "mike@example.com",
	}
	fmt.Println(dados)
	fmt.Println(dados["nome"])
}


func DeletarItemDoMap() {
	dados := map[string]string{
		"nome": "Mike",
		"idade": "30",
		"email": "mike@example.com",
	}
	delete(dados, "nome")
	fmt.Println(dados)
}

func AdicionarItemNoMap() {
	dados := map[string]string{
		"nome": "Mike",
		"idade": "30",
		"email": "mike@example.com",
	}
	dados["telefone"] = "1234567890"
	fmt.Println(dados)
}


