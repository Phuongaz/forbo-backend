package models

import (
	"github.com/phuongaz/forbo/common"
	"gorm.io/gorm"
)

type FeedSkeleton struct {
	UserID  string `json:"user_id"`
	Content string `json:"content"`
	Image   string `json:"image"`
}

type Feed struct {
	gorm.Model
	FeedID   string `json:"feed_id" gorm:"primary_key"`
	UserID   string `json:"user_id" grom:"size:255"`
	Content  string `json:"content"`
	Image    string `json:"image"`
	Like     int32  `json:"like" gorm:"default:0"`
	CreateAt string `json:"create_at" gorm:"default:CURRENT_TIMESTAMP"`
}

func (f *Feed) Create() error {
	return common.SQLDBFeed.Create(f).Error
}

func FindFeedByID(id string) (*Feed, error) {
	var feed Feed
	err := common.SQLDBFeed.Where("feed_id = ?", id).First(&feed).Error
	return &feed, err
}

func FindFeedsByUserID(userID string) ([]Feed, error) {
	var feeds []Feed
	err := common.SQLDBFeed.Where("user_id = ?", userID).Find(&feeds).Error
	return feeds, err
}

func (f *Feed) Update() error {
	return common.SQLDBFeed.Save(f).Error
}

func (f *Feed) Delete() error {
	return common.SQLDBFeed.Delete(f).Error
}

func (nf *FeedSkeleton) ToFeed() *Feed {

	feedID := common.GenerateFeedID(nf.UserID)

	return &Feed{
		FeedID:  feedID,
		UserID:  nf.UserID,
		Content: nf.Content,
		Image:   nf.Image,
	}
}
