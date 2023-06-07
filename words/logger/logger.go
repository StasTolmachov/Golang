package logger

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

type CustomFormatter struct{}

func (f *CustomFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := fmt.Sprintf("%v", entry.Time.Format("2006-01-02T15:04:05.000Z07:00"))
	level := strings.ToUpper(entry.Level.String())

	// Выводим уровень INFO синим цветом
	if level == "INFO" {
		level = "\033[34m" + level + "\033[0m"
	}

	message := entry.Message

	_, file, line, _ := runtime.Caller(9)

	// Изменяем порядок вывода элементов в логе
	return []byte(fmt.Sprintf("%s\t%s\t%s:%d\t%s\n", level, timestamp, file, line, message)), nil
}

func LogSetup() {
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetReportCaller(true)
	logrus.SetFormatter(new(CustomFormatter))
}
