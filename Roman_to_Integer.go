func romanToInt(s string) int {
    arr := map[string]int{"I":1,"V":5,"X":10,"L":50,"C":100,"D":500,"M":1000}
    tmp := arr[string(s[0])]
    sum := 0
    for i:=1;i<len(s);i++{
        if(tmp<arr[string(s[i])]){
            sum-=tmp
        }else{
            sum+=tmp
        }
        tmp = arr[string(s[i])]
    }
    return sum+tmp
}