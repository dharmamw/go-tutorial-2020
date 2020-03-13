package user

import (
	"time"

	"gopkg.in/guregu/null.v3/zero"
)

// JAVA EQUIVALENT -> MODEL

// User object model
type User struct {
	ID           int       `db:"id" json:"user_id"`
	NIP          string    `db:"nip" json:"nip"`
	Nama         string    `db:"nama_lengkap" json:"nama_lengkap"`
	TanggalLahir time.Time `db:"tanggal_lahir" json:"tanggal_lahir"`
	Jabatan      string    `db:"jabatan" json:"jabatan"`
	Email        string    `db:"email" json:"email"`
}

//DataResp Get Data Resp ...
type DataResp struct {
	Data     []User      `json:"data"`
	Metadata interface{} `json:"metadata"`
	Error    interface{} `json:"error"`
}

// Transf object model ...
type Transf struct {
	RunningID     zero.Int  `db:"TransfD_RunningID" json: "running_id"`
	OutCode       string    `db:"TransfD_OutCodeTransf" json: "out_code"`
	NoTransf      string    `db:"TransfD_NoTransf" json: "no_transf"`
	Group         int       `db:"TransfD_Group" json: "group"`
	OutCodeSP     string    `db:"TransfD_OutCodeSP" json: "out_code_sp"`
	Wilayah       zero.Int  `db:"TransfD_Wilayah" json: "wilayah"`
	NoSP          string    `db:"TransfD_NoSP" json: "no_sp"`
	ProCod        string    `db:"TransfD_ProCod" json: "pro_cod"`
	Quantity      zero.Int  `db:"TransfD_Qty" json: "qty"`
	QuantityScan  zero.Int  `db:"TransfD_Qty_Scan" json: "qty_scan"`
	QuantityStock zero.Int  `db:"TransfD_QtyStk" json: "qty_stk"`
	OutCodeOrder  string    `db:"TransfD_OutCodeOrder" json: "out_code_order"`
	NoOrder       string    `db:"TransfD_NoOrder" json: "no_order"`
	ActiveYN      string    `db:"TransfD_ActiveYN" json: "active_yn"`
	UserID        string    `db:"TransfD_UserID" json: "user_id"`
	LastUpdate    time.Time `db:"TransfD_LastUpdate" json: "last_update"`
	DataActiveYN  string    `db:"TransfD_DataAktifYN" json: "data_active_yn"`
}