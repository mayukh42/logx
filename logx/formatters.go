package logx

import (
    "encoding/json"
)

type FormatterFn func(string, ...interface{}) string

// TODO: Add caller and line
type jsonFormat struct {
    Timestamp string        `json:"timestamp"`
    Level     string        `json:"level"`
    Message   []interface{} `json:"message"`
}

/** JSONFormatter()
 * ex:
 * 	JSONFormatter("line_message", "the", "answer", "is", 42)
 * yields
 * 	{"timestamp": "2020-04-03 22:11:20.246", "level": "INFO", "message": ["line_message", "the", "answer", "is", 42]}
 */
func JSONFormatterFn(msg string, args ...interface{}) string {
    if len(args) < 2 {
        return ""
    }

    jf := &jsonFormat{
        Timestamp: args[0].(string),
        Level:     args[1].(string),
        Message:   args[2:],
    }

    body, err := json.Marshal(jf)
    if err != nil {
        return ""
    }

    return string(body)
}
