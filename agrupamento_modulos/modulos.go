package agrupamento_modulos

import (
	"fmt"
	"modulo/modificador_acesso"

	"github.com/badoux/checkmail"
)

func ExecutarModulos() {
	fmt.Println("--------------------------------")
	fmt.Println("MODULOS INTERNOS - MODIFICADOR DE ACESSO PUBLIC")
	fmt.Println("--------------------------------")
	modificador_acesso.FuncaoPublica()
	fmt.Print("\n")

	fmt.Println("--------------------------------")
	fmt.Println("MODULOS EXTERNOS - Checkmail")
	fmt.Println("--------------------------------")
	erro := checkmail.ValidateFormat("teste@teste.com")
	fmt.Println(erro)
	fmt.Print("\n")
}
