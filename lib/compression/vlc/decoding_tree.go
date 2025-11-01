package vlc

import "strings"

//any tree will be defined by its root node => in this type, not the entire tree will be defined, but only its node
//Since each node will have links to child elements, if any, we will be able to access all tree data along the chain
type DecodingTree struct {
	Value string
	Zero  *DecodingTree
	One   *DecodingTree
}

//a function that will build a tree corresponding to it from the encoding table
func (et encodingTable) DecodingTree() DecodingTree {
	res := DecodingTree{}
	for ch, code := range et {
		res.Add(code, ch)
	}
	return res
}

//we get the character code 010000 as input , we need to sort through all these characters one by one, building a tree based on them
//and inside the node that corresponds to the last character, we put the value value 01000(0)<-value
func (dt *DecodingTree) Add(code string, value rune) {
	currentNode := dt
	for _, ch := range code {
		switch ch {
		case '0':
			if currentNode.Zero == nil {
				currentNode.Zero = &DecodingTree{}
			}
			currentNode = currentNode.Zero
		case '1':
			if currentNode.One == nil {
				currentNode.One = &DecodingTree{}
			}
			currentNode = currentNode.One
		}
	}
	currentNode.Value = string(value)
}

//a method for a tree that deals with decoding
func (dt *DecodingTree) Decode(str string) string {
	var buf strings.Builder
	currentNode := dt
	for _, ch := range str {
		if currentNode.Value != "" {
			buf.WriteString(currentNode.Value)
			currentNode = dt
		}
		switch ch {
		case '0':
			currentNode = currentNode.Zero
		case '1':
			currentNode = currentNode.One
		}
	}
	if currentNode.Value != "" {
		buf.WriteString(currentNode.Value)
		currentNode = dt
	}
	return buf.String()
}
