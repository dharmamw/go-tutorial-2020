package user

import (
	"context"
	userEntity "go-tutorial-2020/internal/entity/user"
	"log"

	"go-tutorial-2020/pkg/errors"
	firebaseclient "go-tutorial-2020/pkg/firebaseClient"

	"cloud.google.com/go/firestore"
	"github.com/jmoiron/sqlx"
	"google.golang.org/api/iterator"
)

type (
	// Data ...
	Data struct {
		// db   *sqlx.DB
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
)

var (
	readStmt = []statement{
		{getAllUsers, qGetAllUsers},
	}
)

// New ...
func New(fb *firebaseclient.Client) Data {
	d := Data{
		// db: db,
		fb: fb.Client,
	}

	// d.initStmt()
	return d
}

// func (d *Data) initStmt() {
// 	var (
// 		err   error
// 		stmts = make(map[string]*sqlx.Stmt)
// 	)

// 	for _, v := range readStmt {
// 		stmts[v.key], err = d.db.PreparexContext(context.Background(), v.query)
// 		if err != nil {
// 			log.Fatalf("[DB] Failed to initialize statement key %v, err : %v", v.key, err)
// 		}
// 	}

// 	d.stmt = stmts
// }

// // GetAllUsers digunakan untuk mengambil semua data user
// func (d Data) GetAllUsers(ctx context.Context) ([]userEntity.User, error) {
// 	var (
// 		user  userEntity.User
// 		users []userEntity.User
// 		err   error
// 	)

// 	// Query ke database
// 	rows, err := d.stmt[getAllUsers].QueryxContext(ctx)

// 	// Looping seluruh row data
// 	for rows.Next() {
// 		// Insert row data ke struct user
// 		if err := rows.StructScan(&user); err != nil {
// 			return users, errors.Wrap(err, "[DATA][GetAllUsers] ")
// 		}
// 		// Tambahkan struct user ke array user
// 		users = append(users, user)
// 	}
// 	// Return users array
// 	return users, err
// }

// InsertUsersToFirebase ...
func (d Data) InsertUsersToFirebase(ctx context.Context, user userEntity.User) error {
	_, err := d.fb.Collection("user_test").Doc(user.NIP).Set(ctx, user)
	// if err = rows.StructScan(&user); err != nil {
	// 	return errors.Wrap(err, "[DATA][InsertUsersToFirebase] ")
	// }
	return err
}

//GetUserFromFireBase ...
func (d Data) GetUserFromFireBase(ctx context.Context) ([]userEntity.User, error) {
	var (
		userFirebase []userEntity.User
		err          error
	)
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

// UpdateUserFromFirebase ...
func (d Data) UpdateUserFromFirebase(ctx context.Context, nip string, user userEntity.User) error {
	iter, err := d.fb.Collection("user_test").Doc(nip).Get(ctx)
	userValidate := iter.Data()
	if userValidate == nil {
		return errors.Wrap(err, "Data Not Exist")
	}
	_, err = d.fb.Collection("user_test").Doc(nip).Set(ctx, user)
	return err
}
