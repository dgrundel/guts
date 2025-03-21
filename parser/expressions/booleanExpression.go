package expressions

import (
	"bytes"
)

// Supported boolean operators:
type BooleanOperator string

const (
	Equal              BooleanOperator = "=="
	NotEqual           BooleanOperator = "!="
	GreaterThan        BooleanOperator = ">"
	LessThan           BooleanOperator = "<"
	GreaterThanOrEqual BooleanOperator = ">="
	LessThanOrEqual    BooleanOperator = "<="
	LogicalAnd         BooleanOperator = "&&"
	LogicalOr          BooleanOperator = "||"
	LogicalNot         BooleanOperator = "!"
)

var _booleanOperators = map[string]BooleanOperator{
	"==": Equal,
	"!=": NotEqual,
	">":  GreaterThan,
	"<":  LessThan,
	">=": GreaterThanOrEqual,
	"<=": LessThanOrEqual,
	"&&": LogicalAnd,
	"||": LogicalOr,
	"!":  LogicalNot,
}

type BooleanExpression interface {
	Left() BooleanExpression
	Right() BooleanExpression
	Operator() BooleanOperator
	Parentheses() bool
	Literal() string
	ExpressionType() ExpressionType
	String() string
}

type booleanExpression struct {
	left        BooleanExpression
	right       BooleanExpression
	operator    BooleanOperator
	literal     string
	typ         ExpressionType
	parentheses bool
}

func NewBooleanExpression(s string) (BooleanExpression, error) {
	expr, _, err := ParseBooleanExpression(s)
	return expr, err
}

func (b *booleanExpression) Left() BooleanExpression {
	return b.left
}

func (b *booleanExpression) Right() BooleanExpression {
	return b.right
}

func (b *booleanExpression) Operator() BooleanOperator {
	return b.operator
}

func (b *booleanExpression) Parentheses() bool {
	return b.parentheses
}

func (b *booleanExpression) Literal() string {
	return b.literal
}

func (b *booleanExpression) ExpressionType() ExpressionType {
	return b.typ
}

func (b *booleanExpression) String() string {
	if b.literal != "" {
		if b.typ != nil {
			return b.literal + ":" + b.typ.String()
		}

		return b.literal
	}

	var buf bytes.Buffer
	if b.parentheses {
		buf.WriteByte('(')
	}
	if b.left != nil {
		buf.WriteString(b.left.String())
		buf.WriteByte(' ')
	}
	buf.WriteString(string(b.operator))
	if b.right != nil {
		buf.WriteByte(' ')
		buf.WriteString(b.right.String())
	}
	if b.parentheses {
		buf.WriteByte(')')
	}
	return buf.String()
}
