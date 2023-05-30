// Package qvalue is a parser for [quality values](https://developer.mozilla.org/en-US/docs/Glossary/Quality_values).
package qvalue

import (
	"errors"
)

var (
	// Error to indicate there is invalid quality in provided string
	QualityParseError = errors.New("Quality is not a valid number")
	// Error io indicate a quality is out of range
	QualityRangeError = errors.New("Quality should be in range of [0.0, 1.0]")
)

// Options is configure for this parser.
type Options struct {
	ignoreQualityParseError bool
	qualityForParseError    uint

	ignoreQualityRangeError bool
	qualityForRangeError    uint

	defaultQuality uint
}

// Create a new default parser option.
// The default behavior is report any error and default quality is 1.0.
func NewOptions() *Options {
	return &Options{
		ignoreQualityParseError: false,
		ignoreQualityRangeError: false,
		defaultQuality:          1000,
	}
}

// IgnoreQualityParseError makes this parse option ignore `QualityParseError`,
// and set the item's quality to `value` if this error happened.
func (o *Options) IgnoreQualityParseError(value uint) *Options {
	if value > 1000 {
		panic("quality value can't bigger then 1000")
	}

	o.ignoreQualityParseError = true
	o.qualityForParseError = value

	return o
}

// IgnoreQualityParseError makes this parse option ignore `QualityRangeError`
// and set the item's quality to `value` if this error happened.
func (o *Options) IgnoreQualityRangeError(value uint) *Options {
	if value > 1000 {
		panic("quality value can't bigger then 1000")
	}

	o.ignoreQualityRangeError = true
	o.qualityForRangeError = value

	return o
}

// SetDefaultQuality set the default quality if a item do not have it.
// The default is 1.0, according to the spec.
func (o *Options) SetDefaultQuality(value uint) *Options {
	if value > 100 {
		panic("quality value can't bigger then 1000")
	}

	o.defaultQuality = value

	return o
}
