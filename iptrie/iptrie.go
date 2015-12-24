package iptrie

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type TrieNode struct {
	Entry    	interface{}
	Children 	[]*TrieNode
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		Children: make([]*TrieNode, 2),
	}
}

type IPTrie struct {
	root *TrieNode
}

func NewIPTrie() *TrieNode {
	root := NewTrieNode()
	return root
}

type IPEntry struct {
	cwnd int
}

// func (e IPEntry) Entry() {
// 	fmt.Printf("%v", e.cwnd)
// }

func (e IPEntry) String() string {
	return fmt.Sprintf("cwnd: %v", e.cwnd)
}

func IPToBinary(ip string) (string, error) { // "a.b.c.d" to "1111100000..."
	addr := strings.Split(ip, ".")
	sum := 0
	for i := 0; i < len(addr); i++ {
		sum = sum * 256
		a, err := strconv.Atoi(addr[i])
		if err != nil {
			return "", err
		}
		sum += a
	}
	var res = make([]string, 32)
	for i := 31; i >= 0; i-- {
		res[i] = string(sum&1 + '0')
		sum = sum >> 1
	}
	return strings.Join(res, ""), nil
}


func (t *IPTrie) Set(path string, length int, entry IPEntry) error {
	node := t.root

	binary, err := IPToBinary(path)
	if err != nil {
		return err
	}

	for i := 0; i < length; i++ {
		if v := node.Children[binary[i]-'0']; v != nil {
			node = v
		} else {
			node.Children[binary[i]-'0'] = NewTrieNode()
			node = node.Children[binary[i]-'0']
		}
	}
	node.SetEntry(entry)
	return nil
}

func (t *IPTrie) LoadFromFile(name string) error {
	file, err := os.Open(name)
	if err != nil {
		return err
	}
	defer file.Close()

	for {
		var (
			str    string = "0.0.0.0"
			length int
			entry  IPEntry
		)
		argc, err := fmt.Fscanf(file, "%s%d,%d", &str, &length, &entry.cwnd)
		if (argc > 0 && argc != 3) {
			fmt.Println(str, "Format not right")
		} else {
			if err == io.EOF {
				return nil
			}
			if argc == 0 {
				return nil
			}
		}

		str = str[:len(str)-1] 		// remove comma
		t.Set(str, length, entry)
	}
}

func (t *IPTrie) Get(path string) (entry interface{}, ok bool) {
	node := t.root
	binary, err := IPToBinary(path)

	if err != nil {
		return node.Entry, false
	}
	entry = node.Entry

	for i := 0; i < 32; i++ {
		if node.Children[binary[i]-'0'] != nil {
			node = node.Children[binary[i]-'0']
		} else {
			entry, ok = node.GetEntry()
			if !ok {
				return nil, false
			} else {	// ok == true
				break
			}
		}
	}
	return entry, true
}

func (t *TrieNode) SetEntry(value interface{}) {
	t.Entry = value
}

func (t *TrieNode) GetEntry() (entry interface{}, ok bool) {
	return t.Entry, t.Entry != nil
}
