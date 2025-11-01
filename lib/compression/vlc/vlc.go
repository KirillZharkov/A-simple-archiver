package vlc

import (
	"strings"
	"unicode"
)

// it is used as one of the implementations of the Encoder and Decoder interfaces
// that is, it will compress and decompress files using the variable length code method
type EncoderDecoder struct{}

func New() EncoderDecoder {
	return EncoderDecoder{}
}

func (vlc EncoderDecoder) Encode(str string) []byte {
	////prepare text: M = !m
	str = prepareText(str)
	//let's turn the string into a binary sequence, that is, into an array of bits 0 and 1, such as 10010101
	binStr := encodeBin(str)
	//splitting a binary string into chunks, the size of the chunks is 8 characters,
	//that is, we split the array of characters into bytes and get the type of this - '1001011 10010101 10011011'
	chunks := splitByChunks(binStr, chunksSize)
	// then we represent the received bytes as 16 numbers and get a type like this - '20 30 3C' - a set of 16 numbers
	//chunks.toHex()
	//let's return the resulting sequence of 16 numbers as a string
	return chunks.Bites()
}

// prepares the text for encoding: M = !m
func prepareText(str string) string {
	var buf strings.Builder

	for _, ch := range str {
		if unicode.IsUpper(ch) {
			buf.WriteRune('!')
			buf.WriteRune(unicode.ToLower(ch))
		} else {
			buf.WriteRune(ch)
		}
	}
	return buf.String()
}

// string to string of binary codes (10010101)
func encodeBin(str string) string {
	var buf strings.Builder
	for _, i := range str {
		buf.WriteString(bin(i))
	}
	return buf.String()
}

// which returns the matched characters in binary sequences
func bin(ch rune) string {
	table := getEncodingTable()
	res, ok := table[ch]
	if !ok {
		panic("unknown character:" + string(ch))
	}
	return res
}

// bin will require an encoding operation function (table)
func getEncodingTable() encodingTable {
	return encodingTable{
		' ': "11",
		'e': "101",
		't': "1001",
		'o': "10001",
		'n': "10000",
		'a': "011",
		's': "0101",
		'i': "01001",
		'r': "01000",
		'h': "0011",
		'd': "00101",
		'l': "001001",
		'!': "001000",
		'u': "00011",
		'c': "000101",
		'f': "000100",
		'm': "000011",
		'p': "0000101",
		'g': "0000100",
		'w': "0000011",
		'b': "0000010",
		'y': "0000001",
		'v': "00000001",
		'j': "000000001",
		'k': "0000000001",
		'x': "00000000001",
		'q': "000000000001",
		'z': "000000000000",
	}
}

func (vlc EncoderDecoder) Decode(encodedData []byte) string {
	hChunks := NewBinChunks(encodedData)
	bString := hChunks.Join()
	dTree := getEncodingTable().DecodingTree()
	return exportText(dTree.Decode(bString)) //вернет !my name is !ted
}

/*
exportText is the opposite of prepareText, it prepares the decoded text for export:
it
modifies: ! + <lowercase letter> -> to uppercase letter.
for example: !My name is!ted -> My name is Ted.
*/
func exportText(str string) string {
	var buf strings.Builder
	var isCapital bool
	for _, ch := range str {
		if isCapital {
			buf.WriteRune(unicode.ToUpper(ch))
			isCapital = false
			continue
		}
		if ch == '!' {
			isCapital = true
			continue
		} else {
			buf.WriteRune(ch)
		}

	}
	return buf.String()
}
