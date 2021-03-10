package rule_engine

const LEFT = "left"
const RIGHT = "right"

const TypeNode = "node"   // 节点
const TypeParam = "param" // 参数
const TypeValue = "value" // 值
const TypeFunc = "func"   // 函数

// 节点
type Node struct {
	Type        string
	Logic       string

	Left        *Node
	LeftResult  interface{}

	Right       *Node
	RightResult interface{}

	Result      interface{}

	Name   string
	Params []interface{}
	Value  interface{}
}
