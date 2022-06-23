
import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

// 生成区块数据
func TestBlockChain(t *testing.T) {
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
