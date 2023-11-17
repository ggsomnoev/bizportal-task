package controller

import "os"

// FileInput represents an input source from a file.
type FileInput struct {
	File     *os.File
	FileName string
}

// NewFileInput creates a new FileInput with the given file name.
func (c *TailMovementController) NewFileInput(fileName string) (*FileInput, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	return &FileInput{File: file, FileName: fileName}, nil
}

// Name returns the file name associated with the FileInput.
func (fi *FileInput) Name() string {
	return fi.FileName
}

// Read reads data from the file associated with the FileInput.
func (fi *FileInput) Read(p []byte) (n int, err error) {
	return fi.File.Read(p)
}
