package arif

import (
	"context"
	"log"

	userEntity "tugas-arif/internal/entity/arif"
	"tugas-arif/pkg/errors"

	"github.com/jmoiron/sqlx"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
	firebaseclient "tugas-arif/pkg/firebaseClient"
)

type (
	// Data ...
	Data struct {
		db   *sqlx.DB
		stmt map[string]*sqlx.Stmt
		fb   *firestore.Client
	}

	// statement ...
	statement struct {
		key   string
		query string
	}
)

// Tambahkan query di dalam const
// getAllUser = "GetAllUser"
// qGetAllUser = "SELECT * FROM users"
const (
	updateUserByID  = "UpdateUserByID"
	qUpdateUserByID = "UPDATE Arief SET Nama = ?, Umur = ?, `Tanggal Lahir` = ? WHERE ID = ?"

	deleteUserByID  = "DeleteUserByID"
	qDeleteUserByID = "DELETE FROM Arief WHERE ID = ?"

	getAllUser  = "GetAllUser"
	qGetAllUser = "SELECT * FROM Arief"

	insertUser  = "InsertUser"
	qInsertUser = "INSERT INTO Arief VALUES (?, ?, ?, ?)"

	getUserByID  = "GetUserByID"
	qGetUserByID = "SELECT * FROM Arief WHERE id = ?"
)

var (
	readStmt = []statement{
		{updateUserByID, qUpdateUserByID},
		{deleteUserByID, qDeleteUserByID},
		{getAllUser, qGetAllUser},
		{insertUser, qInsertUser},
		{getUserByID, qGetUserByID},
	}
)

// New func...
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

// UpdateUserByID ...
func (d Data) UpdateUserByID(ctx context.Context, ID string, user userEntity.User) (userEntity.User, error) {
	_, err := d.stmt[updateUserByID].ExecContext(ctx,
		user.Name,
		user.Age,
		user.DOB,
		ID)

	return user, err
}

// DeleteUserByID ...
func (d Data) DeleteUserByID(ctx context.Context, ID string) error {
	_, err := d.stmt[deleteUserByID].ExecContext(ctx, ID)

	return err
}

//GetAllUser . . . .
func (d Data) GetAllUser(ctx context.Context) ([]userEntity.User, error) {
	var (
		user  userEntity.User
		users []userEntity.User
		err   error
	)

	//Query ke database . s. .
	rows, err := d.stmt[getAllUser].QueryxContext(ctx)

	//Looping seluruh row data . . .
	for rows.Next() {
		//Insert row data ke struct user . . .
		if err := rows.StructScan(&user); err != nil {
			return users, errors.Wrap(err, "[DATA][GetAllUser]")
		}
		//Tambah struct user ke array user
		users = append(users, user)
	}
	//Return users array
	return users, err
}

//InsertUser ...
func (d Data) InsertUser(ctx context.Context, user userEntity.User) error {
	_, err := d.stmt[insertUser].ExecContext(ctx,

		user.ID,
		user.Name,
		user.Age,
		user.DOB)

	return err
}

//GetUserByID ...
func (d Data) GetUserByID(ctx context.Context, ID string) (userEntity.User, error) {
	var (
		user userEntity.User
		err  error
	)

	if err = d.stmt[getUserByID].QueryRowxContext(ctx, ID).StructScan(&user); err != nil {
		return user, errors.Wrap(err, "SALAH")
	}
	return user, err
}

// GetPrintPage ...
func (d Data) GetPrintPage(ctx context.Context, page int, length int) ([]userEntity.User, error) {
	var (
		user    userEntity.User
		users   []userEntity.User
		iter    *firestore.DocumentIterator
		lastDoc *firestore.DocumentSnapshot
		err     error
	)

	if page == 1 {
		// Kalau page 1 ambil data langsung dari query
		iter = d.fb.Collection("PrintUser").Limit(length).Documents(ctx)
	} else {
		// Kalau page > 1 ambil data sampai page sebelumnya
		previous := d.fb.Collection("PrintUser").Limit((page - 1) * length).Documents(ctx)
		docs, err := previous.GetAll()
		if err != nil {
			return nil, err
		}
		// Ambil doc terakhir
		lastDoc = docs[len(docs)-1]
		// Query mulai setelah doc terakhir
		iter = d.fb.Collection("PrintUser").StartAfter(lastDoc).Limit(length).Documents(ctx)
	}

	// Looping documents
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}

		if err != nil {
			return users, errors.Wrap(err, "[DATA][GetUserPage] Failed to iterate Document!")
		}
		err = doc.DataTo(&user)
		if err != nil {
			return users, errors.Wrap(err, "[DATA][GetUserPage] Failed to Populate Struct!")
		}
		users = append(users, user)
	}
	return users, err
}
