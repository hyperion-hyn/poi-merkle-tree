package merkletree

import (
	"fmt"
	"github.com/cbergoon/merkletree"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestMakeTree(t *testing.T)  {
	var list []merkletree.Content
	list = append(list, Content{Attr: "Hello"})
	list = append(list, Content{Attr: "Hi"})
	list = append(list, Content{Attr: "Hey"})
	list = append(list, Content{Attr: "Hola"})

	addr := make(map[string]string, 0)
	addr["zh-Hans"] = "Hyperion"
	list = append(list, Content{Attr: addr})

	//Create a new Merkle Tree from the list of Content
	tree, err := merkletree.NewTree(list)
	if err != nil {
		assert.Error(t, err)
	}

	type Test struct {
		RootHash	string
		VT			bool
		VC			bool
		NVC			bool
	}
	test := Test{}

	//Get the Merkle Root of the tree
	mr := tree.MerkleRoot()
	test.RootHash = fmt.Sprintf("%x", mr)
	log.Println(mr)
	log.Printf("%x", mr)

	//Verify the entire tree (hashes for each node) is valid
	vt, err := tree.VerifyTree()
	vtExpiression := true
	if err != nil {
		assert.Error(t, err)
	}
	test.VT = vt
	log.Println("Verify Tree: ", vt)

	//Verify a specific content in in the tree
	vc, err := tree.VerifyContent(list[1])
	vcExpiression := true
	test.VC = vc
	if err != nil {
		assert.Error(t, err)
	}
	log.Println("Verify Content: ", vc)

	//Verify a specific content not in in the tree
	nvc, err := tree.VerifyContent(Content{Attr: "Hyperion"})
	nvcExpiression := false
	test.NVC = nvc
	if err != nil {
		assert.Error(t, err)
	}
	log.Println("Verify not in Content: ", nvc)


	assert.Equal(t, "9684bcab759f9a2c93a4293f2aa052cee4e979e48a609a8d7951e50a8e45f358", test.RootHash)
	assert.Equal(t, nil, err)
	assert.Equal(t, vtExpiression, test.VT)
	assert.Equal(t, vcExpiression, test.VC)
	assert.Equal(t, nvcExpiression, test.NVC)
}