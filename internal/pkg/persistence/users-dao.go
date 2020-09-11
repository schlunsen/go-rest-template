package persistence

import (
	"github.com/antonioalfa22/go-rest-template/internal/pkg/db"
	models "github.com/antonioalfa22/go-rest-template/internal/pkg/models/users"
	"strconv"
)

// UserDAO persists user data in database
type UserDAO struct{}

// NewUserDAO creates a new UserDAO
func NewUserDAO() *UserDAO {
	return &UserDAO{}
}

func (dao *UserDAO) Get(id string) (*models.User, error) {
	var user models.User
	where := models.User{}
	where.ID, _ = strconv.ParseUint(id, 10, 64)
	_, err := First(&where, &user, []string{"Role"})
	if err != nil {
		return nil, err
	}
	return &user, err
}

func (dao *UserDAO) GetByUsername(username string) (*models.User, error) {
	var user models.User
	where := models.User{}
	where.Username = username
	_, err := First(&where, &user, []string{"Role"})
	if err != nil {
		return nil, err
	}
	return &user, err
}

func (dao *UserDAO) All() (*[]models.User, error) {
	var users []models.User
	err := Find(&models.User{}, &users, []string{"Role"}, "id asc")
	return &users, err
}

func (dao *UserDAO) Query(q *models.User) (*[]models.User, error) {
	var users []models.User
	err := Find(&q, &users, []string{"Role"}, "id asc")
	return &users, err
}

func (dao *UserDAO) Add(user *models.User) error {
	err := Create(&user)
	err = Save(&user)
	return err
}

func (dao *UserDAO) Update(user *models.User) error {
	var userRole models.UserRole
	_, err := First(models.UserRole{UserID: user.ID}, &userRole, []string{})
	userRole.RoleName = user.Role.RoleName
	err = Save(&userRole)
	err = db.GetDB().Omit("Role").Save(&user).Error
	user.Role = userRole
	return err
}

func (dao *UserDAO) Delete(user *models.User) error {
	err := db.GetDB().Unscoped().Delete(models.UserRole{UserID: user.ID}).Error
	err = db.GetDB().Unscoped().Delete(&user).Error
	return err
}
