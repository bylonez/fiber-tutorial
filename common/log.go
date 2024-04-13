package common

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"github.com/sirupsen/logrus"
	"github.com/valyala/bytebufferpool"
	"sync"
)

var logger *logrus.Logger

type CustomLogger struct {
	*logrus.Logger
	entry *logrus.Entry // need add usage when necessary
}

func (c *CustomLogger) Tracew(msg string, keysAndValues ...interface{}) {
	c.Logger.Trace(c.fmtMsg(msg, keysAndValues))
}

func (c *CustomLogger) Debugw(msg string, keysAndValues ...interface{}) {
	c.Logger.Debug(c.fmtMsg(msg, keysAndValues))
}

func (c *CustomLogger) Infow(msg string, keysAndValues ...interface{}) {
	c.Logger.Info(c.fmtMsg(msg, keysAndValues))
}

func (c *CustomLogger) Warnw(msg string, keysAndValues ...interface{}) {
	c.Logger.Warn(c.fmtMsg(msg, keysAndValues))
}

func (c *CustomLogger) Errorw(msg string, keysAndValues ...interface{}) {
	c.Logger.Error(c.fmtMsg(msg, keysAndValues))
}

func (c *CustomLogger) Fatalw(msg string, keysAndValues ...interface{}) {
	c.Logger.Fatal(c.fmtMsg(msg, keysAndValues))
}

func (c *CustomLogger) Panicw(msg string, keysAndValues ...interface{}) {
	c.Logger.Panic(c.fmtMsg(msg, keysAndValues))
}

func (c *CustomLogger) SetLevel(level log.Level) {
	c.Logger.SetLevel(c.convertLevel(level))
}

func (c *CustomLogger) WithContext(ctx context.Context) log.CommonLogger {
	entry := c.Logger.WithContext(ctx)
	return &CustomLogger{
		Logger: entry.Logger,
		entry:  entry,
	}
}

func (c *CustomLogger) convertLevel(level log.Level) logrus.Level {
	switch level {
	case log.LevelTrace:
		return logrus.TraceLevel
	case log.LevelDebug:
		return logrus.DebugLevel
	case log.LevelInfo:
		return logrus.InfoLevel
	case log.LevelWarn:
		return logrus.WarnLevel
	case log.LevelError:
		return logrus.ErrorLevel
	case log.LevelFatal:
		return logrus.FatalLevel
	case log.LevelPanic:
		return logrus.PanicLevel
	}
	panic(fmt.Sprintf("invalid log level: %d", level))
}

func (c *CustomLogger) fmtMsg(msg string, keysAndValues []interface{}) string {
	buf := bytebufferpool.Get()
	defer func() {
		buf.Reset()
		bytebufferpool.Put(buf)
	}()

	// Write msg privateLog buffer
	if msg != "" {
		_, _ = buf.WriteString(msg) //nolint:errcheck // It is fine to ignore the error
	}
	var once sync.Once
	isFirst := true
	// Write keys and values privateLog buffer
	if len(keysAndValues) > 0 {
		if (len(keysAndValues) & 1) == 1 {
			keysAndValues = append(keysAndValues, "KEYVALS UNPAIRED")
		}
		for i := 0; i < len(keysAndValues); i += 2 {
			if msg == "" && isFirst {
				once.Do(func() {
					_, _ = fmt.Fprintf(buf, "%s=%v", keysAndValues[i], keysAndValues[i+1])
					isFirst = false
				})
				continue
			}
			_, _ = fmt.Fprintf(buf, " %s=%v", keysAndValues[i], keysAndValues[i+1])
		}
	}
	return buf.String() //nolint:errcheck // It is fine to ignore the error
}

func init() {
	logger = logrus.New()
	c := &CustomLogger{
		Logger: logger,
	}
	c.SetFormatter(&logrus.JSONFormatter{})
	log.SetLogger(c)
}
