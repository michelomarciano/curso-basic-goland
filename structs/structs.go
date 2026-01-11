package structs

import "fmt"

func Structs() {

	type Endereco struct {
		Rua    string
		Numero int
		Cidade string
		Estado string
		CEP    string
	}

	type Usuario struct {
		ID       int
		Nome     string
		Email    string
		Senha    string
		Endereco Endereco
	}

	usuario := Usuario{1, "João", "joao@example.com", "123456", Endereco{"Rua das Flores", 123, "São Paulo", "SP", "1234567890"}}
	fmt.Printf("Usuario: %+v\n", usuario)
	
	usuario2 := Usuario{Nome: "Mike"}
	fmt.Printf("Usuario2: %+v\n", usuario2)
}
