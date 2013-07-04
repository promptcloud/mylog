package mylog

import "fmt"
import "io"
import "os"
import "time"


type FileLogger struct{
    out io.Writer
    level int
    timeFormat string
    prefix string
}


// create a new file logger with output level as 'level', log filename as 'filename' and an empty prefix
func NewFileLogger (filename string,level int)(l *FileLogger){
    var perms os.FileMode
    perms = 0777
    f,err := os.OpenFile(filename,os.O_CREATE|os.O_APPEND|os.O_RDWR, perms)
    if err != nil{
        fmt.Println(err)
        os.Exit(1)
    }
    return &FileLogger{f,level,TIMEFORMAT,""}
}


func (l *FileLogger) output (msg *Message){
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
    if msg.level > FATAL{
        os.Exit(1)
    }
}


// returns the current log level for this logger
func (l *FileLogger) CurrentLevel()int{
    return l.level
}


// sets the level for this logger
func (l *FileLogger) Level(level int){
    if level >= TRACE && level <= FATAL{
        l.level = level
    }
}


// sets the prefix for this logger
func (l *FileLogger) Prefix (pre string){
    l.prefix = pre
}

// sets the time format for this logger
func (l *FileLogger) TimeFormat(format string){
    l.timeFormat = format
}

// closes the log file
func (l *FileLogger) Close() error{
    // find some way to close multiwriter
    return nil
}

// Log functions
func (l *FileLogger) Fatal (format string,v ... interface{}) {
    l.output(&Message{FATAL,fmt.Sprintf(format,v...),time.Now()})
}

func (l *FileLogger) Error (format string,v ... interface{}) {
    l.output(&Message{ERROR,fmt.Sprintf(format,v...),time.Now()})
}


func (l *FileLogger) Warn(format string,v ... interface{}) {
    l.output(&Message{WARN,fmt.Sprintf(format,v...),time.Now()})
}


func (l *FileLogger) Info(format string,v ... interface{}) {
    l.output(&Message{INFO,fmt.Sprintf(format,v...),time.Now()})
}


func (l *FileLogger) Debug(format string,v ... interface{}) {
    l.output(&Message{DEBUG,fmt.Sprintf(format,v...),time.Now()})
}


func (l *FileLogger) Trace(format string,v ... interface{}) {
    l.output(&Message{TRACE,fmt.Sprintf(format,v...),time.Now()})
}
