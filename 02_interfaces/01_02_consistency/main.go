package main

import (
	"database/sql"
	"github.com/sirupsen/logrus"
	"net/http"

	_ "github.com/lib/pq" // Подключение драйвера PostgreSQL
)

// User - структура для хранения данных пользователя
type User struct {
	ID   string
	Name string
}

// UserRepository - интерфейс для работы с пользователями
type UserRepository interface {
	Save(user User) error
	FindByID(id string) (User, error)
	Delete(id string) error
}

// DBUserRepository - реализация интерфейса для работы с базой данных
type DBUserRepository struct {
	db  *sql.DB
	log *logrus.Logger
}

func NewDBUserRepository(db *sql.DB, log *logrus.Logger) *DBUserRepository {
	return &DBUserRepository{
		db:  db,
		log: log}
}

func (repo *DBUserRepository) Save(user User) error {
	repo.log.Infof("Сохранение пользователя %s в базу данных\n", user.Name)

	return nil
}

func (repo *DBUserRepository) FindByID(id string) (User, error) {
	repo.log.Infof("Поиск пользователя с ID: %s в базе данных\n", id)

	return User{ID: id, Name: "Пользователь базы данных"}, nil
}

func (repo *DBUserRepository) Delete(id string) error {
	repo.log.Infof("Удаление пользователя с ID: %s из базы данных\n", id)

	return nil
}

// APIUserRepository - реализация интерфейса для работы через API
type APIUserRepository struct {
	log    *logrus.Logger
	client *http.Client
	apiURL string
}

func NewAPIUserRepository(client *http.Client, apiURL string, log *logrus.Logger) *APIUserRepository {
	return &APIUserRepository{
		client: client,
		apiURL: apiURL,
		log:    log,
	}
}

func (repo *APIUserRepository) Save(user User) error {
	repo.log.Infof("Сохранение пользователя %s через API\n", user.Name)

	return nil
}

func (repo *APIUserRepository) FindByID(id string) (User, error) {
	repo.log.Infof("Поиск пользователя с ID: %s через API\n", id)

	return User{ID: id, Name: "Пользователь API"}, nil
}

func (repo *APIUserRepository) Delete(id string) error {
	repo.log.Infof("Удаление пользователя с ID: %s через API\n", id)
	return nil
}

func manageUsers(repo UserRepository) {
	user := User{ID: "123", Name: "John Doe"}
	err := repo.Save(user)
	if err != nil {
		return
	}
	_, err = repo.FindByID("123")
	if err != nil {
		return
	}
	err = repo.Delete("123")
	if err != nil {
		return
	}
}

func main() {
	log := logrus.New()

	db, err := sql.Open("postgres", "user=postgres dbname=postgres sslmode=disable")
	if err != nil {
		panic("Ошибка подключения к базе данных")
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	dbRepo := NewDBUserRepository(db, log)
	apiRepo := NewAPIUserRepository(http.DefaultClient, "http://example.com/api", log)

	log.Infoln("Использование репозитория базы данных:")
	manageUsers(dbRepo)
	log.Infoln("Использование API репозитория:")
	manageUsers(apiRepo)
}
