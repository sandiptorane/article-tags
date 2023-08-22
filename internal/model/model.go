package model

// UserTag holds article details
type UserTag struct {
	PK          string `dynamodbav:"PK"`
	SK          string `dynamodbav:"SK"`
	Publication string `dynamodbav:"publication"`
}

// UserTagRequest holds request params
type UserTagRequest struct {
	Username    string
	Publication string
	Tag         string
}
