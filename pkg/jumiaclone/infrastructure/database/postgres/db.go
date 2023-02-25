package postgres

import (
	"fmt"
	"log"

	"github.com/VicOsewe/Jumia-clone/pkg/jumiaclone/application"
	"github.com/VicOsewe/Jumia-clone/pkg/jumiaclone/domain/dto"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type JumiaDB struct {
	DB *gorm.DB
}

func NewJumiaDB() *JumiaDB {
	j := JumiaDB{
		DB: Init(),
	}
	j.checkPreConditions()
	return &j
}

func (j *JumiaDB) checkPreConditions() {
	if j.DB == nil {
		log.Panicf("database has not been initialized")
	}
}

func runMigrations(db *gorm.DB) {
	tables := []interface{}{}

	for _, table := range tables {
		if err := db.AutoMigrate(table); err != nil {
			log.Panicf("can't run migrations on table %v: %v", table, err)
		}
	}
}

func Init() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Africa/Nairobi",
		application.MustGetEnvVar("DB_HOST"),
		application.MustGetEnvVar("DB_USER"),
		application.MustGetEnvVar("DB_PASSWORD"),
		application.MustGetEnvVar("DB_NAME"),
		application.MustGetEnvVar("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("can't open connection to the local database: %v", err)
	}
	runMigrations(db)
	return db
}

func (db *JumiaDB) CreateUser(user *dto.User) (*dto.User, error) {
	if user == nil {
		return nil, fmt.Errorf("nil contact")
	}

	if err := db.DB.Create(user).Error; err != nil {
		return nil, fmt.Errorf(
			"can't create a new marketing record: err: %v",
			err,
		)
	}
	return user, nil
}
