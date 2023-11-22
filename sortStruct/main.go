package main

type User struct{
	Name string
	Age int
	Salary int
}
func main(){

}
func (u *User) Len() int{
	return len(u.Name)

}
func (u *User) Less(i,j int){

}
// func (u *User) Swap(i,j int) User{
//     return u.Age<u
// }
