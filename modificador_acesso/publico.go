package modificador_acesso

import (
	"fmt"
)

func FuncaoPublica() {
	fmt.Println("Funcao Publica")
	funcaoNaoPublica()
}
