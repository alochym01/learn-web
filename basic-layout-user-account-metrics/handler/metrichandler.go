package handler

import (
	"database/sql"
	"expvar"
	"fmt"

	"github.com/gin-gonic/gin"
)

type Metrics struct {
	DB *sql.DB
}

func (m Metrics) Metric(c *gin.Context) {

	w := c.Writer
	c.Header("Content-Type", "application/json; charset=utf-8")
	_, _ = w.Write([]byte("{\n"))
	first := true
	expvar.Do(func(kv expvar.KeyValue) {
		if !first {
			_, _ = w.Write([]byte(",\n"))
		}
		first = false
		fmt.Fprintf(w, "%q: %s", kv.Key, kv.Value)
	})
	_, _ = w.Write([]byte("\n}\n"))
	c.AbortWithStatus(200)
}
