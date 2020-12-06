package main

import "fmt"

func main(){
	s := "()[]{}"
	res := isValid(s)
	fmt.Println(res)
}


func isValid(s string) bool{
	if s == ""{
		return true
	}

	if len(s)%2 == 1{
		return false
	}

	keyMap := map[string]string{
		"}": "{",
		"]": "[",
		")": "(",
	}
	var table []string
	for i:=0;i<len(s);i++{
		if len(table) > 0{
			tmp,ok := keyMap[string(s[i])]
			if ok{
				top := table[len(table) - 1]
				if top == tmp {
					table = table[:len(table) - 1]
					continue
				}
			}
		}
		table = append(table,string(s[i]))
	}
	return len(table) == 0
}