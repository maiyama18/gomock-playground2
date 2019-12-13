//go:generate mockgen -source=$GOFILE -destination=mock/mock_$GOFILE -package=mockrepository
package repository

import "github.com/maiyama18/gomock-playground2/model"

type PersonRepository interface {
	Find(id uint64) (*model.Person, error)
	Save(person *model.Person) error
}
