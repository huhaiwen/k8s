package main

import(
	"fmt"
	"os"
)

func main(){
	fmt.println("hello jenkins")
	fmt.Println("BRANCH_NAME:", os.Getenv("branch"))
}
