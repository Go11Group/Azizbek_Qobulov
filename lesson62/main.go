package main

import (
    "github.com/gin-gonic/gin"
    "github.com/go-redis/redis/v8"
    "github.com/casbin/casbin/v2"
    "context"
    "net/http"
    "log"
)

var (
    ctx = context.Background()
    rdb *redis.Client
    e   *casbin.Enforcer
)

func main() {
    r := gin.Default()

    rdb = redis.NewClient(&redis.Options{
        Addr: "localhost:6379",
    })
    
    var err error
    e, err = casbin.NewEnforcer("model.conf", "policy.csv")
    if err != nil {
        log.Fatalf("Casbin initialization failed: %v", err)
    }

    r.POST("/student", authMiddleware("POST"), createStudent)
    r.GET("/student/:id", authMiddleware("GET"), getStudent)
    r.PUT("/student/:id", authMiddleware("PUT"), updateStudent)
    r.DELETE("/student/:id", authMiddleware("DELETE"), deleteStudent)

    r.Run(":8080")
}




func authMiddleware(action string) gin.HandlerFunc {
    return func(c *gin.Context) {
        role := c.GetHeader("Role")
        obj := c.Request.URL.Path

        ok, err := e.Enforce(role, obj, action)
        if err != nil {
            c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        if !ok {
            c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "access denied"})
            return
        }
        c.Next()
    }
}

func createStudent(c *gin.Context) {
    var student map[string]interface{}
    if err := c.ShouldBindJSON(&student); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    id := student["id"].(string)
    err := rdb.HSet(ctx, id, student).Err()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"status": "student created"})
}

func getStudent(c *gin.Context) {
    id := c.Param("id")
    student, err := rdb.HGetAll(ctx, id).Result()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    if len(student) == 0 {
        c.JSON(http.StatusNotFound, gin.H{"error": "student not found"})
        return
    }
    c.JSON(http.StatusOK, student)
}

func updateStudent(c *gin.Context) {
    id := c.Param("id")
    var student map[string]interface{}
    if err := c.ShouldBindJSON(&student); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    err := rdb.HSet(ctx, id, student).Err()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"status": "student updated"})
}

func deleteStudent(c *gin.Context) {
    id := c.Param("id")
    err := rdb.Del(ctx, id).Err()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"status": "student deleted"})
}
