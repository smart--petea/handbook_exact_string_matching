package main

import (
    "testing"
)

func TestBruteForce(t *testing.T) {
    var pattern, str string
    var res []int

    pattern = "ab"
    str = "salut"
    res = BruteForce(pattern, str)
    if len(res) > 0 {
        t.Fatalf("pattern=%s str=%s no occurences should happend", pattern, str)
    }

    pattern = "ab"
    str = "sabutab"
    res = BruteForce(pattern, str)
    if len(res) != 2 {
        t.Fatalf("pattern=%s str=%s should have 2 occurences", pattern, str)
    }

    if res[0] != 1 {
        t.Fatalf("pattern=%s str=%s first occurence should be 1", pattern, str)
    }

    if res[1] != 5 {
        t.Fatalf("pattern=%s str=%s second occurence should be 5", pattern, str)
    }
}

func TestBrute1Force(t *testing.T) {
    var pattern, str string
    var res []int

    pattern = "ab"
    str = "salut"
    res = BruteForce1(pattern, str)
    if len(res) > 0 {
        t.Fatalf("pattern=%s str=%s no occurences should happend", pattern, str)
    }

    pattern = "ab"
    str = "sabutab"
    res = BruteForce1(pattern, str)
    if len(res) != 2 {
        t.Fatalf("pattern=%s str=%s should have 2 occurences", pattern, str)
    }

    if res[0] != 1 {
        t.Fatalf("pattern=%s str=%s first occurence should be 1", pattern, str)
    }

    if res[1] != 5 {
        t.Fatalf("pattern=%s str=%s second occurence should be 5", pattern, str)
    }
}
