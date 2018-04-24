package proc

import (
	"io"
	"os"
	"os/exec"
	"sync"

	"github.com/keekun/gompose/config"
	"github.com/keekun/gompose/logger"
	"github.com/kr/pty"
)

type Processes struct {
	processes map[string]*Process
	Config    *config.Config
	*sync.Mutex
}

type Process struct {
	Config *config.Process
	Input  io.Reader
	Output io.Writer
}

func NewProcesses(conf *config.Config, output io.Writer) *Processes {
	ps := &Processes{
		processes: make(map[string]*Process),
		Config:    conf,
		Mutex:     new(sync.Mutex),
	}
	for _, pConf := range conf.Processes {
		ps.add(pConf, output)
	}
	return ps
}

func (ps *Processes) Spawn() error {
	for _, p := range ps.processes {
		err := p.Spawn()
		if err != nil {
			return err
		}
	}
	return nil
}

func (ps *Processes) add(conf *config.Process, output io.Writer) (*Process, error) {
	logger, err := logger.New(conf, output)
	if err != nil {
		return nil, err
	}
	p := &Process{
		Config: conf,
		Output: logger,
		Input:  nil,
	}
	ps.Lock()
	ps.processes[conf.ID] = p
	ps.Unlock()
	return p, nil
}

func (p *Process) Spawn() error {
	cmdName := p.Config.Spawn[0]
	cmdArgs := append(p.Config.Spawn[1:], p.Config.Command)
	cmd := exec.Command(cmdName, cmdArgs...)
	cmd.Env = os.Environ()

	f, err := pty.Start(cmd)
	if err != nil {
		return err
	}
	_, err = io.Copy(p.Output, f)
	return err
}
