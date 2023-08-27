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

	//While loop i.e for loop with only condition
	j := 0
	for  j < 10 { 
		fmt.Println(j)
		j++
	}
}