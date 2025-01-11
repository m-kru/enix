package regex

import (
	"fmt"
	"regexp"
)

type RegexJSON struct {
	Regex              string
	NegativeLookBehind string
	PositiveLookBehind string
	NegativeLookAhead  string
	PositiveLookAhead  string
}

func (rj RegexJSON) ToRegex() (*Regex, error) {
	if rj.Regex == "" {
		return nil, nil
	}

	var err error

	var re *regexp.Regexp
	re, err = regexp.Compile(rj.Regex)
	if err != nil {
		return nil, fmt.Errorf("can't compile regex: %v", err)
	}

	var nlb *regexp.Regexp
	if rj.NegativeLookBehind != "" {
		nlb, err = regexp.Compile(rj.NegativeLookBehind)
		if err != nil {
			return nil, fmt.Errorf("can't compile negative lookbehind: %v", err)
		}
	}

	var plb *regexp.Regexp
	if rj.PositiveLookBehind != "" {
		plb, err = regexp.Compile(rj.PositiveLookBehind)
		if err != nil {
			return nil, fmt.Errorf("can't compile positive lookbehind: %v", err)
		}
	}

	var nla *regexp.Regexp
	if rj.NegativeLookAhead != "" {
		nla, err = regexp.Compile(rj.NegativeLookAhead)
		if err != nil {
			return nil, fmt.Errorf("can't compile negative lookahead: %v", err)
		}
	}

	var pla *regexp.Regexp
	if rj.PositiveLookAhead != "" {
		pla, err = regexp.Compile(rj.PositiveLookAhead)
		if err != nil {
			return nil, fmt.Errorf("can't compile positive lookahead: %v", err)
		}
	}

	return &Regex{
		Regex:              re,
		NegativeLookBehind: nlb,
		PositiveLookBehind: plb,
		NegativeLookAhead:  nla,
		PositiveLookAhead:  pla,
	}, nil
}
