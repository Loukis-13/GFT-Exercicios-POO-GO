package personagem

import (
	"math/rand"
	"time"
)

type Mago Personagem

func NewMago(nome string) Mago {
	rand.Seed(time.Now().UnixNano())

	p := NewPersonagem(
		nome,
		rand.Intn(35)+15,
		rand.Intn(50)+30,
		rand.Intn(50)+30,
		rand.Intn(50),
	)

	QntPersonagensCriados++

	return Mago(p)
}

// metodos

func (m *Mago) LvlUp() {
	m.level++
	m.vida += rand.Intn(50)
	m.forca += rand.Intn(50)
	m.mana += rand.Intn(50) + 30
	m.inteligencia += rand.Intn(50) + 30
}

func (m *Mago) Attack() int { return m.inteligencia*m.level + rand.Intn(301) }

func (m *Mago) AprenderMagia(magia string) { m.inteligencia += 30 }

// getters

func (p *Mago) Nome() string      { return p.nome }
func (p *Mago) Vida() int         { return p.vida }
func (p *Mago) Mana() int         { return p.mana }
func (p *Mago) Inteligencia() int { return p.inteligencia }
func (p *Mago) Forca() int        { return p.forca }
func (p *Mago) Level() int        { return p.level }
func (p *Mago) Xp() float32       { return p.xp }
