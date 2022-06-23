
import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

// 最简单的区块链实现

// 块数据
type Block struct {
	Hash        string // 当前块的hash
	BlockNumber int64  // 当前块的高度
	PreHash     string // 上个区块的hash
	Timestamp   int64  // 生成当前区块的高度
	Data        string // 区块交易(上链数据)
}

// 基础方法
// 创建创世区块
func genesisBlock() Block {
	b := Block{}
	b.BlockNumber = 0
	b.PreHash = ""
	b.Timestamp = time.Now().Unix()
	b.Data = "genesis block data"
	b.Hash = generateHash(b)
	return b
}

// 创建区块
func generateBlock(oldBlock Block, data string) Block {
	var b Block
	b.BlockNumber = oldBlock.BlockNumber + 1
	b.PreHash = oldBlock.Hash
	b.Timestamp = time.Now().Unix()
	b.Data = data
	b.Hash = generateHash(b)
	return b
}

// 生成hash
func generateHash(b Block) string {
	str := fmt.Sprintf("%d,%d,%s,%s", b.BlockNumber, b.Timestamp, b.PreHash, b.Data)
	hash := sha256.New()
	hash.Write([]byte(str))
	bytes := hash.Sum(nil)
	return hex.EncodeToString(bytes)
}

// 区块"链"
type BlockChain struct {
	ChainData []Block
}

// 数据区块数据是否合法
func (bc *BlockChain) verifyBlock(newBlock Block, oldBlock Block) bool {
	// 对比前后高度是否一致
	if newBlock.BlockNumber != oldBlock.BlockNumber+1 {
		return false
	}
	// 对比前后hash是否一致
	if newBlock.PreHash != oldBlock.Hash {
		return false
	}

	// 对比生成的hash是否一致
	if newBlock.Hash != generateHash(newBlock) {
		return false
	}
	return true
}

// 输出区块数据记录
func (bc *BlockChain) getChainData() {
	for _, v := range bc.ChainData {
		fmt.Println("hash:", v.Hash)
		fmt.Println("blockNum:", v.BlockNumber)
		fmt.Println("timestamp:", v.Timestamp)
		fmt.Println("preHash:", v.PreHash)
		fmt.Println("data:", v.Data)
		fmt.Println("-----------------分割线------------------")
	}
}

// 数据上链
func (bc *BlockChain) uploadChain(data string) bool {
	num := len(bc.ChainData)
	if num == 0 { // 还未产生块
		block := genesisBlock()
		bc.ChainData = append(bc.ChainData, block)
	} else {
		oldBlock := bc.ChainData[len(bc.ChainData)-1]
		var newBlock = generateBlock(oldBlock, data)
		if bc.verifyBlock(newBlock, oldBlock) {
			bc.ChainData = append(bc.ChainData, newBlock)
		} else {
			return false
		}
	}
	return true
}
