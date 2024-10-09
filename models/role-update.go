package models

type Update struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	RoleID   uint   `gorm:"not null" json:"role_id"`
}

func (update *Update) Save() (*Update, error) {
	return update, nil
}
