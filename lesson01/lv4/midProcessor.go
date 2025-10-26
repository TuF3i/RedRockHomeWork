package main

func initProcessor() Processor {
	return Processor{
		s1:     initStack(),
		s2:     initStack(),
		opList: []string{"+", "-", "*", "/", "(", ")"},
	}
}

type Processor struct {
	s1     stack
	s2     stack
	opList []string
}

func (p *Processor) handle(f []string) {

	for i := 0; i < len(f); i++ {
		sig := string(f[i])
		p.logic(sig)
	}

	for !p.s1.ifEmpty() {
		peek, _ := p.s1.pop()
		p.s2.push(peek)
	}
}

func (p *Processor) getWeight(op string) int {
	switch op {
	case "+", "-":
		return 1
	case "*", "/":
		return 2
	default:
		return 0
	}
}

func (p *Processor) isNum(item string) bool {
	for _, tar := range p.opList {
		if item == tar {
			return false
		}
	}
	return true
}

func (p *Processor) logic(sig string) {
	if p.isNum(sig) {
		p.s2.push(sig)
		return
	}

	if p.s1.ifEmpty() {
		p.s1.push(sig)
		return
	}

	if sig == "(" {
		p.s1.push(sig)
		return
	}

	if sig == ")" {
		for !p.s1.ifEmpty() {
			if peek, _ := p.s1.getPeek(); peek == "(" {
				_, _ = p.s1.pop()
				break
			}
			top, _ := p.s1.pop()
			p.s2.push(top)
		}
		return
	}

	if top, _ := p.s1.getPeek(); p.getWeight(top) > p.getWeight(sig) {
		for {
			if peek, _ := p.s1.getPeek(); p.getWeight(peek) < p.getWeight(sig) || p.s1.ifEmpty() {
				p.s1.push(sig)
				break
			}
			top, _ := p.s1.pop()
			p.s2.push(top)
		}
		return
	}

	p.s1.push(sig)
	return
}
