package vlc

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

type encodingTable map[rune]string
type BinaryChunk string
type BinaryChunks []BinaryChunk

const chunksSize = 8

// splitting the received binary code from encodeBin into fragments ('10010111001010110011011' ->'1001011 10010101 10011011')
// the size of the chunk is passed using argument
// and the custom type will be used.
func splitByChunks(binStr string, chunkSize int) BinaryChunks {
	strLen := utf8.RuneCountInString(binStr)
	chunksCount := strLen / chunkSize
	if strLen/chunkSize != 0 {
		chunksCount++
	}
	res := make(BinaryChunks, 0, chunksCount)
	var buf strings.Builder
	for i, ch := range binStr {
		buf.WriteString(string(ch))
		if (i+1)%chunkSize == 0 {
			res = append(res, BinaryChunk(buf.String()))
			buf.Reset()
		}
	}
	if buf.Len() != 0 {
		lastChunk := buf.String()
		lastChunk += strings.Repeat("0", chunkSize-len(lastChunk))
		res = append(res, BinaryChunk(lastChunk))
	}
	return res
}

// the function that will extract the 16th chunks from the received string
func NewBinChunks(data []byte) BinaryChunks {
	res := make(BinaryChunks, 0, len(data))
	for _, code := range data {
		res = append(res, NewBinChunk(code))
	}
	return res
}

func NewBinChunk(code byte) BinaryChunk {
	return BinaryChunk(fmt.Sprintf("%08b", code))
}

// we iterate through the chunks and convert them to bytes
func (bcs BinaryChunks) Bites() []byte {
	res := make([]byte, 0, len(bcs))
	for _, bc := range bcs {
		res = append(res, bc.Byte())
	}
	return res
}

func (bc BinaryChunk) Byte() byte {
	num, err := strconv.ParseUint(string(bc), 2, chunksSize)
	if err != nil {
		panic("can t parse binary chunk:" + err.Error())
	}
	return byte(num)
}

// combining chunks into one line
func (bcs BinaryChunks) Join() string {
	var buf strings.Builder
	for _, bc := range bcs {
		buf.WriteString(string(bc))
	}
	return buf.String()
}
