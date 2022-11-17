package preprocessor

import (
	"errors"
	"fmt"
	"os"
	"path"
	"regexp"
	"strings"
)

type Preprocessor struct {
	Input  string
	Output string
}

func New(input string) *Preprocessor {
	pp := &Preprocessor{
		Input:  input,
		Output: "",
	}
	return pp
}

func (pp *Preprocessor) ResolveImports() {
	reImportStatements := regexp.MustCompile(`import\s+\([\s\n]*("\w+",[\s\n]*)+\)`)
	reImportLibraries := regexp.MustCompile(`\w+`)

	importStmtIndices := reImportStatements.FindAllStringIndex(pp.Input, -1)
	if len(importStmtIndices) == 0 {
		pp.Output = pp.Input
		return
	}

	srcArr := []string{}

	// for each block of imports using their indices
	for _, indices := range importStmtIndices {
		srcCode := ""
		importStr := pp.Input[indices[0]:indices[1]]
		importStr = strings.Replace(importStr, "import", "", -1)
		libNames := reImportLibraries.FindAllString(importStr, -1)

		// for each library name in current import block
		for _, lib := range libNames {
			s, err := readImportFile(lib)
			if err != nil {
				fmt.Println(err)
			} else {
				srcCode += s + "\n"
			}
		}
		srcArr = append(srcArr, srcCode)
	}
	if len(importStmtIndices) != len(srcArr) {
		pp.Output = pp.Input
		return
	}

	lastIdx := 0
	for i, indices := range importStmtIndices {
		pp.Output += pp.Input[lastIdx:indices[0]]
		pp.Output += srcArr[i]
		lastIdx = indices[1]
	}
	pp.Output += pp.Input[lastIdx:]
}

func readImportFile(filepath string) (string, error) {
	usrLibPath, err := createUserLibPath(filepath)
	if err != nil {
		fmt.Println(err)
	}

	if exists, _ := fileExists(usrLibPath); exists {
		s, err := readFileString(usrLibPath)
		return s, err
	}

	stdLibPath, err := createStdLibPath(filepath)
	if err != nil {
		fmt.Println(err)
	}

	if exists, _ := fileExists(stdLibPath); exists {
		s, err := readFileString(stdLibPath)
		return s, err
	}

	return "", errors.New("unable to read import file")
}

func fileExists(filepath string) (bool, error) {
	if _, err := os.Stat(filepath); err == nil {
		return true, nil

	} else if errors.Is(err, os.ErrNotExist) {
		return false, err

	} else {
		fmt.Println(err)
		return false, err
	}
}

func readFileString(filepath string) (string, error) {
	b, err := os.ReadFile(filepath)
	return string(b), err
}

func createUserLibPath(library string) (string, error) {
	workingDir, err := os.Getwd()
	absolutePath := path.Join(workingDir, library+".emp")
	return absolutePath, err
}

func createStdLibPath(library string) (string, error) {
	binPath, err := os.Executable()
	stdLibPath := path.Join(path.Dir(path.Dir(binPath)), "stdlib", library+".emp")
	return stdLibPath, err
}
