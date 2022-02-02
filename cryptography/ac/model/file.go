package model

type File struct {
	FileName  string
	LevelBIBA string
	LevelBell string
}

func NewFile(name, bell, biba string) *File {
	return &File{FileName: name, LevelBIBA: biba, LevelBell: bell}
}
