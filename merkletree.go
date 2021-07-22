package merkletree

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"github.com/cbergoon/merkletree"
)

type Content struct {
	Attr interface{}
}

func (c Content) CalculateHash() ([]byte, error) {
	h := sha256.New()
	b, err := c.GetBytes()
	if err != nil {
		return nil, err
	}
	if _, err := h.Write(b); err != nil {
		return nil, err
	}

	return h.Sum(nil), nil
}

func (c Content) Equals(other merkletree.Content) (bool, error) {
	return c.Attr == other.(Content).Attr, nil
}

func (c *Content) GetBytes() ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(c.Attr)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func MakeTree(attrs []interface{}) (*merkletree.MerkleTree, error) {
	var list []merkletree.Content
	for _, attr := range attrs {
		list = append(list, Content{Attr: attr})
	}

	tree, err := merkletree.NewTree(list)
	if err != nil {
		return nil, err
	}

	return tree, nil
}