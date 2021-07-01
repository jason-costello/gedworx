package gedworx

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const BannedABNF = `banned = %x00-08 / %x0B-0C / %x0E-1F ; C0 other than LF CR and Tab
       / %x7F                        ; DEL
       / %x80-9F                     ; C1
       / %xD800-DFFF                 ; Surrogates
`

func ReadAll(filePath string) ([]byte, error) {

	return ioutil.ReadFile(filePath)

}

func Parse(input []byte) error {

	// An artificial input source.

	scanner := bufio.NewScanner(strings.NewReader(string(input)))

	// Read initial byte and check for BOM (EF BB BF)
	// GEDCOM v7-rc1 UTF-8 only acceptable encoding

	lineCounter := 0
	for scanner.Scan() {
		// fmt.Println("scanning line: ", lineCounter)
		temp := scanner.Text()
		if lineCounter == 0 {
			if !HasUTF8BOM(temp) {
				return errors.New("no utf8 BOM, unable to parse as v7-rc1")
			}
		}
		lineCounter++
		// parse rest of file as needed

	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "should be no error scanning string")
	}

	return nil

}

func HasUTF8BOM(firstLine string) bool {

	if len(firstLine) < 3 {
		return false
	}

	if []byte(firstLine)[0] == 239 &&
		[]byte(firstLine)[1] == 187 &&
		[]byte(firstLine)[2] == 191 {

		return true

	}

	return false

}
