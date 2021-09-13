package piscine

func IterativeFactorial(nb int) int {

	
	a:= nb-1
 
for i := 1 ; i <= a ; i++ {
	nb *= i 
	

}
return nb
}
