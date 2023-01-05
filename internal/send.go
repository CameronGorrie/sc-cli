package cmd

import (
	"errors"
	"flag"
	"fmt"
	"strings"

	"github.com/CameronGorrie/sc"
	"github.com/CameronGorrie/ugens"
)

type Send struct {
	port     int
	ugenList string
}

func (s *Send) Run(args []string) error {
	if len(args[1:]) == 0 {
		return errors.New("no arguments provided to send ")
	}

	fs := flag.NewFlagSet("send", flag.ContinueOnError)
	fs.StringVar(&s.ugenList, "ugens", "", "A comma delimited list of Ugen names")
	fs.IntVar(&s.port, "u", 57120, "UDP port")

	if err := fs.Parse(args[1:]); err != nil {
		return err
	}

	c, err := NewClient(s.port)
	if err != nil {
		return err
	}

	ugenNames := strings.Split(s.ugenList, ",")
	for _, name := range ugenNames {
		if f, ok := ugens.CompleteDictionary[name]; !ok {
			errMsg := fmt.Sprintf("no matching ugen found for name %s ", name)

			return errors.New(errMsg)
		} else {
			return c.SendDef(sc.NewSynthdef(name, f))
		}
	}

	return nil
}
