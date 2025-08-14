package storage

type FileStorage interface {
	Save(filename string, data []byte) (string, error) // returns public URL/path
}
