# **Core**

[![Home](https://godoc.org/github.com/gookit/event?status.svg)](file:///D:/EC-TSJ/Documents/CODE/SOURCE/Go/pkg/lib/cli)
[![Build Status](https://travis-ci.org/gookit/event.svg?branch=master)](https://travis-ci.org/)
[![Coverage Status](https://coveralls.io/repos/github/gookit/event/badge.svg?branch=master)](https://coveralls.io/github/)
[![Go Report Card](https://goreportcard.com/badge/github.com/gookit/event)](https://goreportcard.com/report/github.com/)

> **[EN README](README.md)**

Trie es una librería para realizar la estructura trie.

## GoDoc

- [godoc for github](https://godoc.org/github.com/)

## Funciones Principales
---
--- 

Tiene los tipos siguientes:


Funciones: 

 - *`New() *trie`* 

objeto ***trie***, con métodos:
 - *`Size() int`*
 - *`Add(string, interface{}) `* 
 - *`Get(string) string `*
 - *`Set(string, string) `*
 - *`Find(string) bool`*
 - *`Remove(string)`*
 - *`Prefix(string) bool`*
 - *`Keys(string) []string`*




## Ejemplos
```go

	trie := New()
	words := []string{"sam", "john", "tim", "jose", "rose",
		"cat", "dog", "dogg", "roses"}
	for _, a := range words {
		trie.Add(a)
	}
	wordsToFind := []string{"sam", "john", "tim", "jose", "rose",
		"cat", "dog", "dogg", "roses", "rosess", "ans", "san"}
	trie.Remove("john")
	f := trie.Keys("")
	for _, a := range wordsToFind {
		found := trie.Find(a)
		if found {
			fmt.Printf("Word \"%s\" found in trie\n", a)
		} else {
			fmt.Printf("Word \"%s\" not found in trie\n", a)
		}
	}


```
## Notas





<!-- - [gookit/ini](https://github.com/gookit/ini) INI配置读取管理，支持多文件加载，数据覆盖合并, 解析ENV变量, 解析变量引用
-->
## LICENSE

**[MIT](LICENSE)**
