package main

import (
	"fmt"
)

type Familia interface {
	Dados() string
}

type Pai struct {
	Nome  string
	Idade int
	Cpf   string `json:"cpf"`
}

func (p Pai) Dados() string {
	return fmt.Sprintf("Nome: %s, Idade: %d", p.Nome, p.Idade)
}

type Filho struct {
	Pai
	Email string
}

type Filhos []Filho

func (f Filho) Dados() string {
	return fmt.Sprintf("\nNome: %s, Idade: %d, Email: %s", f.Nome, f.Idade, f.Email)
}

func (f Filhos) Dados() string {
	concat := ""
	for _, v := range f {
		concat += fmt.Sprintf("\nNome: %s, Idade: %d, Email: %s", v.Nome, v.Idade, v.Email)
	}
	return concat
}

func showDados(membro Familia) {
	fmt.Println(membro.Dados())
}

func main() {
	var filhos Filhos
	var filho Filho

	pai := new(Pai)
	pai.Nome = "Jefferson"
	pai.Idade = 38
	pai.Cpf = "00.xxx.xxx-xx"

	filho.Nome = "Arthur"
	filho.Idade = 3
	filho.Email = "arthur@gmail.com"
	filhos = append(filhos, filho)

	filho.Nome = "Francisco"
	filho.Idade = 7
	filho.Email = "francisco@gmail.com"
	filhos = append(filhos, filho)

	filha := new(Filho)
	filha.Nome = "Joyce"
	filha.Idade = 22
	filha.Email = "joyce@gmail.com"

}
