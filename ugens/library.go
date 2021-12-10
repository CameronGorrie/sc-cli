package ugens

import (
	"github.com/CameronGorrie/sc"
)

var Ugens = map[string]sc.UgenFunc{
	"simple_sine": func(p sc.Params) sc.Ugen {
		var (
			freq = p.Add("freq", 440)
			gate = p.Add("gate", 1)
			bend = p.Add("bend", 1)
		)

		env := sc.EnvGen{
			Gate: gate,
			Done: sc.FreeEnclosing,
			Env: sc.EnvADSR{
				A: sc.C(0.01),
				D: sc.C(1),
				S: sc.C(1),
				R: sc.C(1),
			},
		}.Rate(sc.KR)

		sig := sc.SinOsc{
			Freq: freq.Mul(bend),
		}.Rate(sc.AR).Mul(env)

		return sc.Out{
			Bus:      sc.C(0),
			Channels: sc.Multi(sig, sig),
		}.Rate(sc.AR)
	},
}
