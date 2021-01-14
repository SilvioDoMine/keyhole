// Copyright 2020 Kuei-chun Chen. All rights reserved.

package mdb

import (
	"fmt"
	"strings"
	"time"
)

// Logger stores logger info
type Logger struct {
	Collected time.Time `json:"collected" bson:"collected"`
	Warnings  []string  `json:"warnings" bson:"warnings"`
	Logs      []string  `json:"logs" bson:"logs"`
	Params    string    `json:"params" bson:"params"`
	Version   string    `json:"version" bson:"version"`

	nocolor bool
}

// NewLogger returns Logger
func NewLogger(fullVersion string, params string) *Logger {
	p := Logger{Version: fullVersion, Params: params, Warnings: []string{}}
	p.Collected = time.Now()
	p.Logs = []string{fmt.Sprintf(`%v I %v begins`, p.Collected.Format(time.RFC3339), fullVersion)}
	return &p
}

// SetNoColor set nocolor flag
func (p *Logger) SetNoColor(nocolor bool) {
	p.nocolor = nocolor
}

// Add adds a message
func (p *Logger) Add(message string) {
	str := fmt.Sprintf(`%v I %v`, time.Now().Format(time.RFC3339), message)
	p.Logs = append(p.Logs, str)
}

// Warning adds an warning message
func (p *Logger) Warning(message string) {
	p.Warnings = append(p.Warnings, message)
	fmt.Println(CodeRed, "*", message, CodeDefault)
}

// Info adds and prints a message
func (p *Logger) Error(v ...interface{}) {
	p.print("E", v)
}

// Info adds and prints a message
func (p *Logger) Info(v ...interface{}) {
	p.print("I", v)
}

// Warn adds and prints a message
func (p *Logger) Warn(v ...interface{}) {
	p.print("W", v)
}

// Log adds and prints a message
func (p *Logger) Log(v ...interface{}) {
	p.print("I", v)
}

// Log adds and prints a message
func (p *Logger) print(indicator string, v ...interface{}) {
	message := fmt.Sprint(v...)
	if len(message) > 1 {
		message = message[1 : len(message)-1]
	}
	str := fmt.Sprintf(`%v %v %v`, time.Now().Format(time.RFC3339), indicator, message)
	p.Logs = append(p.Logs, str)
	fmt.Println(str)
}

// Print prints keyhole info
func (p *Logger) Print() string {
	if p == nil {
		return ""
	}
	strs := []string{fmt.Sprintf(`{ keyhole: { version: "%v", args: "%v" } }`, p.Version, p.Params)}
	strs = append(strs, p.Logs...)
	if len(p.Warnings) > 0 {
		strs = append(strs, "\nWarnings:")
		for _, warning := range p.Warnings {
			if p.nocolor {
				strs = append(strs, warning)
			} else {
				strs = append(strs, fmt.Sprintf(`%v%v%v`, CodeRed, warning, CodeDefault))
			}
		}
		strs = append(strs, "")
	}
	return strings.Join(strs, "\n")
}
