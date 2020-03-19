package arif

import (
	"context"
	userEntity "tugas-arif/internal/entity/arif"

	"tugas-arif/pkg/errors"
)

// UserData ...
// Masukkan function dari package data ke dalam interface ini
type UserData interface {
	UpdateUserByID(ctx context.Context, ID string, user userEntity.User) (userEntity.User, error)
	DeleteUserByID(ctx context.Context, ID string) error
	InsertUser(ctx context.Context, user userEntity.User) error
	GetAllUser(ctx context.Context) ([]userEntity.User, error)
	GetUserByID(ctx context.Context, ID string) (userEntity.User, error)
	GetPrintPage(ctx context.Context, page int, length int) ([]userEntity.User, error)
}

// Service ...
// Tambahkan variable sesuai banyak data layer yang dibutuhkan
type Service struct {
	userData UserData
}

// New ...
// Tambahkan parameter sesuai banyak data layer yang dibutuhkan
func New(userdata UserData) Service {
	// Assign variable dari parameter ke object
	return Service{
		userData: userdata,
	}
}

// GetAllUser ...
func (s Service) GetAllUser(ctx context.Context) ([]userEntity.User, error) {
	// Panggil method GetAllUsers di data layer user
	users, err := s.userData.GetAllUser(ctx)
	// Error handling
	if err != nil {
		return users, errors.Wrap(err, "[SERVICE][GetAllUser]")
	}
	// Return users array
	return users, err
}

// InsertUser ...
func (s Service) InsertUser(ctx context.Context, user userEntity.User) error {
	// Panggil method GetAllUsers di data layer user
	err := s.userData.InsertUser(ctx, user)
	// Error handling
	if err != nil {
		return errors.Wrap(err, "[SERVICE][GetAllUsers]")
	}
	// Return users array
	return err
}

// GetUserByID ...
func (s Service) GetUserByID(ctx context.Context, ID string) (userEntity.User, error) {
	users, err := s.userData.GetUserByID(ctx, ID)
	if err != nil {
		return users, errors.Wrap(err, "SALAH")
	}
	return users, err
}

// DeleteUserByID ...
func (s Service) DeleteUserByID(ctx context.Context, ID string) error {
	err := s.userData.DeleteUserByID(ctx, ID)
	if err != nil {
		return errors.Wrap(err, "SALAH")
	}
	return err
}

// UpdateUserByID ...
func (s Service) UpdateUserByID(ctx context.Context, ID string, user userEntity.User) (userEntity.User, error) {
	user, err := s.userData.UpdateUserByID(ctx, ID, user)
	if err != nil {
		return user, errors.Wrap(err, "SALAH")
	}
	return user, err
}

// GetPrintPage ...
func (s Service) GetPrintPage(ctx context.Context, page int, length int) ([]userEntity.User, error) {
	userList, err := s.userData.GetPrintPage(ctx, page, length)
	if err != nil {
		return userList, errors.Wrap(err, "SALAH")
	}
	return userList, err
}
