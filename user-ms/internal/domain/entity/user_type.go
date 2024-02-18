package entity

type UserType struct {
	Value string
}

func NewUserType(name string) *UserType {
	return &UserType{Value: name}
}
