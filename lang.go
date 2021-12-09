package pact

import (
	"bytes"
	"strconv"
)

func MakeExpression(programName string, args ...interface{}) string {
	var b bytes.Buffer
	b.WriteString("(")
	b.WriteString(programName)
	if len(args) > 0 {
		b.WriteString(" ")
	}
	for i, arg := range args {
		switch a := arg.(type) {
		case string:
			b.WriteString("\"")
			b.WriteString(a)
			b.WriteString("\"")
		case int:
			b.WriteString(strconv.Itoa(a))
		}
		if i != len(args)-1 {
			b.WriteString(" ")
		}
	}
	b.WriteString(")")
	return b.String()
}
