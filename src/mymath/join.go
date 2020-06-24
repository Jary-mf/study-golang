package mymath

func Join(str1 string,strs...string) string {
	str:=str1
	if len(strs)!=0{
		for _,k :=range strs{
			str+=k
		}
	}
	return str
}