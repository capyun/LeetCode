package Trie

/**
@author Shitanford
@create 2021-10-04 16:21
*/

func main() {

}

type WordDictionary struct {
	trie Trie
}


func Constructor() WordDictionary {
	return WordDictionary{
		trie: Trie{
			charSet: make(map[rune]*Trie),
		},
	}
}


func (this *WordDictionary) AddWord(word string)  {
	this.trie.Insert(word)
}


func (this *WordDictionary) Search(word string) bool {
	return this.trie.Search(word, '.')
}

