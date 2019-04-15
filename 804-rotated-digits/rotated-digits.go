func rotatedDigits(N int) int {
    res := 0
    for i := 1; i <= N; i++ {
        str1 := strconv.Itoa(i)
        tmp := false
        for j := 0; j < len(str1); j++ {
            if str1[j] == '3' || str1[j] == '4' || str1[j] == '7' {
                tmp = false
                break
            } else if str1[j] == '2' || str1[j] == '5' || str1[j] == '6' || str1[j] == '9' {
                tmp = true
            }
        }
        if tmp {
            res += 1
        }
    }
    return res
}