package generator

import "strings"

type SymbolType int

const (
	Type SymbolType = iota
	StructField
	EnumValue
)

type SymTab struct {
	Generated map[string]string
	Types     map[string]SymbolType
}

func (s *SymTab) GetSymbol(str string, t SymbolType) string {
	vg, ok1 := s.Generated[str]
	vt, ok2 := s.Types[str]

	if ok1 && ok2 && vt == t {
		return vg
	}

	formatted := s.formatSymbol(str)

	s.Generated[str] = formatted
	s.Types[str] = t
	return formatted
}

func (s *SymTab) formatSymbol(str string) string {
	str = strings.ReplaceAll(str, "\"", "")
	str = strings.ReplaceAll(str, "'", "")

	chars := make([]rune, len(str))

	for i, c := range str {
		if s.isLetter(c) || s.isNumber(c) || s.isAllowedSymbol(c) {
			chars[i] = c
		} else {
			chars[i] = '_'
		}
	}

	reducedChars := []rune{}

	var lastChar rune

	for _, c := range chars {
		if !(c == lastChar && c == '_') {
			reducedChars = append(reducedChars, c)
		}
		lastChar = c
	}

	str = strings.ToUpper(string(reducedChars[0])) + string(reducedChars[1:])
	return strings.Trim(str, "_ \n\r\t")
}

func (s *SymTab) isLetter(c rune) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

func (s *SymTab) isNumber(c rune) bool {
	return c >= '0' && c <= '9'
}

func (s *SymTab) isAllowedSymbol(c rune) bool {
	return c == '_'
}
