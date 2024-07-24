# logger


usage 


```golang
package main

import "github.com/kuluce/pkg/logger"

const port = 5555

func main() {
	logger.InitLog("order")
	logger.Infof("order server start at :%d", port)

    logger.Info("this is a info log")
	logger.Infof("this is a infof log %d", 1)
	logger.InfoWithCTX(context.WithValue(context.Background(), key("key"), "value"), "this is a info log with context")

	logger.Debug("this is a debug log")
	logger.Debugf("this is a debugf log %d", 1)
	logger.DebugWithCTX(context.WithValue(context.Background(), key("key"), "value"), "this is a debug log with context")

	logger.Warn("this is a warn log")
	logger.Warnf("this is a warnf log %d", 1)
	logger.WarnWithCTX(context.WithValue(context.Background(), key("key"), "value"), "this is a warn log with context")

	logger.Error("this is a error log")
	logger.Errorf("this is a errorf log %d", 1)
	logger.ErrorWithCTX(context.WithValue(context.Background(), key("key"), "value"), "this is a error log with context")

	logger.Fatal("this is a fatal log")
	logger.Fatalf("this is a fatalf log %d", 1)
	logger.FatalWithCTX(context.WithValue(context.Background(), key("key"), "value"), "this is a fatal log with context")
   

}

```

