package qvalue

import (
	"math"
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
			if math.IsNaN(qualityFloat) || qualityFloat < qualityFloatMin || qualityFloat > qualityFloatMax {
				if o.ignoreQualityRangeError {
					quality = o.qualityForRangeError
				} else {
					return nil, QualityRangeError
				}
			} else {
				quality = uint(math.Round(qualityFloat * multiplier))
			}
		}
	}

	return &QValue{
		Value:   value,
		Quality: quality,
	}, nil
}

// Parse a quality value series.
func (o *Options) Parse(s string) ([]*QValue, error) {
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

	return result, nil
}

// Parse a quality value series, and sort items according to their quality.
func (o *Options) ParseAndSort(s string) ([]*QValue, error) {
	result, err := o.Parse(s)
	if err != nil {
		return result, err
	}

	Sort(result)

	return result, nil
}
