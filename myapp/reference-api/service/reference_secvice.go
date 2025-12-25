package service

import (
	"github.com/practice-cloud-run-functions/myapp/reference-api/repository"
)

func FetchUsers() ([]map[string]any, error) {
	return repository.SelectUsers()
}
