package crawler

import "github.com/gin-gonic/gin"

type Http struct {
	service *Service
}

func NewHttp(service *Service) *Http {
	return &Http{service: service}
}

// ParseDataByUrls
// @Summary 	Walk all urls and get data for every url
// @Description Add containers to account
// @accept 		json
// @Param 		input body UrlsToParseSchema true "info"
// @Tags        User
// @Success 	200 	{object} []Response
// @Failure 	500 	{object} BaseResponse
// @Router /walk [post]
func (h *Http) ParseDataByUrls(c *gin.Context) {
	var s []string
	if err := c.ShouldBindJSON(&s); err != nil {
		c.JSON(422, gin.H{"error": "invalid request", "success": false})
		return
	}
	responses, err := h.service.Do(c.Request.Context(), deleteDuplicates(s))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error(), "success": false})
		return
	}
	c.JSON(200, responses)
}
