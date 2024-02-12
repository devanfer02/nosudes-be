package helper

import (
	"time"
)

func CurrentTime() string {
    return time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05")
}

func ErrNotNil() {

}