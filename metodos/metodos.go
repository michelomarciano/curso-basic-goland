package metodos

import "fmt"

type Usuario struct {
	Nome string
	Email string
	Senha string
	Idade int
}

func (u *Usuario) AtualizarIdade() {
	u.Idade++
}

func (u Usuario) Salvar() {
	fmt.Println("Salvando usuario: ", u.Nome)
	fmt.Println("Email: ", u.Email)
	fmt.Println("Senha: ", u.Senha)
}