package logx

import (
    "encoding/json"
    "fmt"
)

type FormatterFn func(...AnyT) string

type meta struct {
    Timestamp  string `json:"timestamp"`
    Level      string `json:"level"`
    CallerFile string `json:"file,omitempty"`
    CallerLine int    `json:"line,omitempty"`
}

// TODO: Add caller and line
type jsonFormat struct {
    meta
    Message []AnyT `json:"message"`
}

/** JSONFormatter()
 * ex:
 * 	JSONFormatter("line_message", "the", "answer", "is", 42)
 * yields
 * 	{"timestamp": "2020-04-03 22:11:20.246", "level": "INFO", "message": ["line_message", "the", "answer", "is", 42]}
 */
func JSONFormatterFn(args ...AnyT) string {
    if len(args) < 2 {
        return ""
    }

    m, ok := args[0].(meta)
    if !ok {
        return ""
    }

    jf := &jsonFormat{
        Message: args[1:],
        meta:    m,
    }

    body, err := json.Marshal(jf)
    if err != nil {
        return ""
    }

    return string(body)
}

func BaseFormatterFn(args ...AnyT) string {
    m, ok := args[0].(meta)
    if !ok {
        return ""
    }

    _as := AnyList{m.Timestamp, m.Level, m.CallerFile, m.CallerLine}
    fs := formattedString(len(_as) + len(args) - 1)

    for _, a := range args[1:] {
        _as = append(_as, a)
    }
    return fmt.Sprintf(fs, _as...)
}
