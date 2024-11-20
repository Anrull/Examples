package models

import (
	"context"
	"Examples/BaseProject/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"github.com/go-redis/redis/v8"
)

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string
	Email    string
	Password string
}

var (
	DB    *gorm.DB
	Redis *redis.Client
	ctx   = context.Background()
)

func New(cfg *config.Config) {
	// Подключение к PostgreSQL
	dsn := "host=" + cfg.Database.Host + " user=" + cfg.Database.User +
		" password=" + cfg.Database.Password + " dbname=" + cfg.Database.DbName +
		" port=" + cfg.Database.Port + " sslmode=" + cfg.Database.SSLMode +
		" TimeZone=" + cfg.Database.Timezone
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	if err != nil {
		panic("Не удалось подключиться к базе данных")
	}

	// Миграция схемы
	DB.AutoMigrate(&User {})

	// Подключение к Redis
	// Redis = redis.NewClient(&redis.Options{
	// 	Addr:     cfg.Redis.Host, // Например, "localhost:6379"
	// 	Password: cfg.Redis.Password, // Если есть пароль
	// 	DB:       cfg.Redis.DB, // Номер базы данных
	// })
}

func AddUser (name, email, password string) {
	var user User
	DB.Where("email = ?", email).First(&user)
	if user.ID == 0 {
		DB.Create(&User {Name: name, Email: email, Password: password})
		// Добавление пользователя в Redis
		// err := Redis.Set(ctx, email, name, 0).Err() // Сохраняем имя пользователя по email
		// if err != nil {
		// 	panic("Не удалось сохранить пользователя в Redis")
		// }
	}
}

func GetUsers() []User  {
	var users []User 
	DB.Find(&users)
	return users
}


// func GetUserFromCache(email string) (string, error) {
// 	name, err := Redis.Get(ctx, email).Result()
// 	if err != nil {
// 		if err == redis.Nil {
// 			return "", nil // Ключ не найден
// 		}
// 		return "", err // Ошибка при получении
// 	}
// 	return name, nil // Возвращаем имя пользователя
// }

