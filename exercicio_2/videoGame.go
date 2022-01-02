package exercicio2

type VideoGame struct {
	produto
	marca   string
	modelo  string
	isUsado bool
}

func NewVideoGame(nome, marca, modelo string, preco float64, qtd int, isUsado bool) Imposto {
	p := produto{
		nome:  nome,
		preco: preco,
		qtd:   qtd,
	}

	v := VideoGame{
		produto: p,
		marca:   marca,
		modelo:  modelo,
		isUsado: isUsado,
	}

	var imp Imposto = v

	return imp
}

// metodos

func (v VideoGame) CalculaImposto() float64 {
	if v.isUsado {
		return v.preco * .25
	}
	return v.preco * .45
}

// getters e setters

func (v *VideoGame) Marca() string  { return v.marca }
func (v *VideoGame) Modelo() string { return v.modelo }
func (v *VideoGame) IsUsado() bool  { return v.isUsado }

func (v *VideoGame) SetIsUsado(isUsado bool) {
	if v.isUsado {
		return
	}
	v.isUsado = isUsado
}
