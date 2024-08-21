package entity

var VehiclesTableName = "vehicles"

type Vehicle struct {
	BaseEntity
	Model string `gorm:"column:model;not null"`
	Type string `gorm:"column:type;not null"`
	NumberPlate string `gorm:"column:number_plate;not null"`
}

func (*Vehicle) TableName() string {
	return VehiclesTableName
}