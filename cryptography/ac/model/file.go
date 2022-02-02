package model

type File struct {
	FileName string
	Level    string
}

func NewFile(name, level string) *File {
	return &File{FileName: name, Level: level}
}
