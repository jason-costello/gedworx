package gedworx

import (
	"bufio"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"strings"
	"testing"
)

func getTestFiles(path string) ([]string, error) {

	var fileNames []string
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	for _, f := range files {
		if !f.IsDir() {
			fileNames = append(fileNames, f.Name())
		}

	}
	return fileNames, nil

}

func TestReadAll(t *testing.T) {
	var testFiles []string
	var err error
	if testFiles, err = getTestFiles("./testdata"); err != nil {
		t.Fatalf(err.Error())
	}

	a := assert.New(t)
	type args struct {
		filePath string
	}
	type test struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}

	var tests []test
	for _, testFile := range testFiles {

		tests = append(tests, test{
			name:    testFile,
			args:    args{filePath: fmt.Sprintf("./testdata/%s", testFile)},
			want:    []byte{},
			wantErr: false,
		})

	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadAll(tt.args.filePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			a.IsType(tt.want, got)
		})
	}
}

func TestParse(t *testing.T) {
	a := assert.New(t)

	var testFiles []string
	var err error
	if testFiles, err = getTestFiles("./testdata"); err != nil {
		t.Fatalf(err.Error())
	}

	type args struct {
		input []byte
	}
	type test struct {
		name    string
		args    args
		wantErr bool
	}
	var tests []test

	for _, testFile := range testFiles {

		tests = append(tests, test{
			name:    fmt.Sprintf("./testdata/%s", testFile),
			args:    args{},
			wantErr: !strings.Contains(testFile, "utf8"),
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.input, err = ReadAll(tt.name)
			if err != nil {
				t.Fail()
				return
			}
			err = Parse(tt.args.input)
			if err != nil {
				if a.Error(err) != tt.wantErr {
					t.Fail()
					return
				}
			}

			// a.NotEqual(tt.want, a.NoError(err ))

		})
	}
}
func Test_HasUTF8BOM(t *testing.T) {
	a := assert.New(t)
	var testFiles []string
	var err error
	if testFiles, err = getTestFiles("./testdata"); err != nil {
		t.Fatalf(err.Error())
	}

	type args struct {
		firstLine []byte
	}
	type test struct {
		name string
		args args
		want bool
	}
	var tests []test

	for _, testFile := range testFiles {

		tests = append(tests, test{
			name: fmt.Sprintf("./testdata/%s", testFile),
			args: args{},
			want: strings.Contains(testFile, "utf8"),
		})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			input, err := ReadAll(tt.name)
			if err != nil {
				t.Log(fmt.Sprintf("%s:  %s ", tt.name, err.Error()))
				t.Fail()
				return
			}
			scanner := bufio.NewScanner(strings.NewReader(string(input)))

			for scanner.Scan() {
				temp := scanner.Text()

				a.Equal(tt.want, HasUTF8BOM(temp))
				break
			}

		})
	}
}
