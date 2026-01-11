package maps

import "fmt"

func AdicionarMapEmOutroMap(){
	dados := map[string]map[string]string{
		"nome": {
			"nome": "Mike",
			"idade": "30",
			"email": "mike@example.com",
		},
	}
	dados["endereco"] = map[string]string{
		"rua": "Rua das Flores",
		"numero": "123",
		"cidade": "São Paulo",
		"estado": "SP",
		"cep": "1234567890",
	}
	fmt.Println("AdicionarMapEmOutroMap: ", dados)
	fmt.Println("AdicionarMapEmOutroMap: ", dados["nome"]["nome"])
	fmt.Println("AdicionarMapEmOutroMap: ", dados["endereco"]["rua"])
}