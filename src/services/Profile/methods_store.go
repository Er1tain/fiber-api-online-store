package profile

import (
	"crypto/md5"
	"encoding/hex"
	"log"
	"os"
)

// save — метод для сохранения объекта в хранилище
func (s *Storage) Save(key string, data []byte) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	//Название(ключ) файла – это хэш-код email`а пользователя
	key_byte := []byte(key)
	md5Hash := md5.Sum(key_byte)
	key = hex.EncodeToString(md5Hash[:]) + ".jpg" //Добавим расширение к файлу

	// Сохраняем данные в памяти
	s.files[key] = data

	// Также сохраняем данные на диск
	err := os.WriteFile(STORAGE_DIR+"/"+key, data, 0644)
	if err != nil {
		log.Printf("Ошибка при сохранении файла %s: %v", key, err)
		return false
	}
	return true
}

// load — метод для загрузки объекта из хранилища
func (s *Storage) Load(key string) ([]byte, bool) {
	s.mu.Lock()         // Захватываем мьютекс перед чтением
	defer s.mu.Unlock() // Освобождаем мьютекс после чтения

	// Проверяем наличие объекта в памяти
	data, exists := s.files[key]
	if exists {
		return data, true
	}

	// Если объект не найден в памяти, пытаемся загрузить его с диска
	data, err := os.ReadFile(STORAGE_DIR + "/" + key)
	if err != nil {
		return nil, false
	}

	// Если загрузка с диска успешна, кэшируем объект в памяти
	s.files[key] = data
	return data, true
}
