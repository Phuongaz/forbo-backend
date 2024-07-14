package models

import (
	"errors"
	"strconv"

	"github.com/phuongaz/forbo/common"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type FollowData struct {
	ID         string `json:"id"`
	FollowerID string `json:"follower_id"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Follower struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
}

type UserModel struct {
	gorm.Model
	UserID       string     `json:"user_id" gorm:"primary_key"`
	Username     string     `json:"username" gorm:"not null;unique"`
	PasswordHash string     `json:"password" gorm:"not null"`
	Email        string     `json:"email" gorm:"size:255;not null;unique"`
	Avatar       string     `json:"avatar"`
	Followers    []Follower `json:"followers" gorm:"type:json"`
	Following    []Follower `json:"following" gorm:"type:json"`
	Role         string     `json:"role" gorm:"default:user"`
	CreateAt     string     `json:"create_at" gorm:"default:CURRENT_TIMESTAMP"`
}

func (u *UserModel) Follow(userID string) {
	u.Following = append(u.Following, Follower{UserID: userID})
}

func (u *UserModel) IsAdmin() bool {
	return u.Role == "admin"
}

func (u *UserModel) Unfollow(userID string) {
	for i, f := range u.Following {
		if f.UserID == userID {
			u.Following = append(u.Following[:i], u.Following[i+1:]...)
			break
		}
	}
}

func (u *UserModel) GetFollowers() []string {
	followers := make([]string, 0, len(u.Followers))
	for _, f := range u.Followers {
		userID, _ := strconv.Atoi(f.UserID)
		followers = append(followers, strconv.Itoa(userID))
	}
	return followers
}

func (u *UserModel) GetFollowing() []string {
	following := make([]string, 0, len(u.Following))
	for _, key := range u.Following {
		userID, _ := strconv.Atoi(key.UserID)
		following = append(following, strconv.Itoa(userID))
	}
	return following
}

func (u *UserModel) IsFollowing(userID string) bool {
	for _, follower := range u.Following {
		if follower.UserID == userID {
			return true
		}
	}
	return false
}

func (u *UserModel) IsFollowedBy(userID string) bool {
	for _, follower := range u.Followers {
		if follower.UserID == userID {
			return true
		}
	}
	return false
}

func (u *UserModel) AddFollower(userID string) {
	if u.IsFollowedBy(userID) {
		return
	}

	follower := Follower{
		UserID:   userID,
		Username: userID,
	}
	u.Followers = append(u.Followers, follower)
}

func (u *UserModel) RemoveFollower(userID string) {
	for i, follower := range u.Followers {
		if follower.UserID == userID {
			u.Followers = append(u.Followers[:i], u.Followers[i+1:]...)
			break
		}
	}
}

func (u *UserModel) SetPassword(password string) error {
	if len(password) == 0 {
		return errors.New("password should not be empty")
	}
	bytePassword := []byte(password)
	passwordHash, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	u.PasswordHash = string(passwordHash)
	return nil
}

func (u *UserModel) CheckPassword(password string) error {
	bytePassword := []byte(password)
	byteHashedPassword := []byte(u.PasswordHash)
	return bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
}

func (u *UserModel) Create() error {
	result := common.SQLDBUser.Create(u)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *UserModel) Update() error {
	result := common.SQLDBUser.Save(u)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (ulg *UserLogin) Register() (*UserModel, error) {
	user := &UserModel{
		UserID:    common.GenerateUID(ulg.Email),
		Username:  ulg.Email,
		Email:     ulg.Email,
		Avatar:    "https://www.gravatar.com/avatar/205e460b479e2e5b48aec07710c08d50",
		Followers: make([]Follower, 0),
		Following: make([]Follower, 0),
	}
	if err := user.SetPassword(ulg.Password); err != nil {
		return nil, err
	}
	if err := user.Create(); err != nil {
		return nil, err
	}
	return user, nil
}

func FindUserByID(id string) (*UserModel, error) {
	var user UserModel
	result := common.SQLDBUser.Where("user_id = ?", id).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func FindUserByEmail(email string) (*UserModel, error) {
	var user UserModel
	result := common.SQLDBUser.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
