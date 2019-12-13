package usecase

import (
	"testing"

	"github.com/maiyama18/gomock-playground2/model"
	mockrepository "github.com/maiyama18/gomock-playground2/repository/mock"

	"github.com/golang/mock/gomock"
)

const (
	dummyID   uint64 = 1
	dummyName        = "Updated"
)

func TestPersonUsecase_ChangeName(t *testing.T) {
	type mock struct {
		person        *model.Person
		findPersonErr error

		savePersonErr error
	}
	type want struct {
		err error
	}

	setup := func(t *testing.T, mock mock) (*gomock.Controller, *PersonUsecase) {
		ctrl := gomock.NewController(t)

		mockPersonRepository := mockrepository.NewMockPersonRepository(ctrl)

		personUsecase := &PersonUsecase{
			personRepository: mockPersonRepository,
		}

		mockPersonRepository.
			EXPECT().
			Find(dummyID).
			Return(mock.person, mock.findPersonErr).
			AnyTimes()

		mockPersonRepository.
			EXPECT().
			Save(gomock.Eq(&model.Person{ID: dummyID, Name: dummyName})).
			// ↑をSave(gomock.Eq(&model.Person{ID: dummyID, Name: "WRONG NAME"}))とかにするとちゃんとエラーになる
			Return(mock.savePersonErr).
			Times(1)

		return ctrl, personUsecase
	}

	tests := []struct {
		name string
		mock mock
		want want
	}{
		{
			name: "success",
			mock: mock{
				person: &model.Person{
					ID:   dummyID,
					Name: "Original",
				},
			},
			want: want{},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl, personUsecase := setup(t, test.mock)
			defer ctrl.Finish()

			err := personUsecase.ChangeName(dummyID, dummyName)
			if test.want.err == nil {
				if err != nil {
					t.Fatalf("err is expected to be nil, but got %q", err)
				}
			} else {
				if err == nil || err.Error() != test.want.err.Error() {
					t.Fatalf("err is expected to be %q, but got %q", test.want.err, err)
				}
			}
		})
	}
}
