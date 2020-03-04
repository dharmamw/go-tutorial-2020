package user

import (
	"context"
	"log"

	"go-tutorial-2020/pkg/errors"

	userEntity "go-tutorial-2020/internal/entity/user"

	"github.com/jmoiron/sqlx"
)

type (
	// Data ...
	Data struct {
		db   *sqlx.DB
		stmt map[string]*sqlx.Stmt
	}

	// statement ...
	statement struct {
		key   string
		query string
	}
)

const (
	getAllUsers  = "GetAllUsers"
	qGetAllUsers = "SELECT * FROM user_test"

	insertUsers  = "InsertUsers"
	qinsertUsers = "INSERT INTO user_test VALUES (NULL, ?, ?, ?, ?, ?)"

	getUserByNIP  = "GetAllUsersByNIP"
	qGetUserByNIP = "SELECT * FROM user_test WHERE nip = ?"

	updateUserByNIP  = "UpdateUserByNIP"
	qUpdateUserByNIP = "UPDATE user_test SET nip = ?, nama_lengkap = ?, tanggal_lahir = ?, jabatan = ?, email = ? WHERE nip = ?"

	deleteUserByNIP  = "DelteUserByNIP"
	qDeleteUserByNIP = "DELETE FROM user_test WHERE NIP = ?"

	insertNipUp  = "InsertNipUp"
	qInsertNipUp = "SELECT MAX(CAST(RIGHT(nip,6)AS INT))+1 FROM user_test"
)

var (
	readStmt = []statement{
		{getAllUsers, qGetAllUsers},
		{insertUsers, qinsertUsers},
		{getUserByNIP, qGetUserByNIP},
		{updateUserByNIP, qUpdateUserByNIP},
		{deleteUserByNIP, qDeleteUserByNIP},
		{insertNipUp, qInsertNipUp},
	}
)

// New ...
func New(db *sqlx.DB) Data {
	d := Data{
		db: db,
	}

	d.initStmt()
	return d
}

func (d *Data) initStmt() {
	var (
		err   error
		stmts = make(map[string]*sqlx.Stmt)
	)

	for _, v := range readStmt {
		stmts[v.key], err = d.db.PreparexContext(context.Background(), v.query)
		if err != nil {
			log.Fatalf("[DB] Failed to initialize statement key %v, err : %v", v.key, err)
		}
	}

	d.stmt = stmts
}

// GetAllUsers digunakan untuk mengambil semua data user
func (d Data) GetAllUsers(ctx context.Context) ([]userEntity.User, error) {
	var (
		user  userEntity.User
		users []userEntity.User
		err   error
	)

	// Query ke database
	rows, err := d.stmt[getAllUsers].QueryxContext(ctx)

	// Looping seluruh row data
	for rows.Next() {
		// Insert row data ke struct user
		if err := rows.StructScan(&user); err != nil {
			return users, errors.Wrap(err, "[DATA][GetAllUsers] ")
		}
		// Tambahkan struct user ke array user
		users = append(users, user)
	}
	// Return users array
	return users, err

}

//InsertUsers ...
func (d Data) InsertUsers(ctx context.Context, user userEntity.User) error {
	_, err := d.stmt[insertUsers].ExecContext(ctx,

		user.NIP,
		user.Nama,
		user.TanggalLahir,
		user.Jabatan,
		user.Email)

	return err

}

//user di if dipakai di var, jadi jangan bingung!

//GetUserByNIP ...
func (d Data) GetUserByNIP(ctx context.Context, NIP string) (userEntity.User, error) {
	var (
		user userEntity.User
		err  error
	)

	if err = d.stmt[getUserByNIP].QueryRowxContext(ctx, NIP).StructScan(&user); err != nil {
		return user, errors.Wrap(err, "SALAH")
	}
	return user, err
}

//UpdateUserByNIP ...
func (d Data) UpdateUserByNIP(ctx context.Context, NIP string, user userEntity.User) (userEntity.User, error) {
	_, err := d.stmt[updateUserByNIP].ExecContext(ctx,
		user.NIP,
		user.Nama,
		user.TanggalLahir,
		user.Jabatan,
		user.Email,
		NIP)

	return user, err

}

// DeleteUserByNIP ...
func (d Data) DeleteUserByNIP(ctx context.Context, NIP string, user userEntity.User) (userEntity.User, error) {
	_, err := d.stmt[deleteUserByNIP].ExecContext(ctx,

		NIP)

	return user, err

}

//InsertNipUp ...
func (d Data) InsertNipUp(ctx context.Context) (int, error) {
	var nipMax int
	err := d.stmt[insertNipUp].QueryRowxContext(ctx).Scan(&nipMax)
	return nipMax, err
}
