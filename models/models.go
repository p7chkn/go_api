package models

import "github.com/jackc/pgtype"

type User struct {
	LAST_LOGIN   pgtype.Timestamp `json:"last_login"`
	ID           pgtype.UUID      `json:"id"`
	PHONE_NUMBER pgtype.BPChar    `json:"phone_number"`
	EMAIL        pgtype.BPChar    `json:"email"`
	PASSWORD     pgtype.BPChar    `json:"paswwod"`
	LAST_CODE    pgtype.Int4      `json:"last_code"`
	IS_STAFF     pgtype.Bool      `json:"is_staff"`
	IS_ACTIVE    pgtype.Bool      `json:"is_active"`
	CREATED_AT   pgtype.Timestamp `json:"created_at"`
	UPDATED_AT   pgtype.Timestamp `json:"updated_at"`
}
