package main

import (
	"github.com/sirupsen/logrus"
)

func main() {
	// Set the log level (to show all levels, including debug)
	logrus.SetLevel(logrus.DebugLevel)

	// Debug level
	logrus.Debug("Some additional information")

	// Info level
	logrus.Info("Working...")

	// Warning level
	logrus.Warn("Watch out!")

	// Error level
	logrus.Error("Oh NO!")

	// Critical level (logrus doesn't have critical, but we can use Fatal which logs and then exits)
	logrus.Fatal("x.x")
}

// DEBU[0000] Some additional information                  
// INFO[0000] Working...                                   
// WARN[0000] Watch out!                                   
// ERRO[0000] Oh NO!                                       
// FATA[0000] x.x                                          
// exit status 1
