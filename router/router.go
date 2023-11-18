// router/router.go

package router

import (
    "github.com/gin-gonic/gin"
    "github.com/adindazenn/kelompok6/controller" 
)

func SetupRouter() *gin.Engine {
    r := gin.Default()

    r.POST("/users/register", controller.RegisterUser)

    r.POST("/users/login", controller.LoginUser)

    r.PUT("/users/update-account", controller.UpdateAccount)

    r.DELETE("/users/delete-account", controller.DeleteAccount)

    //Category
    r.POST("/categories", controller.CreateCategory)

    r.GET("/categories", controller.GetCategories)

    r.PATCH("/categories/:categoryId", controller.UpdateCategory)

    r.DELETE("/categories/:categoryId", controller.DeleteCategory)

    //Task
    r.POST("/tasks", controller.CreateTask)

    r.GET("/tasks", controller.GetTasks)

    r.PUT("/tasks/:taskId", controller.UpdateTask)

    r.PATCH("tasks/update-status/:taskId", controller.UpdateTaskStatus)

    r.PATCH("tasks/update-category/:taskId", controller.UpdateTaskCategory)

    r.DELETE("tasks/:taskId", controller.DeleteTask)

    return r
}
