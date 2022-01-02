package exercicio2

type produto struct {
	nome  string
	preco float64
	qtd   int
}

func NewProduto(nome string, preco float64, qtd int) *produto {
	p := produto{
		nome:  nome,
		preco: preco,
		qtd:   qtd,
	}

	return &p
}

func (p *produto) Nome() string   { return p.nome }
func (p *produto) Preco() float64 { return p.preco }
func (p *produto) Qtd() int       { return p.qtd }
