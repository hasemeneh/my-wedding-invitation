package models

type GuestModels struct {
	ID         int64  `db:"id" json:"id"`
	GuestName  string `db:"guest_name" json:"guest_name"`
	GuestGroup string `db:"guest_group" json:"guest_group"`
	MaxGuest   int64  `db:"max_guest" json:"max_guest"`
	Sides      int64  `db:"sides" json:"sides"`
	GuestKey   string `db:"guest_key" json:"guest_key"`
}
