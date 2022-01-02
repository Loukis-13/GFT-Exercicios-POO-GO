package personagem

import (
	"math/rand"
	"time"
)

type Guerreiro Personagem

func NewGuerreiro(nome string) Guerreiro {
	rand.Seed(time.Now().UnixNano())

	p := NewPersonagem(
		nome,
		rand.Intn(35)+30,
		rand.Intn(50),
		rand.Intn(50),
		rand.Intn(50)+30,
	)

	QntPersonagensCriados++

	return Guerreiro(p)
}

// metodos

func (g *Guerreiro) LvlUp() {
	g.level++
	g.vida += rand.Intn(50) + 30
	g.forca += rand.Intn(50) + 30
	g.mana += rand.Intn(50)
	g.inteligencia += rand.Intn(50)
}

func (g *Guerreiro) Attack() int { return g.forca*g.level + rand.Intn(301) }

func (g *Guerreiro) AprenderHabilidade(habilidade string) { g.forca += 30 }

// getters

func (p *Guerreiro) Nome() string      { return p.nome }
func (p *Guerreiro) Vida() int         { return p.vida }
func (p *Guerreiro) Mana() int         { return p.mana }
func (p *Guerreiro) Inteligencia() int { return p.inteligencia }
func (p *Guerreiro) Forca() int        { return p.forca }
func (p *Guerreiro) Level() int        { return p.level }
func (p *Guerreiro) Xp() float32       { return p.xp }
