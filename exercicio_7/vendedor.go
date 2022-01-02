package funcionario

type Vendedor Funcionario

func (self *Vendedor) Bonificacao() float64 { return self.Salario + 3000 }
