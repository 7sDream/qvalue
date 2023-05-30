package qvalue

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

// QValue is a quality value.
//
// The `Index` field is it's origin index.
type QValue struct {
	Value   string
	Quality uint
	Index   int
}

func (o *Options) parseOne(s string) (*QValue, error) {
	qualityIdx := strings.LastIndex(s, ";q=")

	value := s
	quality := o.defaultQuality

	if qualityIdx != -1 {
		value = s[:qualityIdx]
		qualityStr := s[qualityIdx+3:]
		qualityFloat, err := strconv.ParseFloat(qualityStr, 64)
		if err != nil {
			if o.ignoreQualityParseError {
				quality = o.qualityForParseError
			} else {
				return nil, QualityParseError
			}
		} else {
			if math.IsNaN(qualityFloat) || qualityFloat < 0 || qualityFloat > 1 {
				if o.ignoreQualityRangeError {
					quality = o.qualityForRangeError
				} else {
					return nil, QualityRangeError
				}
			} else {
				quality = uint(math.Round(qualityFloat * 1000))
			}
		}
	}

	return &QValue{
		Value:   value,
		Quality: quality,
	}, nil
}

func (o *Options) parse(s string, sorted bool) ([]*QValue, error) {
	parts := strings.Split(s, ",")

	if len(parts) == 0 {
		return nil, nil
	}

	result := make([]*QValue, 0, len(parts))
	for index, part := range parts {
		qValue, err := o.parseOne(strings.TrimSpace(part))
		if err != nil {
			return nil, err
		}
		qValue.Index = index
		result = append(result, qValue)
	}

	if sorted {
		sort.Sort(byQualityThenIndex(result))
	}

	return result, nil
}

// Parse a quality value series.
func (o *Options) Parse(s string) ([]*QValue, error) {
	return o.parse(s, false)
}

// Parse a quality value series, and sort items according to their quality.
func (o *Options) ParseAndSort(s string) ([]*QValue, error) {
	return o.parse(s, true)
}
