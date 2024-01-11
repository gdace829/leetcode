package main

type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	node := this
	for i := len(word) - 1; i >= 0; i-- {
		num := word[i] - 'a'
		if node.children[num] != nil {
			node = node.children[num]
		} else {
			node.children[num] = &Trie{}
			node = node.children[num]
		}
	}
	node.isEnd = true
}
func (this *Trie) SearchPre(word string) *Trie {
	node := this
	for _, v := range word {
		num := v - 'a'
		if node.children[num] != nil {
			node = node.children[num]
		} else {
			return nil
		}
	}
	return node
}
func (this *Trie) Search(word string) bool {
	return this.SearchPre(word) != nil && this.SearchPre(word).isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	return this.SearchPre(prefix) != nil
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
