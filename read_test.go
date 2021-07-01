package gedworx

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)


func getTestFiles(path string) ([]string, error){

	var fileNames []string
	files, err := ioutil.ReadDir(path)
	if err != nil{
		return nil, err
	}
	for _, f := range  files {
		if !f.IsDir() {
			fileNames = append(fileNames, f.Name())
		}

	}
	return fileNames, nil

}

func TestReadAll(t *testing.T) {
	var testFiles []string
	var err error
	if testFiles, err  = getTestFiles("./testdata"); err != nil{
		t.Fatalf(err.Error())
	}



	a := assert.New(t)
	type args struct {
		filePath string
	}
	type test struct  {
		name    string
		args    args
		want    []byte
		wantErr bool
	}

	var tests []test
	for _, testFile := range testFiles{

		tests = append(tests, test{
			name: testFile,
			args: args{filePath: fmt.Sprintf("./testdata/%s", testFile)},
			want: []byte{},
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
