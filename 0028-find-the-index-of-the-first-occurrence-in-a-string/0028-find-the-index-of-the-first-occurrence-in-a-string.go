func strStr(haystack string, needle string) int {
    nl := len(needle)
    
    for pos, _ := range haystack {
        limit := pos + nl
        if limit > len(haystack) {
            break
        }
        
        val := haystack[pos:limit]
        if val == needle {
            return pos
        }
    }
    
    return -1
}