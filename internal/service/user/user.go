package skeleton

import (
	"context"
	"fmt"
	userEntity "go-tutorial-2020/internal/entity/user"
	"go-tutorial-2020/pkg/errors"
	"go-tutorial-2020/pkg/kafka"
	"net/http"
)

// UserData ...
type UserData interface {
	InsertUsers(ctx context.Context, user userEntity.User) error
	GetAllUsers(ctx context.Context) ([]userEntity.User, error)
	GetUserByNIP(ctx context.Context, NIP string) (userEntity.User, error)
	UpdateUserByNIP(ctx context.Context, NIP string, user userEntity.User) (userEntity.User, error)
	DeleteUserByNIP(ctx context.Context, NIP string, user userEntity.User) (userEntity.User, error)
	InsertNipUp(ctx context.Context) (int, error)
	GetUserFromFireBase(ctx context.Context) ([]userEntity.User, error)
	InsertUsersToFirebase(ctx context.Context, user userEntity.User) error
	InsertMany(ctx context.Context, userList []userEntity.User) error
	GetAllUsersAPI(ctx context.Context, header http.Header) ([]userEntity.User, error)
}

// Service ...
type Service struct {
	userData UserData
	kafka    *kafka.Kafka
}

// New ...
func New(userData UserData, kafka *kafka.Kafka) Service {
	return Service{
		userData: userData,
		kafka:    kafka,
	}
}

//InsertUsers ...
func (s Service) InsertUsers(ctx context.Context, user userEntity.User) error {
	// Panggil method GetAllUsers di data layer user
	nipMax, err := s.userData.InsertNipUp(ctx)
	user.NIP = "P" + fmt.Sprintf("%06d", nipMax)
	err = s.userData.InsertUsers(ctx, user)
	// Error handling
	if err != nil {
		return errors.Wrap(err, "[SERVICE][GetAllUsers]")
	}
	// Return users array
	return err
}
func (s Service) InsertUsersSqlFirebase(ctx context.Context, user userEntity.User) error {
	// Panggil method GetAllUsers di data layer user
	nipMax, err := s.userData.InsertNipUp(ctx)
	user.NIP = "P" + fmt.Sprintf("%06d", nipMax)
	err = s.userData.InsertUsersToFirebase(ctx, user)
	err = s.userData.InsertUsers(ctx, user)
	// Error handling
	if err != nil {
		return errors.Wrap(err, "[SERVICE][GetAllUsers]")
	}
	// Return users array
	return err
}

// GetAllUsers ...
func (s Service) GetAllUsers(ctx context.Context) ([]userEntity.User, error) {
	// Panggil method GetAllUsers di data layer user
	users, err := s.userData.GetAllUsers(ctx)
	// Error handling
	if err != nil {
		return users, errors.Wrap(err, "[SERVICE][GetAllUsers]")
	}
	// Return users array
	return users, err
}

//isi interface

//GetUserByNIP ...
func (s Service) GetUserByNIP(ctx context.Context, NIP string) (userEntity.User, error) {
	users, err := s.userData.GetUserByNIP(ctx, NIP)
	if err != nil {
		return users, errors.Wrap(err, "SALAH")
	}
	return users, err

}

//UpdateUserByNIP ...
func (s Service) UpdateUserByNIP(ctx context.Context, NIP string, user userEntity.User) (userEntity.User, error) {
	// Panggil method GetAllUsers di data layer user
	user, err := s.userData.UpdateUserByNIP(ctx, NIP, user)
	// Error handling
	if err != nil {
		return user, errors.Wrap(err, "[SERVICE][GetAllUsers]")
	}
	// Return users array
	return user, err
}

// DeleteUserByNIP ...
func (s Service) DeleteUserByNIP(ctx context.Context, NIP string, user userEntity.User) (userEntity.User, error) {
	// Panggil method GetAllUsers di data layer user
	user, err := s.userData.DeleteUserByNIP(ctx, NIP, user)
	// Error handling
	if err != nil {
		return user, errors.Wrap(err, "[SERVICE][GetAllUsers]")
	}
	// Return users array
	return user, err
}

// GetUserFromFireBase ...
func (s Service) GetUserFromFireBase(ctx context.Context) ([]userEntity.User, error) {
	var user []userEntity.User
	//	user, err := s.GetUserFromFireBase(ctx)

	user, err := s.userData.GetUserFromFireBase(ctx)
	return user, err
}

//InsertUsersToFirebase ...
func (s Service) InsertUsersToFirebase(ctx context.Context, user userEntity.User) error {
	err := s.userData.InsertUsersToFirebase(ctx, user)
	return err
}

//InsertMany ...
func (s Service) InsertMany(ctx context.Context, userList []userEntity.User) error {
	err := s.userData.InsertMany(ctx, userList)
	return err
}

func (s Service) PublishUser(user userEntity.User) error {
	err := s.kafka.SendMessageJSON("New_User", user)
	if err != nil {
		return errors.Wrap(err, "[SERVICE][publishRO]")
	}
	return err
}

func (s Service) GetAllUsersAPI(ctx context.Context, header http.Header) ([]userEntity.User, error) {
	userList, err := s.userData.GetAllUsersAPI(ctx, header)
	return userList, err
}
