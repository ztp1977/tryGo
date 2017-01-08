package main

//import (
//	"os"
//
//	log "github.com/Sirupsen/logrus"
//)
//
////func main() {
////
////	log.Out = os.Stderr
////
////	log.WithFields(logrus.Fields{
////		"animal": "walrus",
////		"size":   10,
////	}).Info("A group of walrus emerges from the ocean")
////	fmt.Println("vim-go")
////}
//
//func init() {
//	// Log as JSON instead of the default ASCII formatter.
//	log.SetFormatter(&log.JSONFormatter{})
//
//	// Output to stderr instead of stdout, could also be a file.
//	log.SetOutput(os.Stderr)
//
//	// Only log the warning severity or above.uu
//	log.SetLevel(log.Info)
//}
//
//func main() {
//	log.WithFields(log.Fields{
//		"animal": "walrus",
//		"size":   10,
//	}).Info("A group of walrus emerges from the ocean")
//
//	log.WithFields(log.Fields{
//		"omg":    true,
//		"number": 122,
//	}).Warn("The group's number increased tremendously!")
//
//	log.WithFields(log.Fields{
//		"omg":    true,
//		"number": 100,
//	}).Fatal("The ice breaks!")
//	// A common pattern is to re-use fields between logging statements by re-using
//	// the logrus.Entry returned from WithFields()
//	contextLogger := log.WithFields(log.Fields{
//		"common": "this is a common field",
//		"other":  "I also should be logged always",
//	})
//
//	contextLogger.Info("I'll be logged with common and other field")
//	contextLogger.Info("Me too")
//}

import (
	"os"

	"github.com/Sirupsen/logrus"
)

// Create a new instance of the logger. You can have any number of instances.
var log = logrus.New()

func main() {
	// The API for setting attributes is a little different than the package level
	// exported logger. See Godoc.
	log.Out = os.Stderr
	logrus.SetFormatter(&logrus.TextFormatter{})

	log.WithFields(logrus.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	log.Println("aa bb cc")
}
