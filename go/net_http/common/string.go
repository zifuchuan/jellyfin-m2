package common

import (
	"strings"
	"unicode"
)

func String01(val string) string {
	return strings.Replace(val, "a", "b", -1)
}

func String02(val string) string {
	return strings.ReplaceAll(val, "a", "b")
}

func String03(val string) string {
	return strings.Join([]string{val}, ",")
}

func String04(val string) string {
	return strings.Split(val, ",")[0]
}

func String05(val string) string {
	val = strings.Title(val)
	val = strings.ToLower(val)
	val = strings.ToUpper(val)
	val = strings.ToTitle(val)
	val = strings.ToUpperSpecial(unicode.TurkishCase, val)
	val = strings.ToTitleSpecial(unicode.TurkishCase, val)
	val = strings.ToLowerSpecial(unicode.TurkishCase, val)
	return val
}

func String06(val string) string {
	val = strings.TrimSpace(val)
	val = strings.TrimLeft(val, " \t")
	val = strings.TrimRight(val, " \t")
	val = strings.TrimPrefix(val, "abc")
	val = strings.TrimSuffix(val, "abc")
	return val
}

func String07(val string) string {
	sb := strings.Builder{}
	sb.WriteString(val)
	val = sb.String()
	return val
}

func String08(val string) string {
	src := strings.Split(val, ":")
	dst := make([]string, 1)
	copy(dst, src)
	return strings.Join(dst, ":")
}
