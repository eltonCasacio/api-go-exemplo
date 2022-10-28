package database

import (
	"testing"

	"github.com/eltoncasacio/api-go/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreate(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.User{})
	user, _ := entity.NewUser("elton", "elton@mail.com", "123")
	userRepository := NewUserRepository(db)
	err = userRepository.Create(user)
	assert.Nil(t, err)

	var userFound entity.User
	err = db.First(&userFound, "id = ?", user.ID).Error
	assert.Nil(t, err)
	assert.Equal(t, user.Email, userFound.Email)
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Name, userFound.Name)
	assert.NotNil(t, userFound.Password)
}

func TestFindByEmail(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.User{})
	user, _ := entity.NewUser("elton", "elton@mail.com", "123")
	userRepository := NewUserRepository(db)
	userRepository.Create(user)

	userFound, err := userRepository.FindByEmail("elton@mail.com")
	assert.Nil(t, err)
	assert.Equal(t, user.Email, userFound.Email)
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Name, userFound.Name)
	assert.NotNil(t, userFound.Password)
}

func TestFindByEmailIsNotFound(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.User{})
	user, _ := entity.NewUser("elton", "elton@mail.com", "123")
	userRepository := NewUserRepository(db)
	userRepository.Create(user)

	_, err = userRepository.FindByEmail("robert@mail.com")
	assert.NotNil(t, err)
}
