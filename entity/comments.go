package entity

import (
	"time"
)

type Author struct {
	Name string `json:"name,omitempty" bson:"name,omitempty"`
}

type Comment struct {
	ID                string    `bson:"_id" json:"id"`
	CreatedAt         time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt         time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	Author            string    `bson:"author" json:"author"`
	Text              string    `bson:"text" json:"text"`
	PostID            string    `bson:"postId" json:"postId"`
	ParentID          string    `bson:"parentId,omitempty" json:"parentId,omitempty"`
	ChildComments     []string  `bson:"childComments,omitempty" json:"childComments,omitempty"`
	AuthorDetails     *Author    `bson:"authorDetails,omitempty" json:"authorDetails,omitempty"`
	ChildCommentsList []Comment `bson:"childCommentsList,omitempty" json:"childCommentsList,omitempty"`
}
