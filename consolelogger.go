package mylog

import "fmt"
import "io"
import "os"
import "time"


type ConsoleLogger struct{
    out io.Writer
    level int
    timeFormat string
    prefix string
}

// displays the logs on to the console
func NewConsoleLogger(level int)(l *ConsoleLogger){
    return &ConsoleLogger{os.Stdout,level,TIMEFORMAT,""}
}


func (l *ConsoleLogger) output (msg *Message){
    if msg.level < l.level{
        return 
    }
    buf := []byte{}

    buf = append(buf,msg.time.Format(l.timeFormat)...)
    if l.prefix != "" {
        buf = append(buf,' ')
        buf = append(buf, l.prefix ...)
    }
    buf = append(buf,' ')
    buf = append(buf,'[')
    buf = append(buf, levels[msg.level]...)
    buf = append(buf,']')
    buf = append(buf,' ')
    buf = append(buf,msg.msg...)

    if len(msg.msg) > 0 && msg.msg[len(msg.msg)-1] != '\n'{
        buf = append(buf,'\n')
    }
    l.out.Write(buf)
    if msg.level >= FATAL{
        os.Exit(1)
    }
}

// returns the current log level for this logger
func (l *ConsoleLogger) CurrentLevel()int{
    return l.level
}



// sets the level for this logger
func (l *ConsoleLogger) Level(level int){
    if level >= TRACE && level <= FATAL{
        l.level = level
    }
}


// sets the prefix for this logger
func (l *ConsoleLogger) Prefix (pre string){
    l.prefix = pre
}

// sets the time format for this logger
func (l *ConsoleLogger) TimeFormat(format string){
    l.timeFormat = format
}

// closes the log file
func (l *ConsoleLogger) Close() error{
    // find some way to close multiwriter
    return nil
}

// Log functions
func (l *ConsoleLogger) Fatal (format string,v ... interface{}) {
    l.output(&Message{FATAL,fmt.Sprintf(format,v...),time.Now()})
}

func (l *ConsoleLogger) Error (format string,v ... interface{}) {
    l.output(&Message{ERROR,fmt.Sprintf(format,v...),time.Now()})
}


func (l *ConsoleLogger) Warn(format string,v ... interface{}) {
    l.output(&Message{WARN,fmt.Sprintf(format,v...),time.Now()})
}


func (l *ConsoleLogger) Info(format string,v ... interface{}) {
    l.output(&Message{INFO,fmt.Sprintf(format,v...),time.Now()})
}


func (l *ConsoleLogger) Debug(format string,v ... interface{}) {
    l.output(&Message{DEBUG,fmt.Sprintf(format,v...),time.Now()})
}


func (l *ConsoleLogger) Trace(format string,v ... interface{}) {
    l.output(&Message{TRACE,fmt.Sprintf(format,v...),time.Now()})
}
