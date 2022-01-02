package funcionario

type Funcionario struct {
	Nome    string
	Idade   int
	Salario float64
}

func (self *Funcionario) Bonificacao() float64 { return self.Salario }
