package file

import "io/ioutil"

// GetFileBytesPtr reads the file at filePath, and returns the contents as a byte slice.
func GetFileBytesPtr(filePath string) (*[]byte, error) {
	fileBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return &fileBytes, nil
}
