package shop_models

import (
	"log"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

func connect() *gorm.DB {
	//Подключаемся к БД
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=ClothesShop password=userpass sslmode=disable")
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	return db
}

// Список товаров, доступных к покупке
func GetListClothes() []uuid.UUID {
	db := connect()
	defer db.Close()

	list_clothes := []uuid.UUID{}

	//Тут должно быть получение списка вещей....

	return list_clothes
}

// Есть ли товар в наличии?
func CheckThing(id_thing uuid.UUID) bool {
	db := connect()
	defer db.Close()

	//Тут должна быть проверка наличия вещи в продаже

	return false
}
