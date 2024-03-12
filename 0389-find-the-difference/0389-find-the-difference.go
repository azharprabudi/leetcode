func findTheDifference(s string, t string) byte {
    var sum1, sum2 int
    
    for _, char := range s {
        sum1 += int(char)
    }
    
    for _, char := range t {
        sum2 += int(char)
    }
    
    return byte(sum2 - sum1)
}