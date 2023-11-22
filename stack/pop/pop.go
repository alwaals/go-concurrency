package pop

func Pop(s []string)[]string{
	return s[:len(s)-1]
}
func Push(s *[]string,str string){
	*s=append(*s, str)
}