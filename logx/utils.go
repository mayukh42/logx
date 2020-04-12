package logx

import (
    "errors"
    "os"
    "path"
    "strings"
)

const (
    TIME_FORMAT = "2006-01-02 15:04:05.000"
    ERROR       = "ERROR"
    INFO        = "INFO"
    DEBUG       = "DEBUG"
)

// type alias for less verbosity
type AnyT = interface{}
type AnyList = []interface{}

var LEVEL = map[string]int{
    ERROR: 0,
    INFO:  1,
    DEBUG: 2,
}

func createFileHandler(location, name string) (*os.File, error) {
    if _, ok := os.Stat(location); ok != nil {
        // location does not exist. create it
        err := os.MkdirAll(location, os.ModePerm)
        if err != nil {
            return nil, errors.New("Could not create log directory")
        }
    }
    mode := os.FileMode(int(0666))
    f, err := os.OpenFile(path.Join(location, name), os.O_WRONLY|os.O_CREATE|os.O_APPEND, mode)
    if err != nil {
        return nil, errors.New("Could not create log file handler")
    }
    return f, nil
}

func formattedString(n int) string {
    var buf strings.Builder
    buf.WriteString("%s [%s] %s:%d ")
    for i := 0; i < n-4; i++ {
        buf.WriteString("%v ")
    }
    return buf.String()
}
