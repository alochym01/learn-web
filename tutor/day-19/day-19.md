# Gin hello world - Day 19

## Using zap logger

1. We create folder structure

   ```bash
   .
   ├── domain
   │   └── album.go
   ├── go.mod
   ├── go.sum
   ├── handler
   │   └── album.go
   ├── logger
   │   └── log.go
   ├── main.go
   ├── service
   │   └── album.go
   ├── storage
   │   └── memory
   │       └── album.go
   └── router
       └── router.go
   ```

2. We are using `go.uber.org/zap`. Create a [`logger/log.go`](logger/log.go) file to implement log handler

   ```go
   import (
      "go.uber.org/zap"
      "go.uber.org/zap/zapcore"
   )

   var log *zap.Logger

   // NewLogger ...
   func NewLogger() *zap.Logger {
      config := zap.NewProductionConfig()

      encoderConfig := zap.NewProductionEncoderConfig()
      encoderConfig.TimeKey = "timestamp"
      encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
      encoderConfig.StacktraceKey = ""
      config.EncoderConfig = encoderConfig

      log, err := config.Build(zap.AddCallerSkip(1))

      if err != nil {
         panic(err)
      }

      return log
   }

   // Info ...
   func Info(message string, fields ...zap.Field) {
      log.Info(message, fields...)
   }

   // Warn ...
   func Warn(message string, fields ...zap.Field) {
      log.Warn(message, fields...)
   }

   // Error ...
   func Error(message string, fields ...zap.Field) {
      log.Error(message, fields...)
   }
   ```

3. Change `router/router.go` using `ZAP logger`

   ```go
   import (
      "github.com/alochym01/web-w-gin/logger"
      ...
   )

   // Router return a gin.Engine
   func Router(db *sql.DB) *gin.Engine {
      // router := gin.Default()
      router := gin.New()

      // Using ZAP Logger
      logger := logger.NewLogger()

      // Add a ginzap middleware, which:
      //   - Logs all requests, like a combined access and error log.
      //   - Logs to stdout.
      //   - RFC3339 with UTC time format
      router.Use(ginzap.Ginzap(logger, time.RFC3339, true))

      // Logs all panic to error log
      //   - stack means whether output the stack info.
      router.Use(ginzap.RecoveryWithZap(logger, true))

      // New Album Storage - SQLite
      storeAlbum := sqlite3.NewAlbum(db, logger)
      ...
   ```

4. Change `storage/sqlite3/album.go`

   ```go
   // Album ...
   type Album struct {
      db     *sql.DB
      logger *zap.Logger
   }

   // NewAlbum ...
   func NewAlbum(DB *sql.DB, l *zap.Logger) Album {
      return Album{
         db:     DB,
         logger: l,
      }
   }
   ```

## Using curl to test

1. curl localhost:8080/albums
2. curl localhost:8080/albums/1
3. curl -XPOST -H "Content-Type: application/json" -d @createAlbum.json http://127.0.0.1:8080/albums
4. curl -XDELETE -H "Content-Type: application/json" http://127.0.0.1:8080/albums/1
5. curl -XPUT -H "Content-Type: application/json" -d @updateAlbum.json http://127.0.0.1:8080/albums/3
