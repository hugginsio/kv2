package kms

type KeyManagementSystem interface {
	Retrieve(id string) ([]byte, error)
}
