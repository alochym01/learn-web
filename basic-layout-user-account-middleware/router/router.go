package router

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/alochym01/learn-web/basic-layout-user-account-middleware/handler"
	"github.com/alochym01/learn-web/basic-layout-user-account-middleware/logger"
	"github.com/alochym01/learn-web/basic-layout-user-account-middleware/service"
	"github.com/alochym01/learn-web/basic-layout-user-account-middleware/storage"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"go.uber.org/zap"
)

func Router(db *sql.DB) *gin.Engine {
	router := gin.New()
	log := logger.NewLogger()

	// Add a ginzap middleware, which:
	//   - Logs all requests, like a combined access and error log.
	//   - Logs to stdout.
	//   - RFC3339 with UTC time format.
	router.Use(ginzap.Ginzap(log, time.RFC3339, true))

	// Logs all panic to error log
	//   - stack means whether output the stack info.
	router.Use(ginzap.RecoveryWithZap(log, true))

	// router := gin.Default()
	// router.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	// router.Use(ginzap.RecoveryWithZap(logger, true))

	// mem := handler.NewAlbumMemory()
	// // Album routes
	// router.GET("/albums", mem.GetAlbums)
	// router.GET("/albums/:id", mem.GetAlbumByID)       // GET HTTP method with /albums/id will be routed to getAlbumByID function
	// router.POST("/albums", mem.PostAlbums)            // POST HTTP method with /albums will be routed to getAlbums function
	// router.PUT("/albums/:id", mem.UpdateAlbumByID)    // PUT HTTP method with /albums/:id will be routed to updateAlbumsByID function
	// router.DELETE("/albums/:id", mem.DeleteAlbumByID) // DELETE HTTP method with /albums/:id will be routed to deleteAlbumsByID function

	// sql := handler.NewAlbumSQL(db)
	// router.GET("/sqlalbums", sql.SQLGetAlbums)
	// router.GET("/sqlalbums/:id", sql.SQLGetAlbumByID)       // GET HTTP method with /sqlalbums/id will be routed to SQLGetAlbumByID function
	// router.POST("/sqlalbums", sql.SQLPostAlbums)            // POST HTTP method with /sqlalbums will be routed to SQLGetAlbums function
	// router.PUT("/sqlalbums/:id", sql.SQLUpdateAlbumByID)    // PUT HTTP method with /sqlalbums/:id will be routed to SQLUpdateAlbumsByID function
	// router.DELETE("/sqlalbums/:id", sql.SQLDeleteAlbumByID) // DELETE HTTP method with /sqlalbums/:id will be routed to SQLDeleteAlbumsByID functio

	albumRepo := storage.NewAlbumSQL(db, log)
	// albumRepo := storage.NewAlbumMemory()
	albumSVCRepo := service.NewAlbumService(&albumRepo)
	albumHandler := handler.NewAlbumHandler(albumSVCRepo)
	// Album routes
	router.GET("/albums", albumHandler.GetAlbums)
	router.GET("/albums/:id", albumHandler.GetAlbumByID)       // GET HTTP method with /albums/id will be routed to getAlbumByID function
	router.POST("/albums", albumHandler.PostAlbums)            // POST HTTP method with /albums will be routed to getAlbums function
	router.PUT("/albums/:id", albumHandler.UpdateAlbumByID)    // PUT HTTP method with /albums/:id will be routed to updateAlbumsByID function
	router.DELETE("/albums/:id", albumHandler.DeleteAlbumByID) // DELETE HTTP method with /albums/:id will be routed to deleteAlbumsByID function

	userRepo := storage.NewUserSQL(db, log)
	// userRepo := storage.NewAlbumMemory()
	userSVCRepo := service.NewUserService(&userRepo)
	userHandler := handler.NewUserHandler(userSVCRepo)
	// User routes
	router.GET("/users", userHandler.GetUsers)
	router.GET("/users/:id", userHandler.GetUserByID)       // GET HTTP method with /Users/id will be routed to getUserByID function
	router.POST("/users", userHandler.PostUser)             // POST HTTP method with /Users will be routed to getUsers function
	router.PUT("/users/:id", userHandler.UpdateUserByID)    // PUT HTTP method with /Users/:id will be routed to updateUsersByID function
	router.DELETE("/users/:id", userHandler.DeleteUserByID) // DELETE HTTP method with /albums/:id will be routed to deleteAlbumsByID function

	loginRepo := storage.NewLoginSQL(db, log)
	loginSVCRepo := service.NewLoginService(&loginRepo)
	loginHandler := handler.NewLoginHandler(loginSVCRepo)
	// Login routes
	router.POST("/login", loginHandler.Login) // POST HTTP method with /login will be routed to Login function

	// Using middleware
	authRoutes := router.Group("/").Use(authMiddleware(log))
	authRoutes.GET("/protected", userHandler.GetUserByID) // GET HTTP method with /Users/id will be routed to getUserByID function
	return router
}

func authMiddleware(l *zap.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Header format
		// GET /users/2 HTTP/1.1
		// Host: 127.0.0.1:8080
		// User-Agent: curl/7.47.0
		// Accept: */*
		// Content-Type: application/json
		// Authorization: Bearer eyJhbGciO.....
		authorizationHeader := ctx.GetHeader("authorization")

		if len(authorizationHeader) == 0 {
			l.Warn("authorization header is not provided")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"errs": "authorization header is not provided",
			})
			return
		}

		// string split "Bearer eyJ...." => ["Bearer", "ey....J"]
		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			l.Warn("invalid authorization header format")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"errs": "authorization header is not provided",
			})
			return
		}

		// Authorization: Bearer eyJhbGciO.....
		authorizationType := strings.ToLower(fields[0])
		if authorizationType != "bearer" {
			l.Warn(fmt.Sprintf("unsupported authorization type %s", authorizationType))
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"errs": authorizationType,
			})
			return
		}

		// Set tokenString
		tokenString := fields[1]

		// Parse, validate, and return a token.
		// https://pkg.go.dev/github.com/golang-jwt/jwt#example-Parse-Hmac
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				l.Warn(fmt.Sprintf("Unexpected signing method: %v", token.Header["alg"]))
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return []byte(service.HMAC_SAMPLE_SECRET), nil
		})

		if err != nil {
			//  handle token expire
			// fmt.Println(err.Error()) // "Token is expired"
			if err.Error() == "Token is expired" {
				l.Warn("Token is expired")
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"errs": "Token is expired",
				})
				return
			}

			//  handle token invalid
			l.Warn(err.Error())
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"errs": token.Header["alg"],
			})

			return
		}

		if token.Valid {
			ctx.Next()
		}

		// if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// 	// Before go to Application
		// 	fmt.Println(claims["exp"])
		// 	fmt.Println("Go Through Middleware")

		// 	ctx.Next()

		// 	// After application process
		// 	fmt.Println("Response to User")
		// }
	}
}
