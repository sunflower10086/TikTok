package dao

import (
	"context"
	"github.com/sunflower10086/TikTok/interaction/internal/dao/db"
	"github.com/sunflower10086/TikTok/interaction/internal/dao/models"
	"github.com/sunflower10086/TikTok/interaction/internal/dao/models/modelToimpl"
	interaction "github.com/sunflower10086/TikTok/interaction/pb"
)

type Favorite struct {
	UserID  int64
	VideoID int64
}

// 点赞
func AddFavorite(ctx context.Context, userID int64, videoID int64) error {
	conn := db.GetDB().WithContext(ctx)

	// 由于userID是通过token解析出来的，因此userID如果有值则一定存在该用户
	// 判断videoID的合法性
	var video models.Video
	err2 := conn.First(&video, "id = ?", videoID).Error
	if err2 != nil {
		return err2
	}

	// 判断是否已点赞
	check, err := CheckIsFavorite(ctx, videoID, userID)
	if err != nil {
		return err
	}

	// 如果未点赞则点赞
	if !check {
		err = conn.Table("user_favorite").Create(&Favorite{
			UserID:  userID,
			VideoID: videoID,
		}).Error

		if err != nil {
			return err
		}
	}

	return nil
}

// 取消点赞
func DelFavorite(ctx context.Context, userID int64, videoID int64) error {
	conn := db.GetDB().WithContext(ctx)

	// 判断是否已点赞
	check, err := CheckIsFavorite(ctx, videoID, userID)
	if err != nil {
		return err
	}

	// 如果已点赞则取消点赞
	if check {
		err = conn.Table("user_favorite").
			Where("user_id = ? and video_id = ?", userID, videoID).
			Delete(&Favorite{}).Error
		if err != nil {
			return err
		}
	}

	return nil
}

// 获取点赞列表
func GetFavoriteList(ctx context.Context, userID int64) ([]*interaction.Video, error) {
	var videoID []int64                // 用户点赞视频的ID
	videos := make([]*models.Video, 0) // 数据层

	conn := db.GetDB().WithContext(ctx)

	// 查询用户点赞视频的ID
	err := conn.Table("user_favorite").
		Where("user_id = ?", userID).
		Pluck("video_id", &videoID).Error

	if err != nil {
		return nil, err
	}

	// 根据视频ID查询对应视频
	err = conn.Preload("User.OtherInfo").
		Preload("User").
		Where("id in ?", videoID).
		Find(&videos).Error

	videos2 := make([]*interaction.Video, len(videos)) // 业务层

	// 获取实时的Video和User信息
	for i, v := range videos {
		err = RealTimeVideo(ctx, v)
		if err != nil {
			return nil, err
		}
		err = RealTimeUser(ctx, &v.User)
		if err != nil {
			return nil, err
		}

		// 数据层映射到业务层
		videos2[i] = modelToimpl.MapFavorite(v)
	}

	return videos2, nil
}
