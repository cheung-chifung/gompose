package logger

import (
	"bytes"
	"io"
	"sync"
	"time"

	"text/template"

	"github.com/fatih/color"
	"github.com/keekun/gompose/bufferpool"
	"github.com/keekun/gompose/config"
)

type Logger struct {
	conf   *config.Process
	color  *color.Color
	writer io.Writer
	header *template.Template
	*sync.Mutex
}

type headerParams struct {
	Proc *config.Process
	Now  time.Time
}

// FIXME handle lock in a better way
var lock *sync.Mutex

func init() {
	lock = new(sync.Mutex)
}

func New(conf *config.Process, writer io.Writer) (*Logger, error) {
	var err error
	l := &Logger{
		conf:   conf,
		Mutex:  new(sync.Mutex),
		writer: writer,
	}
	l.color = color.New(conf.FGColor, conf.BGColor)
	l.header, err = template.New(conf.ID).Parse(l.conf.HeaderStr)
	if err != nil {
		return nil, err
	}
	return l, nil
}

func (l *Logger) Write(p []byte) (int, error) {
	buf := bytes.NewBuffer(p)
	for {
		tmplBuf := bufferpool.Get()
		defer bufferpool.Free(tmplBuf)

		line, err := buf.ReadBytes('\n')
		if len(line) > 1 {
			err := l.header.Execute(tmplBuf, &headerParams{
				Proc: l.conf,
				Now:  time.Now(),
			})
			if err != nil {
				return 0, err
			}
			lock.Lock()
			l.color.Fprint(l.writer, tmplBuf.String())
			l.writer.Write(line)
			lock.Unlock()
		}
		if err != nil {
			break
		}
	}
	return len(p), nil
}
