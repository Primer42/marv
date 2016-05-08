package marv

import (
	"fmt"
	"regexp"
	"strconv"
)

type ParseRule struct {
	name  string
	regex regexp.Regexp
}

type Parser struct {
	rules []ParseRule
}

func NewParser() *Parser {
	p := new(Parser)
	id := regexp.MustCompile("([[:lower:]][[:alnum:]]*)")
	typ := regexp.MustCompile("([[:upper:]][[:alnum:]]*)")
	litint := regexp.MustCompile("([[:digit:]]+)")
	expr := regexp.MustCompile("(.*)")

	parse_patterns := map[string][]*regexp.Regexp{
		"Vardef": []*regexp.Regexp{id, regexp.MustCompile(" :: "), typ},
		"Assign": []*regexp.Regexp{id, regexp.MustCompile(" = "), expr},
		"Type":   []*regexp.Regexp{typ},
		"Id":     []*regexp.Regexp{id},
		"LitInt": []*regexp.Regexp{litint},
		"Plus":   []*regexp.Regexp{expr, regexp.MustCompile(" \\+ "), expr},
		"Minus":  []*regexp.Regexp{expr, regexp.MustCompile(" - "), expr},
		"Neg":    []*regexp.Regexp{regexp.MustCompile("-"), expr},
	}

	for name, pattern_parts := range parse_patterns {
		regexp_string := "^"
		for _, pp := range pattern_parts {
			regexp_string = regexp_string + pp.String()
		}
		regexp_string = regexp_string + "$"
		reg := regexp.MustCompile(regexp_string)
		p.rules = append(p.rules, ParseRule{name, *reg})
	}

	return p
}

func (p Parser) Parse(lines []Line) (res []Smt, err error) {
	res = make([]Smt, 0)
	for _, line := range lines {
		smt, err := p.ParseSmt(line.content)
		if err != nil {
			return res, err
		}
		res = append(res, smt)
		if err != nil {
			return res, err
		}
	}
	return

}

func (p Parser) ParseSmt(s string) (Smt, error) {
	var res Smt
	for _, rule := range p.rules {
		if rule.regex.MatchString(s) {
			c := rule.regex.FindStringSubmatch(s)
			switch rule.name {
			case "Vardef":
				return p.Vardef(c)
			case "Assign":
				return p.Assign(c)
			}
			return res, fmt.Errorf("Line '%s' did not match statement - matched '%s' instead", s, rule.name)
		}
	}
	return res, fmt.Errorf("Line '%s' did not match any rules!", s)
}

func (p Parser) Vardef(cap []string) (Smt, error) {
	var res Smt
	e, err := p.ParseExp(cap[1])
	if err != nil {
		return res, err
	}
	t := cap[2]
	return Vardef{e, Typ{t}}, nil
}

func (p Parser) Assign(cap []string) (Smt, error) {
	var res Smt
	e1, err := p.ParseExp(cap[1])
	if err != nil {
		return res, err
	}
	e2, err := p.ParseExp(cap[2])
	if err != nil {
		return res, err
	}
	return Assign{e1, e2}, nil
}

func (p Parser) ParseExp(s string) (Expr, error) {
	var res Expr
	for _, rule := range p.rules {
		if rule.regex.MatchString(s) {
			c := rule.regex.FindStringSubmatch(s)
			switch rule.name {
			case "Id":
				return p.Id(c)
			case "LitInt":
				return p.LitInt(c)
			case "Neg":
				return p.Neg(c)
			case "Plus":
				return p.Plus(c)
			case "Minus":
				return p.Minus(c)
			}
			return res, fmt.Errorf("Statement '%s' did not match any expression - matched '%s' instead", s, rule.name)
		}
	}
	return res, fmt.Errorf("Statement %s did not match any rules!", s)
}

func (p Parser) Id(cap []string) (Expr, error) {
	return Id{cap[1]}, nil
}

func (p Parser) LitInt(cap []string) (Expr, error) {
	var res Expr
	l, err := strconv.Atoi(cap[1])
	if err != nil {
		return res, err
	}
	return LitInt{l}, nil
}

func (p Parser) Neg(cap []string) (Expr, error) {
	var res Expr
	e, err := p.ParseExp(cap[1])
	if err != nil {
		return res, err
	}
	return Neg{e}, nil
}

func (p Parser) Plus(cap []string) (Expr, error) {
	var res Expr
	e1, err := p.ParseExp(cap[1])
	if err != nil {
		return res, err
	}
	e2, err := p.ParseExp(cap[2])
	if err != nil {
		return res, err
	}
	return Plus{e1, e2}, nil
}

func (p Parser) Minus(cap []string) (Expr, error) {
	var res Expr
	e1, err := p.ParseExp(cap[1])
	if err != nil {
		return res, err
	}
	e2, err := p.ParseExp(cap[2])
	if err != nil {
		return res, err
	}
	return Minus{e1, e2}, nil
}
