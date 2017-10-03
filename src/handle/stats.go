package handle

import (
	stats_api "github.com/fukata/golang-stats-api-handler"
	"github.com/labstack/echo"
)

func NewStatsHandle(group *echo.Group) {
	group.GET("/stats", StatsHandler)
}
func StatsHandler(c echo.Context) error {
	return c.JSON(200, stats_api.GetStats())
}
