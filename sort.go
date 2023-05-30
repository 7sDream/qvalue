package qvalue

type byQualityThenIndex []*QValue

func (a byQualityThenIndex) Len() int      { return len(a) }
func (a byQualityThenIndex) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a byQualityThenIndex) Less(i, j int) bool {
	if a[i].Quality > a[j].Quality {
		return true
	} else if i < j {
		return true
	}

	return false
}
