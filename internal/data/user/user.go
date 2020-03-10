package user

import (
	"context"
	"log"

	"go-tutorial-2020/pkg/errors"
	firebaseclient "go-tutorial-2020/pkg/firebaseClient"

	userEntity "go-tutorial-2020/internal/entity/user"

	"cloud.google.com/go/firestore"
	"github.com/jmoiron/sqlx"
	"google.golang.org/api/iterator"
)

type (
	// Data ...
	Data struct {
		db   *sqlx.DB
		fb   *firestore.Client
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
func New(db *sqlx.DB, fb *firebaseclient.Client) Data {
	d := Data{
		db: db,
		fb: fb.Client,
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

//GetUserFromFireBase ...
func (d Data) GetUserFromFireBase(ctx context.Context) ([]userEntity.User, error) {
	var (
		userFirebase []userEntity.User
		err          error
	)
	// test := d.fb.Collection("user_test")
	iter := d.fb.Collection("user_test").Documents(ctx)
	for {
		var user userEntity.User
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		log.Println(doc)
		err = doc.DataTo(&user)
		log.Println(user)
		userFirebase = append(userFirebase, user)
	}
	return userFirebase, err
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

//InsertUsersToFirebase ...
func (d Data) InsertUsersToFirebase(ctx context.Context, user userEntity.User) error {
	_, err := d.fb.Collection("user_test").Doc(user.NIP).Set(ctx, user)

	return err
}

//InsertMany ...
func (d Data) InsertMany(ctx context.Context, userList []userEntity.User) error {
	var (
		err error
	)
	for _, i := range userList {
		_, err = d.fb.Collection("user_test").Doc(i.NIP).Set(ctx, i)
	}
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
func (d Data) DeleteUserByNIP(ctx context.Context, NIP string) error {
	_, err := d.stmt[deleteUserByNIP].ExecContext(ctx,

		NIP)

	return err

}

//InsertNipUp ...
func (d Data) InsertNipUp(ctx context.Context) (int, error) {
	var nipMax int
	err := d.stmt[insertNipUp].QueryRowxContext(ctx).Scan(&nipMax)
	return nipMax, err
}

//UpdateByNipFirebase ...
func (d Data) UpdateByNipFirebase(ctx context.Context, nip string, user userEntity.User) error {
	iter,err := d.fb.Collection("user_test").Doc(nip).Get(ctx)
	userValidate := iter.Data()
	if userValidate == nil {
		return errors.Wrap(err, "Data Not Exist")
	}
	_, err = d.fb.Collection("user_test").Doc(nip).Set(ctx, user)
	return err
}

// DeleteByNipFirebase ...
func(d Data) DeleteByNipFirebase(ctx context.Context, nip string) error {
	iter, err := d.fb.Collection("user_test").Doc(nip).Get(ctx)
	userValidate := iter.Data()
	if userValidate == nil {
		return errors.Wrap(err, "Data Not Exist")
	}
	_, err = d.fb.Collection("user_test").Doc(nip).Delete(ctx)
	return err
	//test
}
// DeleteAllFirebase ...
// func(d Data) DeleteAllFirebase(ctx context.Context) error {
// 	err := d.fb.Collection("user_test").Get(ctx)
// 	userValidate := iter.Data()
// 	if userValidate == nil {
// 		return errors.Wrap(err, "Data Not Exist")
// 	}
// 	_, err = d.fb.Collection("user_test").Delete(ctx)
// 	return err
// }
