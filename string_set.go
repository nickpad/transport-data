package transportdata

type StringSet map[string]struct{}

func (stringSet StringSet) Add(value string) {
	if _, ok := stringSet[value]; !ok {
		stringSet[value] = struct{}{}
	}
}

func (stringSet StringSet) Remove(value string) {
	delete(stringSet, value)
}

func (stringSet StringSet) FirstValue() string {
	var result string

	for key := range stringSet {
		result = key
		break
	}

	return result
}
