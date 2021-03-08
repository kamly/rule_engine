package risk_dsl

const LEFT = "left"
const RIGHT = "right"

// 节点
type Node struct {
	Logic       string
	Left        interface{}
	LeftResult  interface{}
	Right       interface{}
	RightResult interface{}
	Result      interface{}
}

// 参数
type NodeParam struct {
	Name  string
	Value interface{}
}

// 函数
type NodeFunc struct {
	Name   string
	Params []interface{}
	Value  interface{}
}

// 值
type NodeValue struct {
	Value interface{}
}
