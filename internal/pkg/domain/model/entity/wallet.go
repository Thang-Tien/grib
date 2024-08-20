package entity

type Wallet struct {
	BaseEntity
	Balance int `gorm:"column:balance"`
	CreditCard string `gorm:"column:credit_card;not null"`
	LinkedBankAccount string `gorm:"linked_bank_account;not null"`
}