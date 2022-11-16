package repl

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/user"
	"path"
	"runtime"
	"strings"

	"github.com/justingodden/empire/pkg/evaluator"
	"github.com/justingodden/empire/pkg/lexer"
	"github.com/justingodden/empire/pkg/object"
	"github.com/justingodden/empire/pkg/parser"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	u, err := user.Current()
	if err != nil {
		panic(err)
	}
	banner(out)
	fmt.Printf("Hello %s! Welcome to the Empire progamming Language.\n", u.Username)
	fmt.Println("Let's get started 🚀")

	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Fprint(out, PROMPT)
		input := ""
		multiLine := true
		newLine := false

		for multiLine {
			if newLine {
				fmt.Fprint(out, ".. ")
			}

			scanned := scanner.Scan()
			if !scanned {
				return
			}

			line := scanner.Text()
			input += " " + line

			lineRStrip := strings.TrimRight(line, " \n\t\r")
			lastChar := string(lineRStrip[len(lineRStrip)-1])
			multiLine = !(lastChar == ";")
			newLine = multiLine
		}

		l := lexer.New(input)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}

func banner(out io.Writer) {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		panic("Could not find Caller of repl.ReadFile")
	}
	absolutePath := path.Join(path.Dir(filename), "ascii.txt")
	b, err := os.ReadFile(absolutePath)
	if err != nil {
		panic(err)
	}
	fmt.Fprint(out, string(b))
}
