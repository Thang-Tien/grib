package entity

var DriversTableName = "drivers"

type Driver struct {
	User
	VehicleID uint
	Vehicle   Vehicle
	WalletID  uint
	Wallet    Wallet
}

func (*Driver) TableName() string {
	return DriversTableName
}
