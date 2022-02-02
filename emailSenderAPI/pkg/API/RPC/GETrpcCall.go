package RPC

import (
	"emailSenderAPI/pkg/DBlogging"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func ClosureGet(dtb *sqlx.DB) func(*gin.Context) {
	db := dtb
	return func(c *gin.Context) {
		if c.Param("fromMail") == "" {
			c.JSON(400, "No id by user")
			return
		}
		res, err := DBlogging.LogRequestFromMail(db, c.Param("fromMail"))
		if err != nil {
			c.JSON(500, "Something went wrong on our server!")
		} else if len(res) == 0 {
			c.JSON(204, "No messages was found from "+c.Param("fromMail"))
		} else {
			c.JSON(200, res)
		}
	}
}
