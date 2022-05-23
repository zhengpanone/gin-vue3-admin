package utils

type Rules map[string][]string

func NotEmpty() string {
	return "notEmpty"
}

/*func Verify(st interface{}, roleMap Rules) (err error) {
	compareMap := map[string]bool{
		"lt": true,
		"le": true,
		"eq": true,
		"ne": true,
		"ge": true,
		"gt": true,
	}

	type:=reflect.TypeOf(st)
	val:=reflect.ValueOf(st) // 获取reflect.Type 类型
}*/
