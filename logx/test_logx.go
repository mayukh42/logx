package logx

const (
    LOG_LOCATION = "/home/rufus/dev/log"
)

func GetLogger() *Logger {
    logger := NewLogger().
        SetMaxLevel(DEBUG).
        SetFormatter(JSONFormatterFn).
        AddFileHandler(LOG_LOCATION, "test.log").
        ConsoleOut(true)

    return logger
}

func TestLog() {
    logger := GetLogger()
    // should always close, else all active writes may not complete before program exits
    defer logger.Close()

    otherLogger := GetLogger()
    defer otherLogger.Close()

    message := AnyList{"Could not validate request params", "the", "answer", "is", 42}
    logger.Infof(message...)

    // logger.Debugf(message...)

    // logger.Errorf(message...)
    otherLogger.Errorf(message...)
}
