package profile

import (
	"sync"
)

const (
	STORAGE_DIR         = "src/services/Profile/store_pict" // ДИРЕКТОРИЯ ДЛЯ ХРАНЕНИЯ АВАТАРОК
	UPLOAD_PREFIX_LEN   = len("/upload/")                   // ДЛИНА ПРЕФИКСА ДЛЯ МАРШРУТА ВЫГРУЗКИ
	DOWNLOAD_PREFIX_LEN = len("/download/")                 // ДЛИНА ПРЕФИКСА ДЛЯ МАРШРУТА ЗАГРУЗКИ
)

// Storage — структура для хранения объектов
type Storage struct {
	mu    sync.Mutex        // Мьютекс для обеспечения потокобезопасности
	files map[string][]byte // Хэш-таблица для хранения данных объектов
}

var storage *Storage

// NewStorage — конструктор для создания(в единственном экземпляре  singleton) нового хранилища
func NewStorage() *Storage {
	if storage == nil {
		storage = &Storage{
			files: make(map[string][]byte),
		}
	}
	return storage
}
