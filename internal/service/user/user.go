package skeleton

import (
	"context"
	"fmt"
	"net/http"

	userEntity "go-tutorial-2020/internal/entity/user"
	"go-tutorial-2020/pkg/errors"
	"go-tutorial-2020/pkg/kafka"
	"log"
)

// UserData ...
type UserData interface {
	InsertUsers(ctx context.Context, user userEntity.User) error
	GetAllUsers(ctx context.Context) ([]userEntity.User, error)
	GetUserByNIP(ctx context.Context, NIP string) (userEntity.User, error)
	UpdateUserByNIP(ctx context.Context, NIP string, user userEntity.User) (userEntity.User, error)
	DeleteUserByNIP(ctx context.Context, NIP string) error
	InsertNipUp(ctx context.Context) (int, error)
	GetUserFromFireBase(ctx context.Context) ([]userEntity.User, error)
	InsertUsersToFirebase(ctx context.Context, user userEntity.User, nipMaxi int) error
	InsertMany(ctx context.Context, userList []userEntity.User) error
	UpdateByNipFirebase(ctx context.Context, nip string, user userEntity.User) error
	DeleteByNipFirebase(ctx context.Context, nip string) error
	//DeleteAllFirebase(ctx context.Context) error
	NipIncrement(ctx context.Context) (int, error)
	GetUserClient(ctx context.Context, headers http.Header) ([]userEntity.User, error)
	InsertUserClient(ctx context.Context, headers http.Header, user userEntity.User) error
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
	nipMaxi, err := s.userData.InsertNipUp(ctx)
	user.NIP = "P" + fmt.Sprintf("%06d", nipMaxi)
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
func (s Service) DeleteUserByNIP(ctx context.Context, NIP string) error {
	// Panggil method GetAllUsers di data layer user
	err := s.userData.DeleteUserByNIP(ctx, NIP)
	// Error handling
	if err != nil {
		return errors.Wrap(err, "[SERVICE][GetAllUsers]")
	}
	// Return users array
	return err
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
	nipMaxi, err := s.userData.NipIncrement(ctx)
	nipMaxi = nipMaxi + 1
	user.NIP = "P" + fmt.Sprintf("%06d", nipMaxi)
	err = s.userData.InsertUsersToFirebase(ctx, user, nipMaxi)
	log.Println(user.NIP, user.ID, user.Email, user.Nama, user.Jabatan, user.TanggalLahir)
	if err != nil {
		return errors.Wrap(err, "[SERVICE][InsertNipIncrementFirebase]")
	}
	return err
}

//InsertMany ...
func (s Service) InsertMany(ctx context.Context, userList []userEntity.User) error {
	err := s.userData.InsertMany(ctx, userList)
	return err
}

//PublishUser ...
func (s Service) PublishUser(user userEntity.User) error {
	err := s.kafka.SendMessageJSON("New_User", user)
	if err != nil {
		return errors.Wrap(err, "[SERVICE][publishRO]")
	}
	return err
}

//UpdateByNipFirebase ...
func (s Service) UpdateByNipFirebase(ctx context.Context, nip string, user userEntity.User) error {
	err := s.userData.UpdateByNipFirebase(ctx, nip, user)
	return err
}

// DeleteByNipFirebase ...
func (s Service) DeleteByNipFirebase(ctx context.Context, nip string) error {
	err := s.userData.DeleteByNipFirebase(ctx, nip)
	return err
}

// DeleteAllFirebase ...
// func(s Service) DeleteAllFirebase(ctx context.Context) error{
// 	err := s.userData.DeleteAllFirebase(ctx)
// 	return err
// }

// GetUserClient ...
func (s Service) GetUserClient(ctx context.Context, headers http.Header) ([]userEntity.User, error) {
	UserClient, err := s.userData.GetUserClient(ctx, headers)

	return UserClient, err
}

//InsertUserClient ...
func (s Service) InsertUserClient(ctx context.Context, headers http.Header, user userEntity.User) error {
	err := s.userData.InsertUserClient(ctx, headers, user)
	if err != nil {
		return errors.Wrap(err, "[SERVICE][InsertUserClient]")
	}
	return err
}
