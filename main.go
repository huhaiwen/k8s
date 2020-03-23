package main

import(
	"fmt"
	"os"
)

func main(){
	fmt.Println("hello jenkins")
	fmt.Println("BRANCH_NAME:", os.Getenv("branch"))
}
