package skeleton

import (
	"context"

	userEntity "go-tutorial-2020/internal/entity/user"
	"go-tutorial-2020/pkg/errors"
)

// UserData ...
type UserData interface {
	// GetAllUsers(ctx context.Context) ([]userEntity.User, error)
	InsertUsersToFirebase(ctx context.Context, user userEntity.User) error
	GetUserFromFireBase(ctx context.Context) ([]userEntity.User, error)
}

// Service ...
type Service struct {
	userData UserData
}

// New ...
func New(userData UserData) Service {
	return Service{
		userData: userData,
	}
}

// // GetAllUsers ...
// func (s Service) GetAllUsers(ctx context.Context) ([]userEntity.User, error) {
// 	// Panggil method GetAllUsers di data layer user
// 	users, err := s.userData.GetAllUsers(ctx)
// 	// Error handling
// 	if err != nil {
// 		return users, errors.Wrap(err, "[SERVICE][GetAllUsers]")
// 	}
// 	// Return users array
// 	return users, err
// }

//InsertUsersToFirebase ...
func (s Service) InsertUsersToFirebase(ctx context.Context, user userEntity.User) error {
	err := s.userData.InsertUsersToFirebase(ctx, user)
	if err != nil {
		return  errors.Wrap(err, "[SERVICE][InsertUsersToFirebase]")
	}
	return err
}

// GetUserFromFireBase ...
func (s Service) GetUserFromFireBase(ctx context.Context) ([]userEntity.User, error) {
	var user []userEntity.User
	//	user, err := s.GetUserFromFireBase(ctx)

	user, err := s.userData.GetUserFromFireBase(ctx)
	if err != nil {
		return  user, errors.Wrap(err, "[SERVICE][GetUserFromFireBase]")
	}
	return user, err
}
