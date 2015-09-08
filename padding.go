package padding

type BlockPadding interface {
	Pad(data []byte, blockSize int) []byte
	Unpad(data []byte) []byte
}
