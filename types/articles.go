package types

// POSTUserTags holds article tags details
type POSTUserTags struct {
	Username string   `json:"username" binding:"required"`
	Tags     []string `json:"tags" binding:"required"`
}

type GetUserTagsResp struct {
	Publication string   `json:"publication"`
	Tags        []string `json:"tags"`
}

type DeleteTagRequest struct {
	Username string `json:"username" binding:"required"`
	Tag      string `json:"tag" binding:"required"`
}
