package mylog

import "fmt"
import "time"
import "os"

const (
    TRACE = iota
    DEBUG
    INFO
    WARN
    ERROR
    FATAL
    LOG

    TIMEFORMAT = "2006/01/02 15:04:05"
)

var (
    default_level = os.Getenv("SDF_LOG_LEVEL")
    stdLog = NewConsoleLogger(Level_from_string(default_level))             // default is a console logger
    levels = [...]string{"TRACE","DEBUG","INFO","WARN","ERROR","FATAL"}
    timeformat = TIMEFORMAT
)

type Logger interface{
    Fatal(string,...interface{})
    Error(string,...interface{})
    Warn(string,...interface{})
    Info(string,...interface{})
    Debug(string,...interface{})
    Trace(string,...interface{})
    Level(int)
    Prefix(string)
    TimeFormat(string)
    Close() error
    output(msg *Message)
}

type Message struct{
    level int
    msg string
    time time.Time
}

// displays the levels and their corresponding integers
func DisplayLevels(){
    for index,level := range levels{
       fmt.Println(level,"\t",index)
    }
}


// returns the current log level for default logger
func CurrentLevel()int{
    return stdLog.level
}

// Sets the output level for the default logger
func Level(level int) {
	stdLog.Level(level)
}

// Sets the output level for logger from a string
func Level_from_string(level string)int{
    switch level{
        case "TRACE": return TRACE
        case "DEBUG": return DEBUG 
        case "INFO" : return INFO
        case "WARN" : return WARN 
        case "ERROR": return ERROR
        case "FATAL": return FATAL
    }
    return INFO
}

// Sets the time format for the default logger
func TimeFormat(f string) {
	stdLog.TimeFormat(f)
}

// Logging functions
func Fatal(format string, v ...interface{}) {
	stdLog.output(&Message{FATAL, fmt.Sprintf(format, v...), time.Now()})
}

func Error(format string, v ...interface{}) {
	stdLog.output(&Message{ERROR, fmt.Sprintf(format, v...), time.Now()})
}

func Warn(format string, v ...interface{}) {
	stdLog.output(&Message{WARN, fmt.Sprintf(format, v...), time.Now()})
}

func Info(format string, v ...interface{}) {
	stdLog.output(&Message{INFO, fmt.Sprintf(format, v...), time.Now()})
}

func Debug(format string, v ...interface{}) {
	stdLog.output(&Message{DEBUG, fmt.Sprintf(format, v...), time.Now()})
}

func Trace(format string, v ...interface{}) {
	stdLog.output(&Message{TRACE, fmt.Sprintf(format, v...), time.Now()})
}
