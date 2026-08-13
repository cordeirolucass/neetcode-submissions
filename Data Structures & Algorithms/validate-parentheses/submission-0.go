type Pilha struct {
	elementos[]rune
}

func (p *Pilha) isEmpty() bool {
	return len(p.elementos) == 0
}

func (p *Pilha) Push(ch rune) {
	p.elementos = append(p.elementos,ch)
}

func (p *Pilha) Pop() (rune, error) {
	if p.isEmpty() {
		return 0, nil
	}
	indice := len(p.elementos)-1
	ch := p.elementos[indice]
	p.elementos = p.elementos[:indice]
	return ch, nil
}

func isValid(s string) bool {
	pilha := Pilha{}
	for _, ch := range s {
		if ch == '(' || ch == '[' || ch == '{' {
			pilha.Push(ch)
		} else if ch == ')' || ch == ']' || ch == '}' {
			if pilha.isEmpty() { return false }
			topo, _ := pilha.Pop()
			if (ch == ')' && topo != '(') ||
				(ch == ']' && topo != '[') ||
				(ch == '}' && topo != '{') { return false  }
		}
	} 
	return pilha.isEmpty()
}
