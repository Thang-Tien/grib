package entity

type Customer struct {
	User
	WalletID uint `gorm:"column:wallet_id"`
	Wallet
}