func longestCommonPrefix(strs []string) string {
    min :=200
    for i:=0; i<len(strs);i++{
        if len(strs[i]) < min {
            min = len(strs[i])
        }
    }
    for i:=0; i<min;i++{
        for j:=0;j<len(strs);j++{
            if strs[0][i] != strs[j][i]{
                min = i
            }
        }
    }
    return strs[0][:min]
}