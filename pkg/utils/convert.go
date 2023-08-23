package utils

func Int8ToString(arr []int8) string {
	var res []byte

	for _, v := range arr {
		if v == 0 {
			break
		}
		res = append(res, byte(v))
	}
	return string(res)
}
