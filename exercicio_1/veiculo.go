package veiculo

import (
	"fmt"
)

type veiculo struct {
	marca             string
	modelo            string
	placa             string
	cor               string
	isLigado          bool
	velocidade        int
	litrosCombustivel int
	km                float32
	preco             float64
}

func NewVeiculo(marca, modelo, placa, cor string, km float32, litrosCombustivel int, preco float64) veiculo {
	veic := veiculo{
		marca:             marca,
		modelo:            modelo,
		placa:             placa,
		cor:               cor,
		litrosCombustivel: litrosCombustivel,
		km:                km,
		preco:             preco,
		isLigado:          false,
		velocidade:        0,
	}

	return veic
}

// metodos

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (veic *veiculo) Acelerar() {
	if veic.isLigado {
		veic.velocidade += 20
		fmt.Println("Acelerou, velocidade:", veic.velocidade)
	} else {
		fmt.Println("O carro está desligado")
	}
}

func (veic *veiculo) Abastecer(combustivel int) {
	veic.litrosCombustivel = min(veic.litrosCombustivel+combustivel, 60)
	fmt.Printf("Litros combustivel: %d\n", veic.litrosCombustivel)
}

func (veic *veiculo) Frear() {
	if veic.velocidade > 0 {
		veic.velocidade -= 20
		fmt.Printf("Desacelerou, velocidade: %d\n", veic.velocidade)
	} else {
		fmt.Println("O carro está parado")
	}
}

func (veic *veiculo) Pintar(cor string) { veic.cor = cor }

func (veic *veiculo) Ligar() {
	if !veic.isLigado {
		veic.isLigado = true
		fmt.Println("Ligou")
	}
}

func (veic *veiculo) Desligar() {
	if veic.velocidade <= 0 && veic.isLigado {
		veic.isLigado = false
		fmt.Println("Desligou")
	} else {
		fmt.Println("Não pode desligar em movimento")
	}
}

// getters

func (veic *veiculo) GetMarca() string          { return veic.marca }
func (veic *veiculo) GetModelo() string         { return veic.modelo }
func (veic *veiculo) GetPlaca() string          { return veic.placa }
func (veic *veiculo) GetCor() string            { return veic.cor }
func (veic *veiculo) IsLigado() bool            { return veic.isLigado }
func (veic *veiculo) GetVelocidade() int        { return veic.velocidade }
func (veic *veiculo) GetLitrosCombustivel() int { return veic.litrosCombustivel }
func (veic *veiculo) GetKm() float32            { return veic.km }
func (veic *veiculo) GetPreco() float64         { return veic.preco }
