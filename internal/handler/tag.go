package handler

import (
	"article-tags/internal/constants"
	"article-tags/internal/model"
	"article-tags/pkg/response"
	svctypes "article-tags/types"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
)

// AddTag add tags to followed list for particular publication
func (app *Application) AddTag(c *gin.Context) {
	var req svctypes.POSTUserTags

	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.BadRequest(c, "bad request", err.Error())
		return
	}

	err = validateAddTagRequest(c)
	if err != nil {
		log.Println("validation error:", err.Error())
		return
	}

	// save all tags
	for _, tag := range req.Tags {
		input := &model.UserTagRequest{
			Username:    req.Username,
			Tag:         tag,
			Publication: c.Param(constants.Publication),
		}

		// save article
		err = app.ArticleStore.Save(c, input)
		if err != nil {
			log.Println("failed to save tag err: ", err.Error())
			response.InternalServerError(c, "failed to save tag", nil)

			return
		}
	}

	response.Created(c, "tags added successfully", nil)
}

func validateAddTagRequest(c *gin.Context) error {
	// validate publication
	publication := c.Param(constants.Publication)
	var isValid bool
	for _, p := range constants.AllowdedPublications {
		if p == publication {
			isValid = true
			break
		}
	}

	if !isValid {
		response.BadRequest(c, "invalid publication", nil)
		return errors.New("invalid request")
	}

	return nil
}

// GetFollowedTags fetch user followed tags for publication
func (app *Application) GetFollowedTags(c *gin.Context) {
	username := c.Query(constants.Username)
	publication := c.Param(constants.Publication)
	log.Println("fetching tags:", "username:", username, "publication:", publication)

	// fetch all followed tags of user
	data, err := app.ArticleStore.Get(c, publication, username)
	if err != nil {
		log.Println("error fetching tags:", "username:", username, "publication:", publication)
		response.InternalServerError(c, "db operation failed", nil)

		return
	}

	resp := svctypes.GetUserTagsResp{
		Publication: publication,
	}

	// add tags to response
	for _, d := range data {
		resp.Tags = append(resp.Tags, d.SK)
	}

	response.Success(c, "success", resp)
}

// GetPopularTags fetch followed tags by other users for publication.
// and exclude already followed tags by user
func (app *Application) GetPopularTags(c *gin.Context) {
	username := c.Query("username")
	publication := c.Param(constants.Publication)
	log.Println("fetching tags:", "username:", username, "publication:", publication)

	// fetch popular tags
	data, err := app.ArticleStore.GetPopularTags(c, username, publication)
	if err != nil {
		log.Println("error fetching tags:", "username:", username, "publication:", publication)
		response.InternalServerError(c, "db operation failed", nil)

		return
	}

	resp := svctypes.GetUserTagsResp{
		Publication: publication,
	}

	// add tags
	for _, d := range data {
		resp.Tags = append(resp.Tags, d.SK)
	}

	response.Success(c, "success", resp)
}

// DeleteTag delete tag from followed list for particular publication
// and decrement total_count of tag for that tag
func (app *Application) DeleteTag(c *gin.Context) {
	var req svctypes.DeleteTagRequest

	err := c.BindJSON(&req)
	if err != nil {
		response.BadRequest(c, "bad request", err.Error())
		return
	}

	input := &model.UserTagRequest{
		Username:    req.Username,
		Publication: c.Param(constants.Publication),
		Tag:         req.Tag,
	}

	// delete tag
	err = app.ArticleStore.Delete(c, input)
	if err != nil {
		log.Println("failed to save delete err: ", err.Error())
		response.InternalServerError(c, "failed to delete tag", nil)

		return
	}

	response.Success(c, "tag deleted successfully", nil)
}
