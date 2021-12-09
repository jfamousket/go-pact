package pact

import "bytes"

func MakeExpression(programName string, args ...string) string {
	var b bytes.Buffer
	b.WriteString("(")
	b.WriteString(programName)
	b.WriteString(" ")
	for i, arg := range args {
		b.WriteString(arg)
		if i != len(args)-1 {
			b.WriteString(" ")
		}
	}
	return b.String()
}
