package entity

var CustumersTableName = "customers"

type Customer struct {
	User
	WalletID uint
	Wallet   Wallet
}

func (*Customer) TableName() string {
	return CustumersTableName
}
