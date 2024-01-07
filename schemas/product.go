package schemas

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	name  string
	price float32
}
