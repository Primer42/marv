package marv

type Line struct {
	content string
}

type Typ struct {
	name string
}

type Smt interface {
	isSmt() bool
}

type Expr interface {
	isExpr() bool
}

type Vardef struct {
	val Expr
	typ Typ
}

func (v Vardef) isSmt() bool {
	return true
}

type Assign struct {
	tar Expr
	val Expr
}

func (a Assign) isSmt() bool {
	return true
}

type Id struct {
	id string
}

func (i Id) isExpr() bool {
	return true
}

type LitInt struct {
	val int
}

func (l LitInt) isExpr() bool {
	return true
}

type Neg struct {
	tar Expr
}

func (n Neg) isExpr() bool {
	return true
}

type Plus struct {
	a Expr
	b Expr
}

func (p Plus) isExpr() bool {
	return true
}

type Minus struct {
	a Expr
	b Expr
}

func (m Minus) isExpr() bool {
	return true
}
