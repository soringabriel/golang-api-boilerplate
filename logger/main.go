package logger

import (
	"strconv"

	"api/helpers"

	"github.com/google/uuid"
	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
)

var Instance *logrus.Entry

type LoggerSettings struct {
	LogFile       string
	LogMaxSizeMB  int
	LogMaxBackups int
	LogMaxAge     int
}

func SetupLogger() {
	if Instance == nil {
		logrus.SetFormatter(&logrus.JSONFormatter{})
		// Setup logger
		ip, err := helpers.GetOutboundIP()
		if err != nil {
			logrus.Fatalf("Failed to get outbound IP: %v", err)
		}
		appName := helpers.EnvVariable("APP_NAME")
		Instance = logrus.WithFields(logrus.Fields{
			"hostname": ip,
			"appname":  appName,
			"session":  uuid.New().String(),
		})

		// Logging to file
		logFile := helpers.EnvVariable("LOG_FILE")
		if len(logFile) > 0 {
			logMaxSizeMB, err := strconv.Atoi(helpers.EnvVariable("LOG_MAX_SIZE_MB"))
			if err != nil {
				Instance.Fatal("Failed to convert LOG_MAX_SIZE_MB env variable to int", err)
			}
			logMaxBackups, err := strconv.Atoi(helpers.EnvVariable("LOG_MAX_BACKUPS"))
			if err != nil {
				Instance.Fatal("Failed to convert LOG_MAX_BACKUPS env variable to int", err)
			}
			logMaxAge, err := strconv.Atoi(helpers.EnvVariable("LOG_MAX_AGE"))
			if err != nil {
				Instance.Fatal("Failed to convert LOG_MAX_AGE env variable to int", err)
			}
			Instance.Logger.SetOutput(&lumberjack.Logger{
				Filename:   logFile,
				MaxSize:    logMaxSizeMB,
				MaxBackups: logMaxBackups,
				MaxAge:     logMaxAge,
			})
		}
	}
}
