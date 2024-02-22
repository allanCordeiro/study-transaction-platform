package entity

type UserInterface interface {
	Save(user *User) error
	FindByMail(email string) (*User, error)
	FindByID(id string) (*User, error)
	Update(user *User) (*User, error)
	Delete(user *User) error
}
