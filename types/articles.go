package types

// POSTUserTags holds article tags details
type POSTUserTags struct {
	Username    string   `json:"username" bind:"required"`
	Publication string   `json:"publication" bind:"required"`
	Tags        []string `json:"tags" bind:"required"`
}

type GetUserTagsResp struct {
	Publication string   `json:"publication"`
	Tags        []string `json:"tags"`
}

type DeleteTagRequest struct {
	Username string `json:"username" bind:"required"`
	Tag      string `json:"tag" bind:"required"`
}
