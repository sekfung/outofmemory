package models

import (
	"github.com/jinzhu/gorm"
	"outofmemory/errors"
)

type Comment struct {
	BaseModel
	CommentId  uint32
	TopicId    uint32
	TopicType  uint8
	Content    string
	FromUid    uint32
	FromAvatar string
}

func AddComment(data map[string]interface{}) (interface{}, error){
	var (
		commentData = parseCommentData(data)
		comment Comment
	)

	err := db.Create(&commentData)
	if err != nil {
		return nil, errors.ErrCreateCommentFailed
	}
	return comment, nil
}

func existCommentByID(commentID uint32) (Comment, error) {
	var comment Comment
	err := db.Where("comment_id = ?", commentID).Find(&comment).Error
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return comment, errors.ErrCommentNotExist
		default:
			return comment, errors.ErrGetCommentFailed

		}
	}
	return comment, nil
}

func parseCommentData(data map[string]interface{}) Comment  {
	comment := Comment{
		CommentId:  data["comment_id"].(uint32),
		TopicId:    data["topic_id"].(uint32),
		TopicType:  data["topic_type"].(uint8),
		FromUid:    data["from_uid"].(uint32),
		FromAvatar: data["from_avatar"].(string),
	}
	return comment
}
