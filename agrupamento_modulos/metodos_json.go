package agrupamento_modulos

import (
	"encoding/json"
	"fmt"
	jsons "modulo/json"
	"modulo/metodos"
)

func ExecutarMetodosEJSON() {
	fmt.Println("--------------------------------")
	fmt.Println("METODOS")
	fmt.Println("--------------------------------")
	usuario := metodos.Usuario{Nome: "Mike", Email: "mike@example.com", Senha: "123456", Idade: 20}
	usuario.Salvar()
	usuario.AtualizarIdade()
	fmt.Println("Idade: ", usuario.Idade)
	fmt.Print("\n")

	fmt.Println("--------------------------------")
	fmt.Println("JSONS")
	fmt.Println("--------------------------------")
	jsonExemplo := jsons.Usuario{Nome: "Mike", Idade: 20, Email: "mike@example.com"}
	fmt.Println(jsonExemplo)
	jsonBytes, err := json.Marshal(jsonExemplo)
	if err != nil {
		fmt.Println("Erro ao fazer Marshal:", err)
	} else {
		fmt.Println("JSON:", string(jsonBytes))
	}
}

