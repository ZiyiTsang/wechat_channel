package BST

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"
)

type Node struct {
	index   int
	payload interface{}
	left    *Node
	right   *Node
}

type Tree struct {
	count int
	root  *Node
}

func NewTree() *Tree {
	return &Tree{
		count: 0,
		root:  nil,
	}
}

//i=index,p=payload
func NewNode(i int, p interface{}) *Node {
	return &Node{index: i, payload: p}
}
func (BST *Tree) Insert(i int, p interface{}) {
	if BST.root == nil {
		//println("Add root")
		node := NewNode(i, p)
		BST.root = node
		BST.count++
	} else {
		currentNode := BST.root
		for {
			if i < currentNode.index {
				if currentNode.left == nil {
					newNode := NewNode(i, p)
					currentNode.left = newNode
					//fmt.Printf("find place! locate at %d's left\n", currentNode.index)
					BST.count++
					break
				} else {
					currentNode = currentNode.left
				}
			} else {
				if currentNode.right == nil {
					newNode := NewNode(i, p)
					currentNode.right = newNode
					BST.count++
					//fmt.Printf("find place! locate at %d's right\n", currentNode.index)
					break
				} else {
					currentNode = currentNode.right
				}
			}
		}
	}
	//println("insert " + strconv.Itoa(i) + " successful")

}

func (BST *Tree) Find(i int) (c *Node, p *Node, err error) {
	if BST.root == nil {
		return nil, nil, errors.New("tree is empty")
	}
	var parentNode *Node
	currentNode := BST.root
	for {
		if currentNode.index == i {
			break
		}
		if i < currentNode.index {
			if currentNode.left != nil {
				parentNode = currentNode
				currentNode = currentNode.left
				continue
			} else {
				return nil, nil, errors.New("node is node exist")
			}
		}
		if i > currentNode.index {
			if currentNode.right != nil {
				parentNode = currentNode
				currentNode = currentNode.right
				continue
			} else {
				return nil, nil, errors.New("node is node exist")
			}
		}
	}
	return currentNode, parentNode, nil
}
func (BST *Tree) Delete(i int) error {
	currentNode, parentNode, err := BST.Find(i)
	if err != nil {
		return errors.New("Delete fail,because " + err.Error())
	}
	var leftOrRight bool
	if parentNode.left == currentNode {
		leftOrRight = true
	} else {
		leftOrRight = false
	}
	if currentNode.right == nil && currentNode.left == nil {
		switch leftOrRight {
		case true:
			parentNode.left = nil
		case false:
			parentNode.right = nil
		}
	} else if currentNode.right != nil && currentNode.left != nil {
		var (
			minNode, minNodeParent *Node
		)
		minNodeParent = currentNode
		minNode = currentNode.right
		for minNode.left != nil {
			minNodeParent = minNode
			minNode = minNode.left
		}
		minNodeParent.left = nil
		currentNode.index = minNode.index
		currentNode.payload = minNode.payload
	} else if currentNode.right != nil || currentNode.left != nil {
		if currentNode.right != nil {
			currentNode = currentNode.right
		} else {
			currentNode = currentNode.left
		}
	}
	return nil
}
func (BST *Tree) InOrderTraverseNormal() {
	InOrderTraverseNormal(BST.root)
	println(" ")
}
func (BST *Tree) InOrderTraverseWithTime(ctx context.Context) error {
	select {
	case <-InOrderTraverse(BST.root):
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
func InOrderTraverse(root *Node) chan interface{} {
	ch := make(chan interface{}, 0)
	go InOrderTraverseIterate(root, 1, ch)
	return ch
}

func InOrderTraverseIterate(node *Node, level int, ch chan interface{}) {
	if node.left != nil {
		InOrderTraverseIterate(node.left, level+1, ch)
	}
	print(node.index, " ")
	if node.right != nil {
		InOrderTraverseIterate(node.right, level+1, ch)
	}
	if level == 1 {
		println(" ")
		ch <- "done"
	}
	return
}
func InOrderTraverseNormal(node *Node) {
	if node.left != nil {
		InOrderTraverseNormal(node.left)
	}
	print(node.index, " ")
	if node.right != nil {
		InOrderTraverseNormal(node.right)
	}
	return
}
func TestInsert(t *testing.T) {
	tree := NewTree()
	tree.Insert(5, "")
	tree.Insert(4, "")
	tree.Insert(9, "")
	tree.Insert(100, "")
}
func TestPostOrder(t *testing.T) {
	tree := NewTree()
	tree.Insert(5, "")
	tree.Insert(4, "")
	tree.Insert(9, "")
	tree.Insert(100, "")
	tree.Insert(1, "")
	cfx, _ := context.WithTimeout(context.Background(), time.Second)
	err := tree.InOrderTraverseWithTime(cfx)
	if err != nil {
		fmt.Println(err)
	}
}
func TestFind(t *testing.T) {
	tree := NewTree()
	tree.Insert(5, "")
	tree.Insert(4, "")
	tree.Insert(9, "")
	tree.Insert(100, "")
	tree.Insert(1, "")
	_, _, err := tree.Find(100)
	if err != nil {
		return
	}
}
func TestDelete(t *testing.T) {
	tree := NewTree()
	tree.Insert(5, "")
	tree.Insert(4, "")
	tree.Insert(9, "")
	tree.Insert(100, "")
	tree.Insert(1, "")
	println("Before:")
	tree.InOrderTraverseNormal()
	err := tree.Delete(199)
	if err != nil {
		fmt.Println(err)
	}
	println("After")
	tree.InOrderTraverseNormal()

}
