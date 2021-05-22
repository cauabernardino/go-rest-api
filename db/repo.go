package db

// Repo represents which methods the database will handle
type Repo interface {
	Create(model interface{}) (string, error)
}
