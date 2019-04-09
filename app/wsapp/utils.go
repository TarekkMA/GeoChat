package wsapp

import (
	"strconv"
	"time"

	"github.com/google/uuid"
)

func TimeNowNanoStr() string {
	return strconv.FormatInt(time.Now().UnixNano(), 10)
}

func NewMessageID() string {
	return "msg-" + uuid.New().String()
}
