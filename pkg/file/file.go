package file

import "io/ioutil"

func GetFileBytesPtr(filePath string) (*[]byte, error) {
	fileBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return &fileBytes, nil
}
