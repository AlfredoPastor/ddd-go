package vo

import (
	"errors"

	"github.com/AlfredoPastor/ddd-go/shared/uuid"
)

type ID struct {
	value string
}

var ErrorInvalidID = errors.New("invalid id")

func NewIDFromString(value string) (ID, error) {
	_, err := uuid.NewUuidFromString(value)
	if err != nil {
		return ID{}, ErrorInvalidID
	}
	return ID{
		value: value,
	}, nil
}

func NewID() ID {
	return ID{
		value: uuid.NewID(),
	}
}

func (v ID) String() string {
	return v.value
}
