package user

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/phuongaz/forbo/common"
	"github.com/phuongaz/forbo/helper"
	"github.com/phuongaz/forbo/models"
)

func getUser(c *gin.Context) {
	userID := c.Param("id")
	user, err := models.FindUserByID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func registerUser(c *gin.Context) {
	var user models.UserLogin
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := user.Register(); err != nil {

		fmt.Println(err)

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

func loginUser(c *gin.Context) {
	var user models.UserLogin
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	foundUser, err := models.FindUserByEmail(user.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := foundUser.CheckPassword(user.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password or email"})
		return
	}

	token, err := helper.GenerateJWT(foundUser.UserID, foundUser.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token, "user": foundUser, "message": "Login successfully"})
}

func followUser(c *gin.Context) {
	var followData models.FollowData

	if err := c.ShouldBindJSON(&followData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if followData.ID == followData.FollowerID {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You can't follow yourself"})
		return
	}

	user, err := models.FindUserByID(followData.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if user.IsFollowing(followData.FollowerID) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You are already following this user"})
		return
	}

	user.AddFollower(followData.FollowerID)
	if err := user.Update(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Followed successfully"})
}

func unfollowUser(c *gin.Context) {
	var followData models.FollowData

	if err := c.ShouldBindJSON(&followData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := followData.ID
	followerID := followData.FollowerID

	user, err := models.FindUserByID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if !user.IsFollowing(followerID) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You are not following this user"})
		return
	}

	user.RemoveFollower(followerID)
	if err := user.Update(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Unfollowed successfully"})
}

func getFollowers(c *gin.Context) {
	userID := c.Param("id")
	user, err := models.FindUserByID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user.Followers)
}

func getAvatar(c *gin.Context) {
	userID := c.Param("id")
	user, err := models.FindUserByID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if user.Avatar == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "Avatar not found"})
		return
	}

	client, err := common.ConnectMinIO()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	objectName := user.Avatar
	bucketName := "avatars"
	file, err := common.DownloadFile(client, bucketName, objectName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Header("Content-Disposition", "attachment; filename="+objectName)
	c.Header("Content-Type", "application/octet-stream")
	c.File(file)
}

func uploadAvatar(c *gin.Context) {
	userID := c.Param("id")
	user, err := models.FindUserByID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	file, err := c.FormFile("avatar")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bucketName := "avatars"
	objectName := userID + "/" + file.Filename
	client, err := common.ConnectMinIO()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = common.UploadFile(client, bucketName, objectName, file)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user.Avatar = objectName
	if err := user.Update(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Avatar uploaded successfully"})
}
