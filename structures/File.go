package structures

type File struct {
	Id uint32
	Path string
	Md5 string
}

func (File) TableName() string {
	return "files"
}
