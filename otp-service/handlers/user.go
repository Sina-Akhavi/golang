package handlers

import (
	"net/http"
	"otp-service/models"
	"sort"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetSingleUserByPhone(c *gin.Context) {
	phone := c.Query("phone")

	user, exists := models.Users[phone]

	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}


func GetUsersWithPagination(c *gin.Context) {
	// Default values for pagination
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
