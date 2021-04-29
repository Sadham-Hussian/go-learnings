package database

// ProductItem struct to describe product object
type ProductItem struct {
	ID    uint32 `gorm:"primaryKey;autoIncrement"`
	Name  string
	Type  string
	Prize int32
}
