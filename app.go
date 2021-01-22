package demo

import (
	"demo/beer"
	_ "demo/docs"

	"github.com/gin-gonic/gin"
	"github.com/penglongli/gin-metrics/ginmetrics"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//StartServer na
func StartServer() {
	// r := gin.Default()
	r := gin.New()
	r.Use(gin.Recovery())
	r.GET("/")
	// r.GET("/ping/:name", gin.Logger(), ping)

	// get global Monitor object
	m := ginmetrics.GetMonitor()

	// +optional set metric path, default /debug/metrics
	m.SetMetricPath("/metrics")
	// +optional set slow time, default 5s
	m.SetSlowTime(10)
	// +optional set request duration, default {0.1, 0.3, 1.2, 5, 10}
	// used to p95, p99
	m.SetDuration([]float64{0.1, 0.3, 1.2, 5, 10})

	// set middleware for gin
	m.Use(r)
	beer.NewRoutes(r)
	// ===== Route of documents
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":1234")

	// r.GET("/beer", ping)

}

// func ping(c *gin.Context) {
// 	// TODO :: call service
// 	name := c.Param("name")
// 	c.JSON(200, gin.H{
// 		"message": "pong" + name,
// 	})
// }
