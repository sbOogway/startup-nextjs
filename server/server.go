package main

import (
    "database/sql"
    // "fmt"
    "log"
    "net/http"

    "github.com/gin-gonic/gin"
    _ "github.com/go-sql-driver/mysql" // Import the MySQL driver
)

// Database connection string
const dsn = "username:password@tcp(127.0.0.1:3306)/dbname"

func main() {
    r := gin.Default()

    r.POST("/add-user", addUser)

    if err := r.Run(":8080"); err != nil {
        log.Fatal(err)
    }
}

// Handler function for adding a user
func addUser(c *gin.Context) {
    // Open a new database connection
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to the database"})
        return
    }
    defer db.Close() // Ensure the database connection is closed when done

    // Check if the connection is alive
    if err := db.Ping(); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection is not alive"})
        return
    }

    // Get the user data from the request body
    var user struct {
        Name  string `json:"name" binding:"required"`
        Email string `json:"email" binding:"required"`
    }

    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    // Perform the database operation (e.g., insert a new user)
    _, err = db.Exec("INSERT INTO users (name, email) VALUES (?, ?)", user.Name, user.Email)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert user"})
        return
    }

    // Respond with success
    c.JSON(http.StatusOK, gin.H{"message": "User added successfully"})
}
