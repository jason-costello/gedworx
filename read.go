package gedworx

import (
	"io/ioutil"
)

func ReadAll(filePath string) ([]byte, error) {

	 return ioutil.ReadFile(filePath)

}
