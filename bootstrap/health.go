package bootstrap

import (
    h "net/http"
	s "database/sql"
	f "fmt"
	r "runtime"
    t "time"

	"github.com/gin-gonic/gin"
)

type healthInfo struct {
    Alloc       string      `json:"alloc"`
    TotalAlloc  string      `json:"total_alloc"`
    Version     string      `json:"go_version"`
    Lookups     uint64      `json:"pointer_lookups"`
    Sys         string      `json:"sys"`
}

func bToMb(b uint64) uint64 {
    return b / 1024 / 1024
}

func getMemStats() *healthInfo {
    var m r.MemStats
    r.ReadMemStats(&m)

    healthInfo := &healthInfo{
        Alloc: f.Sprintf("%v MiB", bToMb(m.Alloc)),
        TotalAlloc: f.Sprintf("%v MiB",bToMb(m.Alloc)),
        Version: r.Version(),
        Lookups: m.Lookups,
        Sys: f.Sprintf("%v MiB",bToMb(m.Sys)),
    }

    return healthInfo
}

func GetHealthCheck(db *s.DB) gin.HandlerFunc {
    return func (ctx *gin.Context) {
        err := db.Ping()
        info := getMemStats()
        curr := t.Now().UTC().Format("2006-01-02 15:04")
        
        if err != nil {
            ctx.JSON(h.StatusServiceUnavailable, gin.H{
                "status": "error could not ping sql db",
                "time": curr,
                "info": info,
            })
            return
        }

        ctx.JSON(h.StatusOK, gin.H{
            "status": "server running ok",
            "time": curr,
            "info": info,
        })
    }
}