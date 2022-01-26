func twoSum(nums []int, target int) []int { 
    tmpMap := make(map[int]int);
    for i := 0 ; i < len(nums) ; i++ {
        if tmpIndex,tmpNum:=tmpMap[target - nums[i]];tmpNum{
            return []int{tmpIndex,i}
        }
        tmpMap[nums[i]] = i ;
    }
    return []int{}
}