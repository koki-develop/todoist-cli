package util

func InterfacesToStrings(is []interface{}) []string {
	ss := make([]string, len(is))
	for i, v := range is {
		ss[i] = v.(string)
	}
	return ss
}
