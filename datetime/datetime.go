package datetime

import (
	"fmt"
	"time"
)

func TimeNOW_To_Format_YYYYMMDD() string {

	YYYYMMDD := "2006-01-02"
	TO_DAY := fmt.Sprintf("%v", time.Now().UTC().Format(YYYYMMDD))

	return TO_DAY
}
