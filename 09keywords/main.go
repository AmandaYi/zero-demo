package main

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/stringx"
)

func main() {
	//关键词替换-具有最长替换特性
	replacer := stringx.NewReplacer(map[string]string{
		"JS":   "Golang",
		"喜欢JS": "非常喜欢Golang",
	})
	println(replacer.Replace("我喜欢JS语言")) // 我非常喜欢Golang
	//查找敏感词
	trie := stringx.NewTrie([]string{
		"坏人", "讨厌",
	})
	trieList := trie.FindKeywords("这有坏人")
	fmt.Println(trieList) // [坏人]
	//敏感词过滤,使用选项模式配置过滤选项
	filterTrie := stringx.NewTrie([]string{
		"色情",
	}, stringx.WithMask('*'))
	sentence, keywords, found := filterTrie.Filter("这里有色情网站!")
	fmt.Printf("sentence = %s, keywords = %v, found = %v", sentence, keywords, found)
}
