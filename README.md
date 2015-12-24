# iptrie
gen_data/gen_data.go  生成测试数据

    WORKDIR     string  = "../iptrie/"  生成数据的目录
    INPUTSIZE   int     = 2000000       生成数据的个数
    TESTSIZE    int     = 10000         测试数据的个数
    INPUTFILE   string  = WORKDIR + "input.txt"  与 iptrie_test.go 中 Load 的文件相同
    TESTFILE    string  = WORKDIR + "test.txt"   与 iptrie_test.go 中 Test 的文件相同

iptrie/iptrie.go
    type TrieNode struct {
        Entry       interface{}
        Children    []*TrieNode
    }

    type IPTrie struct {
        root *TrieNode
    }
    
    type IPEntry srtuct {
        cwnd int
    }

    func (t *IPTrie) Set(path string, length int, entry IPEntry) error {}
        path is "a.b.c.d", use IPToBinary(path) -> binary form
        length is prefix length
        entry is the value stored at leaf node

    func (t *IPTrie) Get(path string) (entry interface{}, ok bool) {}
    if ok == true { entry found }

    func (t *IPTrie) LoadFromFile(filename string) error {}
        read "a.b.c.d, length, cwnd" from file named filename
        and call t.Set(path, length, IPEntry(cwnd))

