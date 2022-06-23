# 简易版区块链实现
# 使用
```
package main
import (
	"fmt"
	"strconv"
	"testing"
	"time"
)
    
func main() {
    blockchain := BlockChain{}
	for i := 0; i < 5; i++ {
		b := blockchain.uploadChain("上链数据:" + strconv.Itoa(i))
		if !b {
			fmt.Println("验证区块错误")
		}
		time.Sleep(time.Second * 1)
	}
	// 打印链数据
	blockchain.getChainData()
}
```
打印出的数据：
```go
=== RUN   TestBlockChain
hash: 266cd15c59789786be91873464938cb955ed24a1b5133742fccee6e964ae1bf2
blockNum: 0
timestamp: 1655955949
preHash: 
data: genesis block data
-----------------分割线------------------
hash: 9d9f8fe10611eaa153f22c2874fe97eef57fe6d581a94c2821e5f53ef252abba
blockNum: 1
timestamp: 1655955950
preHash: 266cd15c59789786be91873464938cb955ed24a1b5133742fccee6e964ae1bf2
data: 上链数据:1
-----------------分割线------------------
hash: fd6355fd343ed65300726002b7a17bc1edaeb904adff9661106158625f09d4df
blockNum: 2
timestamp: 1655955951
preHash: 9d9f8fe10611eaa153f22c2874fe97eef57fe6d581a94c2821e5f53ef252abba
data: 上链数据:2
-----------------分割线------------------
hash: a46b1e05464c191b7dab8d2d00048a926562f474b2bf6fe8ac313e9220c65165
blockNum: 3
timestamp: 1655955952
preHash: fd6355fd343ed65300726002b7a17bc1edaeb904adff9661106158625f09d4df
data: 上链数据:3
-----------------分割线------------------
hash: e782e286cc58e49525592e86c89fab3fbb0ccb160f8c674dc7a25070f77a2370
blockNum: 4
timestamp: 1655955953
preHash: a46b1e05464c191b7dab8d2d00048a926562f474b2bf6fe8ac313e9220c65165
data: 上链数据:4
-----------------分割线------------------
最后附完整代码
```
