package main
import "fmt"

func main(){

	//For loop with condition
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	//For loop without condition (runs forever)
	for i := 0; ; i++ {
		fmt.Println(i)
		if i==1000{
			break
		}
	}
}