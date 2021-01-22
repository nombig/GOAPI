package beer

import "github.com/gin-gonic/gin"

//NewRoutes d
func NewRoutes(r *gin.Engine) {
	b := r.Group("/beer")
	b.GET("/", getAllBeer)
}

//BeerResponse d
//@Produce json
// @Summary Retrieves users from mongodb
// @Description Get Users
// @Produce json
type BeerResponse struct {
	Message string `json:"message"`
}

// getAllBeer From Database
// @Summary Retrieves users from mongodb
// @Description Get Users
// @Produce json
// @Param name query string false "Name"
// @Param age query int false "Age"
// @Success 200 {object} string
// @Router /beer [get]
func getAllBeer(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Get all beer",
	})
}
