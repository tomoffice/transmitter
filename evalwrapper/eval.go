package evalwapper

import (
	"bytes"
	"html/template"

	"github.com/antonmedv/expr"
)

// Eval is github.com/antonmedv/expr wapper
func Eval(vars map[string]interface{}, code string) (interface{}, error) {

	var buf bytes.Buffer

	t := template.Must(template.New("").Parse(code))
	err := t.Execute(&buf, vars)
	if err != nil {
		panic(err)
	}
	program, err := expr.Compile(buf.String(), expr.Env(vars))
	if err != nil {
		panic(err)
	}
	output, err := expr.Run(program, vars)
	if err != nil {
		panic(err)
	}
	var result interface{}
	switch output.(type) {
	case float64:
		result = output.(float64)
	case int:
		result = float64(output.(int))
	case string:
		result = output
	}
	return result, nil

}
