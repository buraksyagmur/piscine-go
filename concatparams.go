package piscine

func ConcatParams(args []string) string {
	result := ""
	a := len(args)
	for i := 0; i < a; i++ {

		if i != 0 && string(args[i]) != " " {
			result += "\n"
		}
		{
			result += string(args[i])
		}
	}
	return result
}
