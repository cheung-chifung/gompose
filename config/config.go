package config

import (
	"io/ioutil"
	"strings"

	"text/template"

	"github.com/fatih/color"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Processes map[string]*Process
}

type Process struct {
	ID      string
	Name    string
	Command string
	Spawn   []string
	*Format
}

type Format struct {
	FGColorStr string             `yaml:"fgcolor"`
	FGColor    color.Attribute    `yaml:"-"`
	BGColorStr string             `yaml:"bgcolor"`
	BGColor    color.Attribute    `yaml:"-"`
	HeaderStr  string             `yaml:"header"`
	Header     *template.Template `yaml:"-"`
}

func Load(path string) (*Config, error) {
	conf := Config{}

	if path == "" {
		path = ".gompose.yaml"
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	if err = yaml.Unmarshal(data, &conf); err != nil {
		return nil, err
	}
	for pid, p := range conf.Processes {
		p.ID = pid
		p.Sanitize()
	}
	return &conf, nil
}

func (p *Process) Sanitize() {
	if len(p.Spawn) == 0 {
		p.Spawn = []string{"/bin/sh", "-c"}
	}

	p.sanitizeFormat()
}

var fgColorTable = map[string]color.Attribute{
	"black":   color.FgBlack,
	"red":     color.FgRed,
	"green":   color.FgGreen,
	"yellow":  color.FgYellow,
	"blue":    color.FgBlue,
	"magenta": color.FgMagenta,
	"cyan":    color.FgCyan,
	"white":   color.FgWhite,
}

var bgColorTable = map[string]color.Attribute{
	"black":   color.BgBlack,
	"red":     color.BgRed,
	"green":   color.BgGreen,
	"yellow":  color.BgYellow,
	"blue":    color.BgBlue,
	"magenta": color.BgMagenta,
	"cyan":    color.BgCyan,
	"white":   color.BgWhite,
}

func (p *Process) sanitizeFormat() {
	if c, ok := fgColorTable[strings.ToLower(p.FGColorStr)]; ok {
		p.FGColor = c
	}
	if c, ok := bgColorTable[strings.ToLower(p.BGColorStr)]; ok {
		p.BGColor = c
	}
	if p.HeaderStr == "" {
		p.HeaderStr = "[{{.Proc.Name}}] "
	}
}
