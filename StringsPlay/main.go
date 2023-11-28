package main 
  
import ( 
    "fmt"
) 
  
// Main method 
func main() { 

	
	mp:=[5]int{}
	mp[0]=10

	changeMap(mp)
	fmt.Println(mp)

} 
func changeMap(mp [5]int){
	mp[1]=20
}