// package api

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	"gorm.io/gorm"
// )

// type healthCheckHandler struct{
// 	db *gorm.DB
// }

// func(h *healthCheckHandler) healthCheckHandler(c *gin.Context) {
//     // Check database connectivity
//     if err := h.db.ConnPool.(); err != nil {
//         c.JSON(http.StatusInternalServerError, gin.H{
//             "status": "error",
//             "message": "Database connection error",
//         })
//         return
//     }

//     // If database connectivity is successful, return a success message
//     c.JSON(http.StatusOK, gin.H{
//         "status": "ok",
//         "message": "Application and database are healthy",
//     })
// }