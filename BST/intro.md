# Goalng之二叉排序树
### 什么是二叉排序树
二叉排序树也叫二叉搜索树、二叉查找树，简称 ***BST***（Binary Search/Sort Tree）树，它是一种特殊的二叉树。

一棵二叉排序树需要满足以下三个**特征**：
1. 若它的左子树不空，则左子树上所有结点的值均小于它的根结点的值；
2. 若它的右子树不空，则右子树上所有结点的值均大于它的根结点的值；
3. 它的左、右子树也分别为二叉排序树。
   ![alt text](https://tse1-mm.cn.bing.net/th/id/OIP-C.BIj4F72syuUDfCe3_QIOFwHaER?pid=ImgDet&rs=1)

### BST的优势
二叉搜索树作为一种经典的数据结构，它既有链表的快速插入与删除操作的特点，又有数组快速查找的优势；所以应用十分广泛，例如在文件系统和数据库系统一般会采用这种数据结构进行高效率的排序与检索操作。因此其这也是最经典的数据结构之一。

## 定义二叉树
### 理论基础
首先我们需要定义二叉树及其底层数据结构。
BST本质上是一棵树（tree），对于一棵树，我们有两种**存储形式**：

1. **链表法**：浪费空间，但增删查改效率高
2. **数组法**：节约空间，但增删查改效率低

本篇我们将采用链表法构建BST。
### 构建树
首先构建树的底层结构：**节点**。

节点包含几个部分：

1. 索引是BST排序的标准
2. 负载是该节点所带的数据
3. 左右节点地址指向其下一节点
~~~go
type Node struct {
	index   int
	payload interface{}
	left    *Node
	right   *Node
}
~~~
顺便写个构造函数
~~~go
func NewNode(i int, p interface{}) *Node {
	return &Node{index: i, payload: p}
}
~~~
接下来就可以构建BST啦！

其实BST这个class的属性特别简单，只需要一个**根节点地址**和**节点数**（可选）。
~~~go
type Tree struct {
	count int
	root  *Node
}
~~~
顺便建立一个构造函数
~~~go
func NewTree() *Tree {
	return &Tree{ 
		count: 0,  //0个节点
		root:  nil, //没有根节点嗷
	}
}
~~~
## 定义树的操作
对一个数据结构的操作，无外乎就是“**增删查改**”了！！  
择日不如撞日，咱们这就把代码给他写出来叭！


### 增加节点
我就不多加解释了叭，毕竟还蛮简单的。看不懂的代码里有注释哈。
~~~go
//i:index,c:currentNode,p:parentNode
func (BST *Tree) Insert(i int, p interface{}) {
	if BST.root == nil { //当根节点为空时
		node := NewNode(i, p)
		BST.root = node
		BST.count++
	} else { //当根节点非空时
		currentNode := BST.root
		for {
			if i < currentNode.index { //找到需要加入的位置
				if currentNode.left == nil {
					newNode := NewNode(i, p) //新建一个节点
					currentNode.left = newNode //把新建节点加入当前节点的左侧
					BST.count++ //节点数+1
					break
				} else {
					currentNode = currentNode.left
				}
			} else {
				if currentNode.right == nil {
					newNode := NewNode(i, p) 
					currentNode.right = newNode
					BST.count++
					break
				} else {
					currentNode = currentNode.right
				}
			}
		}
	}
}
~~~
顺便写个测试函数
~~~go
func TestInsert(t *testing.T) {
	tree := NewTree()
	tree.Insert(5, "")
	tree.Insert(4, "")
	tree.Insert(9, "")
	tree.Insert(100, "")
}
~~~
### 查找节点
这个也不解释了哈。
~~~go
//i:index,c:currentNode,p:parentNode
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
~~~
顺便写个测试函数：
~~~go
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
~~~

### 删除节点

二叉排序树的删除相对而言要复杂一些，需要分三种情况来处理：

* **第一种情况**：如果要删除的节点没有子节点，我们只需要直接将父节点中指向要删除节点的指针置为 nil。比如下图中的删除节点 55。
* **第二种情况**：如果要删除的节点只有一个子节点（只有左子节点或者右子节点），我们只需要更新父节点中指向要删除节点的指针，让它指向要删除节点的子节点就可以了。比如下图中的删除节点 13。
* **第三种情况**：如果要删除的节点有两个子节点，这就比较复杂了。  
  我们需要先找到这个节点的右子树中节点值最小的节点，把它替换到要删除的节点上，然后再删除掉这个最小节点。因为最小节点肯定没有左子节点（如果有左子结点，那就不是最小节点了），所以，我们可以应用上面两条规则来删除这个最小节点。比如下图中的删除节点 18。（同理，也可以通过待删除节点的左子树中的最大节点思路来实现）

![删除节点的三种情况](https://laravel.gstatics.cn/storage/uploads/images/gallery/2019-10/scaled-1680-/3e6b26363e49ffa0760b1ddc494f32a271951676d2d03397948ecff119d8f932.jpg)

代码如下：
~~~go
func (BST *Tree) Delete(i int) error {
 //使用前面定义的Find函数，找出待删除节点及其父节点
	currentNode, parentNode, err := BST.Find(i) 
	if err != nil { //如果找不到，将执行此处逻辑
		return errors.New("Delete fail,because " + err.Error())
	}
  //确定子节点和父节点的关系。左节点：0，右节点：1
	var leftOrRight bool  
	if parentNode.left == currentNode {
		leftOrRight = true
	} else {
		leftOrRight = false
	}
  //第一种情况，也是最简单的
	if currentNode.right == nil && currentNode.left == nil {
		switch leftOrRight {
		case true:
			parentNode.left = nil
		case false:
			parentNode.right = nil
		}
	} else if currentNode.right != nil && currentNode.left != nil {//第三种情况
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
	} else if currentNode.right != nil || currentNode.left != nil {//第二种情况
		if currentNode.right != nil {
			currentNode = currentNode.right
		} else {
			currentNode = currentNode.left
		}
	}
	return nil
}
~~~
顺便写个测试函数：
~~~go
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
~~~
### 改动节点

改动节点需要依次执行以下操作：
1. 找到该节点(Find函数)
2. 删除该节点(Delete函数)
3. 加入新节点(Insert函数)

这三个函数都已经给出代码实现了，这块我就不写啦！~~绝对不是因为我懒得写~~

### 中序遍历
#### 普通版

~~~go
func (BST *Tree) InOrderTraverseNormal() {
	InOrderTraverseNormal(BST.root) //调用递归函数
	println(" ") //输出空行
}
func InOrderTraverse(node *Node) {
	if node.left != nil { //左子树进行递归
		InOrderTraverse(node.left)
	}
	print(node.index, " ") //输出自己
	if node.right != nil { //右子树进行递归
		InOrderTraverse(node.right) 
	}
	return
}
~~~

#### PRO版(带超时控制)
这个有点难度，涉及到协程和上下文操作，没法理解就算了。
~~~go
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
~~~

## 总结

二叉排序树作为最经典也是最常用的数据结构之一，其在业务环境里被广泛使用。  
此篇文章也仅是一个*抛砖引玉*的效果，很多概念和代码并未展出。例如前序遍历，后续遍历，平衡版的BST等等。  
需要指出的是，本文的树class中count字段并未被Mutex所保护，因此本篇所写BST并不是**线程安全**的。

###### 尾注
参考资料：
1. Go 数据结构和算法篇（十七）：二叉排序树，https://geekr.dev/posts/go-binary-search-tree
2. 《数据结构与算法之美》，极客时间

Github地址：  
https://github.com/ZiyiTsang/wechat_channel  
[点我跳转](https://github.com/ZiyiTsang/wechat_channel)






