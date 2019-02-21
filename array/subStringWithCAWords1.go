package array

func FindSubstring(s string, words []string) []int {
	res := make([]int,0)
	if len(words)==0{
		return res
	}
	wordsMap := buildWordsMap(words)
	countPer := len(words[0])
	strsCount := len(words) * countPer
	for i:=0;i<len(words[0]);i++{
		left := i
		right := i + countPer //
		m := make(map[string]int)
		for right<=len(s){
			curWord := s[right-countPer:right]
			if wordsMap[curWord]==0{
				left = right
				right += countPer
				m = make(map[string]int)
				continue
			}
			m[curWord]++
			if right-left<strsCount{
				right += countPer
				continue
			}
			if right-left==strsCount{
				if match(m,wordsMap){
					res = append(res,left)
				}
				m[s[left:left+countPer]]--
				left += countPer
				right += countPer
			}
		}
	}
	return res

}

func match(m,words map[string]int)bool{
	for key := range words{
		if m[key]!=words[key]{
			return false
		}
	}
	return true
}

func buildWordsMap(words[]string)map[string]int{
	m := make(map[string]int)
	for _,word := range words{
		m[word]++
	}
	return m
}
