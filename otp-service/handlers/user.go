package handlers

import (
	"fmt"
	"net/http"
	"otp-service/models"
	"sort"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// GetSingleUserByPhone retrieves a user by phone number
// @Summary Get a user by phone number
// @Description Retrieve a single user by their phone number
// @Tags User
// @Accept json
// @Produce json
// @Param phone query string true "Phone number of the user"
// @Success 200 {object} models.User "User found"
// @Failure 404 {object} map[string]interface{} "User not found"
// @Router /user [get]
func GetSingleUserByPhone(c *gin.Context) {
	phone := c.Query("phone")

	user, exists := models.Users[phone]

	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}


// GetUsersWithPagination retrieves a paginated list of users
// @Summary Get a paginated list of users
// @Description Retrieve all users with pagination based on query parameters
// @Tags User
// @Accept json
// @Produce json
// @Param page query int false "Page number (default: 1)"
// @Param limit query int false "Number of users per page (default: 10)"
// @Success 200
// @Failure 400 {object} map[string]interface{} "Invalid query parameters"
// @Router /users [get]
func GetUsersWithPagination(c *gin.Context) {
	// Default values for pagination
	fmt.Println("Hi Sina!!!")

	defaultPage := 1
	defaultLimit := 10

	// Parse 'page' query parameter
	pageStr := c.Query("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 {
		page = defaultPage // Default to page 1 if invalid
	}

	// Parse 'limit' query parameter
	limitStr := c.Query("limit")
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = defaultLimit // Default to limit 10 if invalid
	}

	// Calculate offset
	offset := (page - 1) * limit

	// Convert the map of users to a slice for sorting
	var allUsers []models.User
	for _, user := range models.Users {
		allUsers = append(allUsers, user)
	}

	// Sort the users by phone number (or name, if desired)
	sort.Slice(allUsers, func(i, j int) bool {
		return allUsers[i].Phone < allUsers[j].Phone // Sort by phone number
	})

	// Apply pagination logic
	totalUsers := len(allUsers)
	end := offset + limit
	if end > totalUsers {
		end = totalUsers
	}

	if offset >= totalUsers {
		// If offset is greater than total users, return an empty list
		c.JSON(http.StatusOK, gin.H{
			"users":      []models.User{},
			"total":      totalUsers,
			"page":       page,
			"limit":      limit,
			"totalPages": (totalUsers + limit - 1) / limit, // Calculate total pages
		})
		return
	}

	// Paginate the sorted user slice
	paginatedUsers := allUsers[offset:end]

	// Respond with paginated users and metadata
	c.JSON(http.StatusOK, gin.H{
		"users":      paginatedUsers,
		"total":      totalUsers,
		"page":       page,
		"limit":      limit,
		"totalPages": (totalUsers + limit - 1) / limit, // Calculate total pages
	})
}


// @Summary Create a new user
// @Description Create a new user with phone number and optional name
// @Tags User
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{} "User created successfully"
// @Failure 400 {object} map[string]interface{} "Invalid request payload"
// @Failure 409 {object} map[string]interface{} "User already exists"
// @Router /users [post]
// CreateUser handles user creation
func CreateUser(c *gin.Context) {
	// Define the request structure
	var request struct {
		Phone string `json:"phone" binding:"required"` // Phone is required
		Name  string `json:"name"`                    // Name is optional
	}

	// Parse JSON payload
	err := c.ShouldBindJSON(&request) // Bind the JSON payload to the `request` variable
	if err != nil {                  // Check if an error occurred
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Check if the phone number is already registered
	_, exists := models.Users[request.Phone];
	if exists {	
		c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
		return
	}

	// Create the new user
	newUser := models.User{
		Phone:           request.Phone,
		Name:            request.Name,
		RegistrationDate: time.Now(), // Set registration date to current time
	}

	// Store the new user in the in-memory map
	models.Users[request.Phone] = newUser

	// Respond with success and the created user data
	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"user":    newUser,
	})
}
