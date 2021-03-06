package logx

import (
    "errors"
    "fmt"
    "os"
    "path/filepath"
    "runtime"
    "strings"
    "sync"
    "time"
)

/** A goroutines-based logger that works with signals
 * TODO:
 *  use signals/ channels
 *  custom (and multiple) writers
 *  test race conditions
 */
type Logger struct {
    defaultLevel string
    maxLevel     string
    timeFormat   string
    fileHandler  *os.File
    formatter    FormatterFn
    console      bool
    wg           *sync.WaitGroup
}

// only Logger should be able to access
// TODO: not implemented in 0.20.3 version
type channels struct {
    msg  chan<- jsonFormat
    kill chan<- struct{}
}

// default formatter: BaseFormatterFn [fmt.Sprintf()]
func NewLogger() *Logger {
    logger := new(Logger)
    logger.defaultLevel = INFO
    logger.timeFormat = TIME_FORMAT
    logger.maxLevel = logger.defaultLevel
    logger.formatter = BaseFormatterFn
    logger.wg = &sync.WaitGroup{}
    return logger
}

func (log *Logger) ConsoleOut(flag bool) *Logger {
    log.console = flag
    return log
}

func (log *Logger) SetMaxLevel(level string) *Logger {
    level = strings.ToUpper(level)
    if _, ok := LEVEL[level]; ok {
        log.maxLevel = level
    }

    return log
}

func (log *Logger) SetTimeFormat(tf string) *Logger {
    if len([]rune(tf)) > 0 {
        // not a proper check, but if format string is wrong, output will be wrong too w/o loss of correctness
        log.timeFormat = tf
    }

    return log
}

func (log *Logger) SetFormatter(fn FormatterFn) *Logger {
    if fn != nil {
        log.formatter = fn
    }
    return log
}

func (log *Logger) AddFileHandler(location, name string) *Logger {
    if log.fileHandler != nil {
        // already a file handler was added, so skip
        return log
    }

    f, err := createFileHandler(location, name)
    if err == nil {
        log.fileHandler = f
    } else {
        fmt.Printf("%s\n", err.Error())
    }

    return log
}

func (log *Logger) Close() error {
    // wait for all writes to end
    // TODO: use kill channel
    log.wg.Wait()

    if log.fileHandler != nil {
        return log.fileHandler.Close()
    }
    return nil
}

func (log *Logger) log(m meta, args ...AnyT) error {
    // this is always called in a goroutine
    defer log.wg.Done()

    as_ := AnyList{m}
    for _, a := range args {
        as_ = append(as_, a)
    }

    content := log.formatter(as_...)
    if content == "" {
        return errors.New("Nothing to log")
    }

    // append newline
    line := fmt.Sprintf("%s\n", content)
    if log.fileHandler != nil {
        // use writer interface
        log.fileHandler.WriteString(line)
    }

    // Console out if set to true
    if log.console {
        fmt.Printf(line)
    }

    return nil
}

func (log *Logger) controller(level string, args ...AnyT) error {
    // use this for select channels
    m := meta{
        Timestamp: time.Now().Format(TIME_FORMAT),
        Level:     level,
    }

    // Add caller info
    if _, f, n, ok := runtime.Caller(2); ok {
        m.CallerFile = filepath.Base(f)
        m.CallerLine = n
    }

    log.wg.Add(1)
    go log.log(m, args...)

    return nil
}

func (log *Logger) Errorf(args ...AnyT) error {
    // Errorf is always logged
    log.controller(ERROR, args...)
    return nil
}

func (log *Logger) Infof(args ...AnyT) error {
    if LEVEL[INFO] <= LEVEL[log.maxLevel] {
        log.controller(INFO, args...)
    }

    return nil
}

func (log *Logger) Debugf(args ...AnyT) error {
    if LEVEL[DEBUG] <= LEVEL[log.maxLevel] {
        log.controller(DEBUG, args...)
    }

    return nil
}
