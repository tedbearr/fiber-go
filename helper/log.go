package helper

import (
	"fmt"
	"os"
	"time"

	"github.com/gookit/slog"
	"github.com/gookit/slog/handler"
)

func Log() {
	date := time.Now().Local().Format("2006-01-02")

	yesterday := time.Now().Local().AddDate(0, 0, -90).Format("2006-01-02")

	delErrLog := os.Remove("log/" + yesterday + ".log")
	delAllLog := os.Remove("log/" + yesterday + "-error.log")

	if delAllLog != nil && delErrLog != nil {
		fmt.Println(delAllLog.Error(), delErrLog.Error())
	}

	errorLog := handler.MustFileHandler("log/"+date+"-error.log", handler.WithLogLevels(slog.DangerLevels))

	allLog := handler.MustFileHandler("log/"+date+".log", handler.WithLogLevels(slog.AllLevels))

	slog.PushHandler(errorLog)
	slog.PushHandler(allLog)
}
