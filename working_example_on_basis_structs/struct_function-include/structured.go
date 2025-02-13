package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"

	"github.com/apex/log"
	"github.com/apex/log/handlers/text"
)

func main() {
	fmt.Println("Logrus:")
	Logrus()

	fmt.Println()
	fmt.Println("Apex:")
	Apex()
}

// ThrowError throws an error that we'll trace
func ThrowError() error {
	err := errors.New("a crazy failure")
	log.WithField("id", "123").Trace("ThrowError").Stop(&err)
	return err
}

// CustomHandler splits to two streams
type CustomHandler struct {
	id      string
	handler log.Handler
}

// HandleLog adds a hook and does the emitting
func (h *CustomHandler) HandleLog(e *log.Entry) error {
	e.WithField("id", h.id)
	return h.handler.HandleLog(e)
}

// Apex has a number of useful tricks
func Apex() {
	log.SetHandler(&CustomHandler{"123", text.New(os.Stdout)})
	err := ThrowError()

	//With error convenience function
	log.WithError(err).Error("an error occurred")
}

//---

// Hook will implement the logrus
// hook interface
type Hook struct {
	id string
}

// Fire will trigger whenever you log
func (hook *Hook) Fire(entry *logrus.Entry) error {
	entry.Data["id"] = hook.id
	return nil
}

// Levels is what levels this hook will fire on
func (hook *Hook) Levels() []logrus.Level {
	return logrus.AllLevels
}

// Logrus demonstrates some basic logrus functionality
func Logrus() {
	// we're emitting in text format
	logrus.SetFormatter(&logrus.TextFormatter{})
	logrus.SetLevel(logrus.InfoLevel)
	logrus.AddHook(&Hook{"123"})

	fields := logrus.Fields{}
	fields["success"] = true
	fields["complex_struct"] = struct {
		Event string
		When  string
	}{"Something happened", "Just now"}

	x := logrus.WithFields(fields)
	x.Warn("warning!")
	x.Error("error!")
}

// Logrus:
// WARN[0000] warning!                                      complex_struct="{Something happened Just now}" id=123 success=true
// ERRO[0000] error!                                        complex_struct="{Something happened Just now}" id=123 success=true

// Apex:
//   INFO[0000] ThrowError                id=123
//  ERROR[0000] ThrowError                duration=0 error=a crazy failure id=123
//  ERROR[0000] an error occurred         error=a crazy failure
