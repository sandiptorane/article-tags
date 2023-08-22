package handler

import (
	"article-tags/internal/model"
	"article-tags/pkg/response"
	svctypes "article-tags/types"
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

	// check table if exists or not. if not present create new
	err = app.ArticleStore.DescribeTable(c)
	if err != nil {
		err = app.ArticleStore.CreateTable(c)
		if err != nil {
			log.Println("create table failed")
			response.InternalServerError(c, "db operation failed", nil)

			return
		}
	}

	// save all tags
	for _, tag := range req.Tags {
		input := &model.UserTag{
			PK:          req.Username + "#" + req.Publication,
			SK:          tag,
			Publication: req.Publication,
		}

		// check tag is already added or not
		existingTag, err := app.ArticleStore.GetByPublicationTag(c, &model.UserTagRequest{
			Username:    req.Username,
			Publication: req.Publication,
			Tag:         tag,
		})
		if err != nil {
			response.InternalServerError(c, "failed to save tag", nil)
			return
		}

		// skip if already added
		if existingTag != nil {
			continue
		}

		err = app.ArticleStore.Save(c, input)
		if err != nil {
			log.Println("failed to save tag err: ", err.Error())
			response.InternalServerError(c, "failed to save tag", nil)

			return
		}
	}

	response.Created(c, "tags added successfully", nil)
}

// GetFollowedTags fetch user followed tags for publication
func (app *Application) GetFollowedTags(c *gin.Context) {
	username := c.Query("username")
	publication := c.Param("publication")
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
	publication := c.Param("publication")
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
		Publication: c.Param("publication"),
		Tag:         req.Tag,
	}

	// check tag is already deleted or not
	existingTag, err := app.ArticleStore.GetByPublicationTag(c, input)
	if err != nil {
		response.InternalServerError(c, "failed to delete tag", nil)
		return
	}

	// return resp if already deleted
	if existingTag == nil {
		response.Success(c, "success", nil)
		return
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
