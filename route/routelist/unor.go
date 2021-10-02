package routelist

import (
	// "code-echo/api/peremajaan"

	"code-echo/api/peremajaan"

	"github.com/labstack/echo/v4"
)

func SettingUnor(echo *echo.Group) {
	echo.POST("/add", peremajaan.SimpanUnor)
	echo.POST("/update", peremajaan.UpdateUnor)
	echo.GET("/get", peremajaan.GetUnor)
	echo.DELETE("/delete", peremajaan.DeleteUnor)
}
