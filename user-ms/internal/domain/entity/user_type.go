package entity

type UserType uint8

const (
	Customer UserType = iota + 1
	Distributor
)

func (u UserType) EnumIndex() uint8 {
	return uint8(u)
}

func (u UserType) String() string {
	switch u {
	case Customer:
		return "customer"
	case Distributor:
		return "distributor"
	}
	return "unknow"
}
