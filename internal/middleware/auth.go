package middleware

import (
    "api-gateway/pkg/constant"
    "api-gateway/pkg/response"
    "fmt"
    "net/http"
    "strings"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v4"
)

var jwtSecret = []byte("secretOfKey")

var publicPaths = []string{
    "/auth/signin",
    "/auth/signup",
    "/auth/google",
    "/auth/google/callback",
    "/auth/verify",
    "/health",
}

func isPublicPath(path string) bool {
    for _, p := range publicPaths {
        if p == path {
            return true
        }
    }
    return false
}

func Auth() gin.HandlerFunc {
    return func(c *gin.Context) {
        if isPublicPath(c.Request.URL.Path) {
            c.Next()
            return
        }

        authHeader := c.GetHeader(constant.HeaderAuthorization)
        if authHeader == "" {
            response.Error(c, http.StatusUnauthorized, constant.ErrCodeUnauthorized, "Missing authorization header")
            c.Abort()
            return
        }

        tokenString := strings.TrimPrefix(authHeader, "Bearer ")
        if tokenString == authHeader {
            response.Error(c, http.StatusUnauthorized, constant.ErrCodeUnauthorized, "Invalid Authorization header format")
            c.Abort()
            return
        }

        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
            }
            return jwtSecret, nil
        })

        if err != nil || !token.Valid {
            response.Error(c, http.StatusUnauthorized, constant.ErrCodeUnauthorized, "Invalid or expired token")
            c.Abort()
            return
        }

        claims, ok := token.Claims.(jwt.MapClaims)
        if !ok {
            response.Error(c, http.StatusUnauthorized, constant.ErrCodeUnauthorized, "Invalid token claims")
            c.Abort()
            return
        }
        fmt.Printf("Claims: %+v\n", claims)

        if exp, ok := claims["exp"].(float64); ok {
            if time.Unix(int64(exp), 0).Before(time.Now()) {
                response.Error(c, http.StatusUnauthorized, constant.ErrCodeUnauthorized, "Token has expired")
                c.Abort()
                return
            }
        }

        userIDFloat, ok := claims["id"].(float64)
        if !ok {
            response.Error(c, http.StatusUnauthorized, constant.ErrCodeUnauthorized, "Invalid user ID in token")
            c.Abort()
            return
        }

        userID := fmt.Sprintf("%.0f", userIDFloat)

        email, _ := claims["email"].(string)
        username, _ := claims["username"].(string)
        role, _ := claims["role"].(string)

        if userID == "" {
            response.Error(c, http.StatusUnauthorized, constant.ErrCodeUnauthorized, "Invalid user ID in token")
            c.Abort()
            return
        }

        c.Set("user_id", userID)
        c.Set("email", email)
        c.Set("username", username)
        c.Set("role", role)

        c.Request.Header.Set("X-User-ID", userID)
        fmt.Println("X-User-ID Header Set:", userID)

        c.Next()
    }
}
