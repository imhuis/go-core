package functions

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
)

type Node struct {
	Type                    html.NodeType
	Data                    string
	Attr                    []html.Attribute
	FirstChild, NextSibling *Node
}

type NodeType int32

const (
	ErrorNode NodeType = iota
	TextNode
	DocumentNode
	ElementNode
	CommentNode
	DoctypeNode
)

type Attribute struct {
	key, val string
}

//func Parse(r io.Reader) (*html.Node, error)

func htmlParse(in io.Reader) {
	doc, err := html.Parse(in)
	if err != nil {
		panic(err)
	}
	fmt.Print(doc)
}
