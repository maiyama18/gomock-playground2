package usecase

import "github.com/maiyama18/gomock-playground2/repository"

type PersonUsecase struct {
	personRepository repository.PersonRepository
}

func (u *PersonUsecase) ChangeName(id uint64, name string) error {
	p, err := u.personRepository.Find(id)
	if err != nil {
		return err
	}

	p.Name = name
	return u.personRepository.Save(p)
}
