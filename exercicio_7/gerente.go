package funcionario

type Gerente Funcionario

func (self *Gerente) Bonificacao() float64 { return self.Salario + 10000 }
