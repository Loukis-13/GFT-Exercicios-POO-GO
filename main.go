package main

import (
	"fmt"

	veiculo "github.com/Loukis-13/GFT-Exercicios-POO-GO/exercicio_1"
	exercicio2 "github.com/Loukis-13/GFT-Exercicios-POO-GO/exercicio_2"
	personagem "github.com/Loukis-13/GFT-Exercicios-POO-GO/exercicio_3"
	pessoa "github.com/Loukis-13/GFT-Exercicios-POO-GO/exercicio_4"
	funcionario "github.com/Loukis-13/GFT-Exercicios-POO-GO/exercicio_7"
)

func main() {
	testarVeiculo()
	testarLoja()
	testarPersonagem()
	testarPessoa()
	testarFuncionario()
}

func testarVeiculo() {
	carro := veiculo.NewVeiculo("Volksvagen", "Gol", "GFT-0311", "branco", 0, 30, 700000)

	carro.Acelerar()
	carro.Ligar()
	carro.Acelerar()
	carro.Acelerar()
	fmt.Println()

	fmt.Println("Litros combustivel:", carro.GetLitrosCombustivel())
	carro.Abastecer(20)
	carro.Abastecer(200)
	fmt.Println()

	fmt.Println("Cor anterior: " + carro.GetCor())
	carro.Pintar("preto")
	fmt.Println("Cor atual: " + carro.GetCor())
	fmt.Println()

	carro.Desligar()
	carro.Frear()
	carro.Frear()
	carro.Frear()
	carro.Frear()
	carro.Desligar()

	fmt.Println("---------------------------------------------------------------------")
}

func testarLoja() {
	l1 := exercicio2.NewLivro("Harry poter", "J. K. Rolling", "fantasia", 40, 50, 300)
	l2 := exercicio2.NewLivro("Senhor dos Anéis", "J. R. R. Tolkien", "fantasia", 60, 30, 500)
	l3 := exercicio2.NewLivro("Java POO", "GFT", "educativo", 20, 50, 500)

	ps4 := exercicio2.NewVideoGame("PS4", "Sony", "SLim", 1800, 100, false)
	ps4Usado := exercicio2.NewVideoGame("PS4", "Sony", "SLim", 1000, 7, true)
	xbox := exercicio2.NewVideoGame("XBOX", "Microsoft", "One", 1500, 500, false)

	livros := []exercicio2.Livro{l1, l2, l3}

	games := []exercicio2.VideoGame{ps4.(exercicio2.VideoGame), ps4Usado.(exercicio2.VideoGame), xbox.(exercicio2.VideoGame)}

	americanas := exercicio2.NewLoja("Americanas", "123456789", livros, games)

	fmt.Println("Imposto livro:", l2.CalculaImposto())
	fmt.Println("Imposto livro educativo:", l3.CalculaImposto())

	fmt.Println("Imposto PS4 novo:", ps4.CalculaImposto())
	fmt.Println("Imposto PS4 usado:", ps4Usado.CalculaImposto())
	fmt.Println()

	americanas.ListaLivros()
	americanas.ListaVideoGames()
	fmt.Printf("O patrimonio da loja: %s é R$%.2f\n", americanas.Nome(), americanas.CalculaPatrimonio())

	fmt.Println("---------------------------------------------------------------------")
}

func testarPersonagem() {
	fmt.Println("Quantidade de personagens criados:", personagem.QntPersonagensCriados)

	mago := personagem.NewMago("Gandalf")
	guerreiro := personagem.NewGuerreiro("Talion")

	fmt.Println("Quantidade de personagens criados:", personagem.QntPersonagensCriados)

	mago.LvlUp()
	fmt.Println("Mago Level:", mago.Level())
	fmt.Println("Mago ataca:", mago.Attack())

	guerreiro.LvlUp()
	guerreiro.LvlUp()
	fmt.Println("Guerreiro Level:", guerreiro.Level())
	fmt.Println("Guerreiro ataca:", guerreiro.Attack())

	fmt.Println("Inteligencia mago:", mago.Inteligencia())
	mago.AprenderMagia("Impedir que o inimigo passe")
	fmt.Println("Inteligencia mago:", mago.Inteligencia())

	fmt.Println("---------------------------------------------------------------------")
}

func testarPessoa() {
	pessoas := []pessoa.Pessoa{
		{Nome: "João", Idade: 15},
		{Nome: "Leandro", Idade: 21},
		{Nome: "Paulo", Idade: 17},
		{Nome: "Jessica", Idade: 18},
	}

	var pessoaMaisVelha pessoa.Pessoa
	var jessica *pessoa.Pessoa
	pessoasMaisVelhasQue18 := 0

	for _, i := range pessoas {
		if i.Idade > pessoaMaisVelha.Idade {
			pessoaMaisVelha = i
		}
		if i.Idade > 18 {
			pessoasMaisVelhasQue18++
		}
		if i.Nome == "Jessica" {
			jessica = &i
		}
	}

	fmt.Println("A pessoa mais velha é", pessoaMaisVelha.Nome)

	fmt.Println()
	fmt.Println("Quantidade de pessoas:", len(pessoas))
	fmt.Println("Quantidade de pessoas maiores que 18:", pessoasMaisVelhasQue18)
	fmt.Println()

	if jessica != nil {
		fmt.Println("Jessica encontrada, idade:", jessica.Idade)
	} else {
		fmt.Println("Jessica não encontrada")
	}

	fmt.Println("---------------------------------------------------------------------")
}

func testarFuncionario() {
	gerente := funcionario.Gerente{Nome: "Tomate", Idade: 40, Salario: 5000}
	supervisor := funcionario.Supervisor{Nome: "Pepino", Idade: 30, Salario: 4000}
	vendedor := funcionario.Vendedor{Nome: "Beringela", Idade: 20, Salario: 3000}

	fmt.Printf("Gerente %s bonificado %.2f\n", gerente.Nome, gerente.Bonificacao())
	fmt.Printf("Supervisor %s bonificado %.2f\n", supervisor.Nome, supervisor.Bonificacao())
	fmt.Printf("Vendedor %s bonificado %.2f\n", vendedor.Nome, vendedor.Bonificacao())

	fmt.Println("---------------------------------------------------------------------")
}
