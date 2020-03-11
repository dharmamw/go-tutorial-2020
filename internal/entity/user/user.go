package user

import "time"

// JAVA EQUIVALENT -> MODEL

// User object model
type User struct {
	ID           int       `db:"id" json:"user_id"`
	NIP          string    `db:"nip" json:"nip"`
	Nama         string    `db:"nama_lengkap" json:"nama_lengkap"`
	Jabatan      string    `db:"jabatan" json:"jabatan"`
	TanggalLahir time.Time `db:"tanggal_lahir" json:"tanggal_lahir"`
	Email        string    `db:"email" json:"email"`
}
