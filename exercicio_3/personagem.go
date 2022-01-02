package personagem

type Personagem struct {
	nome         string
	vida         int
	mana         int
	inteligencia int
	forca        int
	level        int
	xp           float32
}

var QntPersonagensCriados = 0

func NewPersonagem(nome string, vida, mana, inteligencia, forca int) Personagem {
	p := Personagem{
		nome:         nome,
		vida:         vida,
		mana:         mana,
		inteligencia: inteligencia,
		forca:        forca,
		level:        1,
		xp:           0,
	}

	return p
}

func (p *Personagem) lvlUp() {
}

func (p *Personagem) Nome() string      { return p.nome }
func (p *Personagem) Vida() int         { return p.vida }
func (p *Personagem) Mana() int         { return p.mana }
func (p *Personagem) Inteligencia() int { return p.inteligencia }
func (p *Personagem) Forca() int        { return p.forca }
func (p *Personagem) Level() int        { return p.level }
func (p *Personagem) Xp() float32       { return p.xp }
