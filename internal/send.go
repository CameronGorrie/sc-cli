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
	ugenList string
}

func (s *Send) Run(c *sc.Client) error {
	ugenNames := strings.Split(s.ugenList, ",")
	for _, name := range ugenNames {
		if f, ok := ugens.CompleteDictionary[name]; !ok {
			errMsg := fmt.Sprintf("no matching ugen found for name %s ", name)

			return errors.New(errMsg)
		} else {
			if err := c.SendDef(sc.NewSynthdef(name, f)); err != nil {
				return err
			}

			fmt.Printf("%s ugen sent to server\n", name)
		}
	}

	return nil
}

func (s *Send) ParseFlags(fs *flag.FlagSet, args []string) error {
	if len(args) == 0 {
		return errors.New("no arguments provided to send ")
	}

	fs.StringVar(&s.ugenList, "ugens", "", "A comma delimited list of Ugen names")

	if err := fs.Parse(args[0:]); err != nil {
		return err
	}

	return nil
}
