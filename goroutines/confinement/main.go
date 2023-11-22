package main

import (
	"fmt"
	"sync"
)

func main(){

	nums:=[]int{2,3,4,5}
	var wg sync.WaitGroup

	for i:=0;i<len(nums);i++{
		wg.Add(1)
		go processData(&wg,&nums[i])
	}
	wg.Wait()
	fmt.Println(nums)
}

func processData(wg *sync.WaitGroup,i *int){
	defer wg.Done()
	*i = *i**i
}