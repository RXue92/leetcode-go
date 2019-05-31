/*
Runtime: 40 ms, faster than 93.17% of Go online submissions for Word Ladder.
Memory Usage: 8.6 MB, less than 27.20% of Go online submissions for Word Ladder.
https://leetcode.com/submissions/detail/232632731/
BFS on undirected graph based on official Python solution.
 */

package problem127

func BFS(beginWord string, endWord string, wordList []string) int {
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

	combos := make(map[string][]string)
	for _, word := range wordList {
		for i:=0; i<len(word); i++ {
			tmpState := word[:i]+"*"+word[i+1:]
			if _, ok := combos[tmpState]; ok {
				combos[tmpState] = append(combos[tmpState], word)
			} else {
				combos[tmpState] = []string{word}
			}
		}
	}

	visited := map[string]bool {beginWord: true}
	toVisit := []string{beginWord}
	toVisitLevel := []int{1}

	for len(toVisit) > 0 {
		word := toVisit[0]
		level := toVisitLevel[0]
		toVisit = toVisit[1:]
		toVisitLevel = toVisitLevel[1:]

		for i:=0; i<len(word); i++ {
			tmpState := word[:i] + "*" + word[i+1:]

			// breath first search guarantees shorter path returned before longer path
			for _, node := range combos[tmpState] {
				if node == endWord {
					return level+1
				}

				if _, ok := visited[node]; !ok {
					visited[node] = true
					toVisit = append(toVisit, node)
					toVisitLevel = append(toVisitLevel, level+1)
				}
			}
			delete(combos, tmpState)   // avoid checking visited tmp state
		}
	}
	return 0
}