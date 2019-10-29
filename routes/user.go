package routes

import (
	"net/http"

	"github.com/PacoDw/introduction_to_GORM/models"
	s "github.com/PacoDw/introduction_to_GORM/models/structs"
	"github.com/PacoDw/introduction_to_GORM/services"
	"github.com/gin-gonic/gin"
)

/*User routes */
func User(route *gin.RouterGroup) {

	// --------------------------------------------------------------------
	// Create a user
	route.POST("/create", func(ctx *gin.Context) {
		user := &models.User{}

		if err := ctx.BindJSON(user); err != nil {
			ctx.Status(http.StatusBadRequest)
			return
		}

		response := services.Create(user)
		code := http.StatusOK
		if !response.Success {
			code = http.StatusBadRequest
		}

		ctx.JSON(code, response)
	})

	// --------------------------------------------------------------------
	// Get All Users
	route.GET("/getAll", func(ctx *gin.Context) {
		response := services.FindAll(&models.User{})

		code := http.StatusOK
		if !response.Success {
			code = http.StatusBadRequest
		}

		ctx.JSON(code, response)
	})

	// --------------------------------------------------------------------
	// Get one users by id
	route.GET("/getOne/:id", func(ctx *gin.Context) {

		response := services.FindOneByID(&models.User{ID: ctx.Param("id")})

		code := http.StatusOK
		if !response.Success {
			code = http.StatusBadRequest
		}

		ctx.JSON(code, response)
	})

	// --------------------------------------------------------------------
	// Update a User
	route.PUT("/update/:id", func(ctx *gin.Context) {
		user := &models.User{}

		if err := ctx.BindJSON(user); err != nil {
			ctx.Status(http.StatusBadRequest)
			return
		}

		user.ID = ctx.Param("id")

		response := services.Update(user)
		code := http.StatusOK
		if !response.Success {
			code = http.StatusBadRequest
		}

		ctx.JSON(code, response)
	})

	// --------------------------------------------------------------------
	// Delete a User by id
	route.DELETE("/delete/:id", func(ctx *gin.Context) {
		response := services.DeleteOneByID(&models.User{ID: ctx.Param("id")})
		code := http.StatusOK
		if !response.Success {
			code = http.StatusBadRequest
		}

		ctx.JSON(code, response)
	})

	// --------------------------------------------------------------------
	// Delete a Users by ids
	route.POST("/delete", func(ctx *gin.Context) {
		ids := &s.IDs{}

		if err := ctx.BindJSON(ids); err != nil {
			ctx.Status(http.StatusBadRequest)
			return
		}

		response := services.DeleteByIDS(ids, &models.User{})
		code := http.StatusOK
		if !response.Success {
			code = http.StatusBadRequest
		}

		ctx.JSON(code, response)
	})
}
