package uuid

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func NewID() string {
	id := primitive.NewObjectID()
	return id.Hex()
}

func NewUuidFromString(value string) (string, error) {
	id, err := primitive.ObjectIDFromHex(value)
	if err != nil {
		return "", err
	}
	return id.String(), nil
}

func NewGoogleUuid() string {
	return uuid.New().String()
}

func GoogleUuidParse(value string) (string, error) {
	id, err := uuid.Parse(value)
	if err != nil {
		return "", err
	}

	return id.String(), nil
}
