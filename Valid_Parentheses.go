func isValid(s string) bool {
    arr := map[string]string{
        ")": "(",
        "}": "{",
		"]": "[",
	}
    stack := make([]string,0,(len(s)/2)+1)
    for _,val := range s{
        switch val{
            case '(','{','[' :
                stack = append(stack,string(val))
                break
            case ')','}',']' :
                if len(stack)>0 && stack[len(stack)-1] == arr[string(val)]{
                    stack = stack[:len(stack)-1]
                }else{
                    return false
                }
        }
    }
    if(len(stack)==0){
        return true
    }else{
        return false
    }
}