func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    
    ss, tt := countChar(s), countChar(t)
    for char, count := range ss {
        tcount, valid := tt[char]
        if !valid {
            return false
        }
        
        if count != tcount {
            return false
        }
    }
    
    return true
}

func countChar(s string) map[rune]int {
    chars := make(map[rune]int)  
    for _, char := range s {
        _, valid := chars[char]
        if !valid {
            chars[char] = 0
        }
        
        chars[char] += 1
    }
    
    return chars
}