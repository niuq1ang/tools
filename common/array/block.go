package array

type Block struct {
	Since int
	End   int
}

var DefaultBlockSize = 10

func NewBlocks(length int, blockSize int) []*Block {
	blocks := []*Block{}
	if blockSize <= 0 {
		return blocks
	}
	since := 0
	for i := 0; i < length; {
		i += blockSize
		if i > length {
			i = length
		}
		blocks = append(blocks, &Block{Since: since, End: i})
		since = i
	}
	return blocks
}

func (b *Block) Range() int {
	return b.End - b.Since
}

func NewDefalutBlocks(length int) []*Block {
	return NewBlocks(length, DefaultBlockSize)
}
