package piscine
import "sort"
func Abort(a, b, c, d, e int) int {


s:= []int{a,b,c,d,e}

sort.Ints(s)

return s[2]
}
