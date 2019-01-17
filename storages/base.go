// Package storages allows multiple implementation on how to store short URLs.
package storages

type IStorage interface {
	Code() string
	Save(string) string
	Load(string) (string, error)
}

// a type for holding (database) credentials
type DbCredentials struct {
	Host string
	Port int
	Name int
	User string
	Pass string
}
