package main
import(
	"fmt"
)
func main(){
	xi:=[]int{1,2,3,4,5,6,7,8}
	fmt.Println(foo(xi...))
}
func foo(ii ...int)int{
	t:=0
	for_,v:=range ii {
		t += v
	}
	return t
}