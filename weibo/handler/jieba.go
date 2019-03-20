package handler

import (
	"fmt"
	"github.com/yanyiwu/gojieba"
	"regexp"
	"sort"
	"strings"
)

//段落分词
func jiebaFreq(texts string, num int) (freq []string) {
	text := removeKey(texts)
	use_hmm := true
	x := gojieba.NewJieba()
	defer x.Free()
	//使用精确模式
	words := x.CutForSearch(text, use_hmm)
	strings.Join(words, "/")
	freq = append(freq, sliceMax(words, num)...)
	return
}

//词频统计
func sliceMax(text []string, num int) (freq []string) {
	txtMap := make(map[int]string)
	freqMap := make(map[string]int)
	var intArry []int
	for i, tex := range text {
		txtMap[i] = tex
	}
	for k, v := range txtMap {
		freqMap[v] = 1
		delete(txtMap, k)
		for ki, vs := range txtMap {
			if v == vs {
				freqMap[v] += 1
				delete(txtMap, ki)
			}

		}
	}
	for _, val := range freqMap {
		intArry = append(intArry, val)
	}
	sort.Ints(intArry)
	for _, ar := range intArry[len(intArry)-num:] {
		for k, v := range freqMap {
			if v == ar {
				freq = append(freq, k)
				break
			}
		}
	}
	sliceReverse(freq)
	fmt.Println(freq)
	return
}

//常用词剔除
func removeKey(text string) (txt string) {
	fmt.Println("Remove daily keywords...")
	txt = strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(
		strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(
		strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.
		Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.
		Replace(strings.Replace(strings.Replace(strings.Replace(text, "了", "", -1),
		"都", "", -1), "是", "", -1), "？", "", -1), "#",
		"", -1), " ", "", -1), "一下", "", -1), "！", "",
		-1), ":", "", -1), "]", "", -1), "[", "", -1), "）",
		"", -1), "（", "", -1), "@", "", -1), "回复", "", -1),
		"：", "", -1), "：", "", -1), "。", "", -1), "，", "",
		-1), "她", "", -1), "他", "", -1), "我", "", -1), "你",
		"", -1), "的", "", -1)
	//过滤所有数字和字母
	re, _ := regexp.Compile("([a-zA-Z]|[0-9])")
	txt = re.ReplaceAllString(txt, "")

	return
}

//反转数组
func sliceReverse(sli []string) (s []string) {
	for i, j := 0, len(sli)-1; i < j; i, j = i+1, j-1 {
		sli[i], sli[j] = sli[j], sli[i]
	}
	return
}

// 通过两重循环过滤重复元素
func RemoveRepByLoop(slc []string) []string {
	result := []string{} // 存放结果
	for i := range slc {
		flag := true
		for j := range result {
			if slc[i] == result[j] {
				flag = false // 存在重复元素，标识为false
				break
			}
		}
		if flag { // 标识为false，不添加进结果
			result = append(result, slc[i])
		}
	}
	return result
}

// 通过map主键唯一的特性过滤重复元素
func RemoveRepByMap(slc []string) []string {
	result := []string{}
	tempMap := map[string]byte{} // 存放不重复主键
	for _, e := range slc {
		l := len(tempMap)
		tempMap[e] = 0
		if len(tempMap) != l { // 加入map后，map长度变化，则元素不重复
			result = append(result, e)
		}
	}
	return result
}

// 元素去重
func RemoveRep(slc []string) []string {
	if len(slc) < 1024 {
		// 切片长度小于1024的时候，循环来过滤
		return RemoveRepByLoop(slc)
	} else {
		// 大于的时候，通过map来过滤
		return RemoveRepByMap(slc)
	}
}
