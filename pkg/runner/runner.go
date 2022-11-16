package runner

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/justingodden/empire/pkg/evaluator"
	"github.com/justingodden/empire/pkg/lexer"
	"github.com/justingodden/empire/pkg/object"
	"github.com/justingodden/empire/pkg/parser"
)

func Run(filePath string) {
	if filepath.Ext(filePath) != ".emp" {
		panic("Empire can only read .emp files.")
	}

	input, err := readFile(filePath)
	if err != nil {
		panic(err)
	}

	l := lexer.New(input)
	p := parser.New(l)
	env := object.NewEnvironment()

	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		for _, msg := range p.Errors() {
			fmt.Println(msg)
		}
	}

	evaluated := evaluator.Eval(program, env)
	if evaluated != nil {
		fmt.Println(evaluated.Inspect())
	}
}

func readFile(pathFromCaller string) (string, error) {
	workingDir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	absolutePath := path.Join(workingDir, pathFromCaller)
	f, err := os.ReadFile(absolutePath)
	str := string(f[:])
	return str, err
}
