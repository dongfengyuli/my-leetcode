package main

import "fmt"

func main(){

	word := "apple"
	obj := Constructor();
	obj.Insert(word);
	fmt.Println(obj.next)
	param_2 := obj.Search(word);
	param_3 := obj.StartsWith("app");

	fmt.Println("param_2",param_2)
	fmt.Println("param_3",param_3)
}


type Trie struct {
	isEnd bool
	next [26]*Trie
}


/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
	for _,v := range word {
		if this.next[v-'a'] == nil {
			this.next[v-'a'] = new(Trie)
		}
		this = this.next[v-'a']
	}
	this.isEnd = true
}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	for _,v := range word{
		if this.next[v-'a'] == nil {
			return false
		}
		this = this.next[v-'a']
	}
	if this.isEnd == false{
		return false
	}
	return true
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	for _,v := range prefix{
		if this.next[v-'a'] == nil {
			return false
		}
		this = this.next[v-'a']
	}
	return true
}