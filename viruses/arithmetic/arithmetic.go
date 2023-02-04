package arithmetic

import (
	"go/ast"
	"go/token"

	"github.com/gtramontina/ooze/viruses"
)

type Arithmetic struct {
	mutations map[token.Token]token.Token
}

func New() *Arithmetic {
	return &Arithmetic{
		mutations: map[token.Token]token.Token{
			token.ADD: token.SUB,
			token.SUB: token.ADD,
			token.MUL: token.QUO,
			token.QUO: token.MUL,
			token.REM: token.MUL,
		},
	}
}

func (v *Arithmetic) Incubate(node ast.Node) []*viruses.Infection {
	expression, matches := node.(*ast.BinaryExpr)
	if !matches {
		return nil
	}

	originalOperation := expression.Op

	mutatedOperation, matches := v.mutations[expression.Op]
	if !matches {
		return nil
	}

	return []*viruses.Infection{
		viruses.NewInfection(
			"Arithmetic",
			func() { expression.Op = mutatedOperation },
			func() { expression.Op = originalOperation },
		),
	}
}
