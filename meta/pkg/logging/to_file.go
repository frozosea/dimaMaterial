package logging

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
)

func createDirsAndFiles(dirName string) error {
	if _, err := os.ReadDir(`logs`); err != nil {
		if mkDirErr := os.Mkdir(`logs`, 0700); mkDirErr != nil {
			return err
		}
	}
	if _, err := os.ReadDir(`logs`); err != nil {
		if createLogDirErr := os.Mkdir(`logs`, 0700); createLogDirErr != nil {
			return createLogDirErr
		}
	}
	if err := os.Mkdir(fmt.Sprintf(`logs/%s`, dirName), 0700); err != nil {
		return err
	}
	logs := []string{"info.log", "exception.log", "warning.log", "fatal.log"}
	for _, v := range logs {
		if info := os.WriteFile(fmt.Sprintf(`logs/%s/%s`, dirName, v), nil, 0700); info != nil {
			return info
		}
	}
	return nil
}

func initOneLogger(filePath string) *log.Logger {
	logger := log.New()
	d, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return logger
	}
	logger.Out = d
	logger.Formatter = &log.JSONFormatter{}
	return logger
}

type Logger struct {
	SaveLogsDir     string
	infoLogger      log.Logger
	exceptionLogger log.Logger
	warningLogger   log.Logger
	fatalLogger     log.Logger
}

func (l *Logger) InfoLog(logString string) {
	l.infoLogger.Info(logString)
}
func (l *Logger) ExceptionLog(logString string) {
	l.exceptionLogger.Error(logString)
}
func (l *Logger) WarningLog(logString string) {
	l.warningLogger.Warning(logString)
}
func (l *Logger) FatalLog(logString string) {
	l.fatalLogger.Fatal(logString)
}

func NewToFile(saveLogsDir string) *Logger {
	if err := createDirsAndFiles(saveLogsDir); err != nil {
		return &Logger{}
	}
	return &Logger{
		SaveLogsDir:     saveLogsDir,
		infoLogger:      *initOneLogger(fmt.Sprintf(`logs/%s/info.log`, saveLogsDir)),
		exceptionLogger: *initOneLogger(fmt.Sprintf(`logs/%s/exception.log`, saveLogsDir)),
		warningLogger:   *initOneLogger(fmt.Sprintf(`logs/%s/warning.log`, saveLogsDir)),
		fatalLogger:     *initOneLogger(fmt.Sprintf(`logs/%s/fatal.log`, saveLogsDir)),
	}
}
