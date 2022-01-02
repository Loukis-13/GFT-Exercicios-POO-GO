package funcionario

type Supervisor Funcionario

func (self *Supervisor) Bonificacao() float64 { return self.Salario + 5000 }
