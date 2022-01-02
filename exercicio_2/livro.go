package exercicio2

type Livro struct {
	produto
	autor  string
	tema   string
	qntPag int
}

func NewLivro(nome, autor, tema string, preco float64, qtd int, qntPag int) Livro {
	p := produto{
		nome:  nome,
		preco: preco,
		qtd:   qtd,
	}

	l := Livro{
		produto: p,
		autor:   autor,
		tema:    tema,
		qntPag:  qntPag,
	}

	return l
}

// metodos

func (v Livro) CalculaImposto() float64 {
	if v.tema == "educativo" {
		return 0
	}
	return v.preco * .1
}

// getters e setters

func (v *Livro) Autor() string { return v.autor }
func (v *Livro) Tema() string  { return v.tema }
func (v *Livro) QntPag() int   { return v.qntPag }
