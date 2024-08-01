package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"war_ticket/internal/domain"
	"war_ticket/internal/repository"
	"war_ticket/internal/repository/db_repo"
	"war_ticket/pkg"
)

type contextKey string

const ContextUserKey contextKey = "user"

func AuthMiddleware(next http.HandlerFunc, userRepo repository.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			logger pkg.LogFormat
			user   *domain.User
			err    error
		)

		defer func() {
			logger.Execute()
		}()

		apiKey := r.Header.Get("X-API-KEY")

		if user, err = isValidAPIKey(apiKey, userRepo); err != nil {
			logger.Error = err.Error()
			logger.Message = "Unauthorized request"
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), ContextUserKey, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

func AuthMiddlewareGin(userRepo repository.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			logger pkg.LogFormat
			user   *domain.User
			err    error
		)

		defer func() {
			logger.Execute()
		}()

		apiKey := c.GetHeader("X-API-KEY")

		if user, err = isValidAPIKey(apiKey, userRepo); err != nil {
			logger.Error = err.Error()
			logger.Message = "Unauthorized request"
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		log.Println("user middle :: ", user)
		c.Set(string(ContextUserKey), user)
		c.Next()
	}
}

func isValidAPIKey(apiKey string, userRepository db_repo.UserRepository) (*domain.User, error) {

	user, err := userRepository.FindByApiKey(apiKey)

	if err != nil {
		return nil, err
	}

	return user, nil
}
