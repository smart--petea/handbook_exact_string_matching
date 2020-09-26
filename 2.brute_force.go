package main

func BruteForce(pattern string, str string) []int {
    results := make([]int, 0)
    strLen := len(str)
    patternLen := len(pattern)

    for i := 0; i <= strLen - patternLen; i++ {
        var j int
        for j = 0; j < patternLen && pattern[j] == str[i + j]; j++ {}
        if j >= patternLen {
            results = append(results, i)
        }
    }

    return results
}

func BruteForce1(pattern string, str string) []int {
    results := make([]int, 0)
    strLen := len(str)
    patternLen := len(pattern)

    for i := 0; i <= strLen - patternLen; i++ {
        if str[i: i + patternLen] == pattern {
            results = append(results, i)
        }
    }

    return results
}
