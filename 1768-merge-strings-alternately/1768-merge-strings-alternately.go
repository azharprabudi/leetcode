func mergeAlternately(word1 string, word2 string) string {
    var lastPos int
    var mergedString string
    
    for pos, char := range word1 {
        if pos <= len(word2) - 1 {
            mergedString = fmt.Sprintf("%s%c%c", mergedString, char, word2[pos])
        } else {
            mergedString = fmt.Sprintf("%s%c", mergedString, char)
        }

        lastPos = pos
    }
    
    if lastPos < len(word2) - 1 {
        mergedString = fmt.Sprintf("%s%s", mergedString, word2[lastPos + 1:len(word2)])
    }

    return mergedString
}