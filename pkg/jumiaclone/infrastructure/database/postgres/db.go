package postgres

import (
	"fmt"
	"log"

	"github.com/VicOsewe/Jumia-clone/pkg/jumiaclone/application"
	"github.com/VicOsewe/Jumia-clone/pkg/jumiaclone/domain/dao"

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
	tables := []interface{}{
		&dao.OTPPayload{},
		&dao.UserProfile{},
	}

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

func (db *JumiaDB) CreateUser(user *dao.UserProfile) (*dao.UserProfile, error) {
	if user == nil {
		return nil, fmt.Errorf("nil user")
	}

	if err := db.DB.Create(user).Error; err != nil {
		return nil, fmt.Errorf(
			"can't create a new user record: err: %v",
			err,
		)
	}
	return user, nil
}

func (db *JumiaDB) GetUserByPhoneNumber(phoneNumber string) (*dao.UserProfile, error) {

	user := dao.UserProfile{}

	if err := db.DB.Where(
		&dao.UserProfile{
			PhoneNumber: phoneNumber,
		}).
		Find(&user).
		Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (db *JumiaDB) GetUserByEmail(email string) (*dao.UserProfile, error) {
	user := dao.UserProfile{}
	if err := db.DB.Where(
		&dao.UserProfile{
			Email: email,
		}).
		Find(&user).
		Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (db *JumiaDB) UpdateUser(user *dao.UserProfile) (*dao.UserProfile, error) {
	if err := db.DB.Where(
		&dao.UserProfile{
			PhoneNumber: user.PhoneNumber,
		}).
		Updates(user).
		Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (db *JumiaDB) SaveOTP(otp *dao.OTPPayload) error {
	if otp == nil {
		return fmt.Errorf("nil otp")
	}

	if err := db.DB.Create(otp).Error; err != nil {
		return fmt.Errorf("can't save otp record: %v", err)
	}
	return nil
}

func (db *JumiaDB) GetOTP(phoneNumber, otp string) (*dao.OTPPayload, error) {
	OTPPayload := dao.OTPPayload{}
	if err := db.DB.Where(&dao.OTPPayload{PhoneNumber: phoneNumber, OTPPassword: otp}).Find(&OTPPayload).Error; err != nil {
		return nil, err
	}
	return &OTPPayload, nil
}

func (db *JumiaDB) UpdateOTP(otp *dao.OTPPayload) (*dao.OTPPayload, error) {
	if err := db.DB.Where(
		&dao.OTPPayload{
			PhoneNumber: otp.PhoneNumber,
			IsValid:     true,
		}).
		Updates(otp).
		Error; err != nil {
		return nil, err
	}
	return otp, nil
}
