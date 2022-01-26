func isPalindrome(x int) bool {
    str := strconv.Itoa(x)
    left := 0
    right := len(str)-1
    for left < right{
        if(str[left] != str[right]){
            return false
        }
        left+=1
        right-=1
    }
    return true
}