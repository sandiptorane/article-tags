package types

// POSTUserTags holds article tags details
type POSTUserTags struct {
	Username    string   `json:"username" bind:"required"`
	Publication string   `json:"" bind:"required"`
	Tags        []string `json:"tags" bind:"required"`
}
