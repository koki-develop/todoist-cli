package util

func InterfacesToStrings(is []interface{}) []string {
	ss := make([]string, len(is))
	for i, v := range is {
		ss[i] = v.(string)
	}
	return ss
}

func StringsToInterfaces(ss []string) []interface{} {
	is := make([]interface{}, len(ss))
	for i, v := range ss {
		is[i] = v
	}
	return is
}
