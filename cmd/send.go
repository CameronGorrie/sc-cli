package cmd

import (
	"errors"
	"flag"
	"fmt"
	"strings"

	"github.com/CameronGorrie/sc"
)

type Send struct {
	ugenList string
	client   *sc.Client
}

func (s *Send) Run(args []string) error {
	if len(args[1:]) == 0 {
		return errors.New("no arguments provided to send ")
	}

	fs := flag.NewFlagSet("send", flag.ContinueOnError)
	fs.StringVar(&s.ugenList, "ugens", "", "A comma delimited list of Ugen names")

	if err := fs.Parse(args[1:]); err != nil {
		return err
	}

	ugenNames := strings.Split(s.ugenList, ",")
	for _, name := range ugenNames {
		// do a lookup in the complete dictionary of sc-ugens package
		if f, ok := ugens.Ugens[name]; !ok {
			errMsg := fmt.Sprintf("no matching ugen found for name %s ", name)

			return errors.New(errMsg)
		} else {
			return s.client.SendDef(sc.NewSynthdef(name, f))
		}
	}

	return nil
}
