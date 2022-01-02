package exercicio2

import (
	"fmt"
)

type loja struct {
	nome       string
	cnpj       string
	livros     []Livro
	videoGames []VideoGame
}

func NewLoja(nome, cnpj string, livros []Livro, videoGames []VideoGame) loja {
	l := loja{
		nome:       nome,
		cnpj:       cnpj,
		livros:     livros,
		videoGames: videoGames,
	}

	return l
}

// metodos

func (l *loja) ListaLivros() {
	if len(l.livros) > 0 {
		for _, livro := range l.livros {
			fmt.Printf("Titulo: %s, Preço: R$%.2f, Quantidade: %d em estoque\n", livro.nome, livro.preco, livro.qtd)
		}
		fmt.Println()
	} else {
		fmt.Println("A loja não tem livros no seu estoque")
	}
}

func (l *loja) ListaVideoGames() {
	if len(l.videoGames) > 0 {
		for _, videoGame := range l.videoGames {
			fmt.Printf("Titulo: %s, Preço: R$%.2f, Quantidade: %d em estoque\n", videoGame.nome, videoGame.preco, videoGame.qtd)
		}
		fmt.Println()
	} else {
		fmt.Println("A loja não tem video-games no seu estoque")
	}
}

func (l *loja) CalculaPatrimonio() (total float64) {
	for _, i := range l.livros {
		total += i.preco * float64(i.qtd)
	}
	for _, i := range l.videoGames {
		total += i.preco * float64(i.qtd)
	}
	return
}

// getters e setters

func (l *loja) Nome() string            { return l.nome }
func (l *loja) Cnpj() string            { return l.cnpj }
func (l *loja) Livros() []Livro         { return l.livros }
func (l *loja) VideoGames() []VideoGame { return l.videoGames }

func (l *loja) SetLivros(livros []Livro)             { l.livros = livros }
func (l *loja) SetVideoGames(videoGames []VideoGame) { l.videoGames = videoGames }
