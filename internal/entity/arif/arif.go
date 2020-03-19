package arif

import "time"


type User struct {
	ID   string    `db:"ID" json:"id" firestore:"ID"`
	Name string    `db:"Nama" json:"name" firestore:"Nama"`
	Age  int       `db:"Umur" json:"age" firestore:"Umur"`
	DOB  time.Time `db:"Tanggal Lahir" json:"dob" firestore:"Tanggal Lahir"`
}
