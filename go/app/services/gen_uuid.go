package services

import "github.com/google/uuid"

func GenUuid() string {
	id := uuid.New()
	return id.String()
}
