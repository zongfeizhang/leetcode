package array

func findSubstring1(s string, words []string) []int {
	res := make([]int,0)
	if len(words)==0||len(s)<len(words[0]){
		return res
	}
	subStrLen := len(words)*len(words[0])
	left := 0
	for left<=len(s)-subStrLen{
		if match1(s,left,words){
			res = append(res,left)
		}
		left++
	}
	return res
}

func match1(s string,left int,words []string)bool{
	m := make(map[string]int)
	num := 0
	for start:=left;num<len(words);start+=len(words[0]){
		m[s[start:start+len(words[0])]]++
		num++
	}

	for _,word := range words{
		m[word]--
		if m[word]<0{
			return false
		}
	}
	return true
}



