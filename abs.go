package marv

type Line struct {
	content string
}

type Typ struct {
	name string
}

type Smt struct{}
type Expr struct{}

type Vardef struct {
	Smt
	val Expr
	typ Typ
}

type Assign struct {
	Smt
	tar Expr
	val Expr
}

type Id struct {
	Expr
	id string
}

type LitInt struct {
	Expr
	val int
}

type Neg struct {
	Expr
	tar Expr
}

type Plus struct {
	Expr
	a Expr
	b Expr
}

type Minus struct {
	Expr
	a Expr
	b Expr
}
