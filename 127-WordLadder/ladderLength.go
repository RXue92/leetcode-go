/*
Given two words (beginWord and endWord), and a dictionary's word list, find the length of shortest transformation sequence from beginWord to endWord, such that:

Only one letter can be changed at a time.
Each transformed word must exist in the word list. Note that beginWord is not a transformed word.
Note:

Return 0 if there is no such transformation sequence.
All words have the same length.
All words contain only lowercase alphabetic characters.
You may assume no duplicates in the word list.
You may assume beginWord and endWord are non-empty and are not the same.
Example 1:

Input:
beginWord = "hit",
endWord = "cog",
wordList = ["hot","dot","dog","lot","log","cog"]

Output: 5

Explanation: As one shortest transformation is "hit" -> "hot" -> "dot" -> "dog" -> "cog",
return its length 5.
Example 2:

Input:
beginWord = "hit"
endWord = "cog"
wordList = ["hot","dot","dog","lot","log"]

Output: 0

Explanation: The endWord "cog" is not in wordList, therefore no possible transformation.

=================================
Should work, but too slow to pass all test cases.
https://leetcode.com/submissions/detail/232588381/
*/

package problem127

import "fmt"

func ladderLength(beginWord string, endWord string, wordList []string) int {
	contain := 0
	bwidx := -1
	for i:=0; i<len(wordList); i++ {
		if wordList[i] == endWord {
			contain++
		}
		if wordList[i] == beginWord {
			bwidx = i
		}
	}
	if contain == 0 {
		return 0
	}

	// remove beginWord from wordList, if there exists one
	if bwidx >= 0 {
		wordList = append(wordList[:bwidx], wordList[bwidx+1:]...)
	}
	//fmt.Println(transformable("ait", "aat"))
	ret := helper(beginWord, endWord, wordList, 0, len(wordList))
	if ret < 0 {
		return 0
	}
	return ret+1
}

func helper(beginWord string, endWord string, wordList []string, count, maxLength int) int {
	fmt.Println(beginWord, wordList, count)
	if transformable(beginWord, endWord) {
		return count+1
	}

	shortestPath := maxLength+1
	for i:=0; i<len(wordList); i++ {
		word := wordList[i]
		if transformable(word, beginWord) {
			tmpList := make([]string, len(wordList))
			copy(tmpList, wordList)
			tmpList = append(tmpList[:i], tmpList[i+1:]...)  // review https://play.golang.org/p/i-4FlM1VEMK.go
			//fmt.Println("tlist:", tmpList)
			if tmp := helper(word, endWord, tmpList, count+1, maxLength); tmp > 0 && tmp < shortestPath {
				shortestPath =  tmp
			}
		}
	}

	if shortestPath == maxLength+1 {
		return -maxLength-1
	}

	return shortestPath
}

// assume str1 and str2 have same length
func transformable(str1, str2 string) bool {
	//fmt.Println(str1, str2)
	flag := false
	for i:=0; i<len(str1); i++ {
		d := str1[i] - str2[i]
		if !flag {
			if d > 0 {
				flag = true
			}
		} else {
			if d != 0 {
				return false
			}
		}
	}
	return flag  // if str1 == str2, return false
}
