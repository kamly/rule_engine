package dsl_engine

import (
	"fmt"

	"github.com/Knetic/govaluate"
)

func (dsl *Dsl) getNextNode(node *Node) bool {
	// 先计算 left
	if _, ok := node.Left.(*Node); ok {
		node.LeftResult = dsl.getNextNode(node.Left.(*Node))
	} else {
		dsl.getNodeValue(LEFT, node)
	}

	if _, ok := node.Right.(*Node); ok {
		node.RightResult = dsl.getNextNode(node.Right.(*Node))
	} else {
		dsl.getNodeValue(RIGHT, node)
	}

	// 进行逻辑判断 > < <= >= == || &&
	return dsl.calculate(node)
}

// 进行逻辑判断 > < <= >= == || &&
func (dsl *Dsl) calculate(node *Node) bool {
	switch node.Logic {
	case ">", "<", ">=", "<=", "==":
		// 执行逻辑
		var params = make(map[string]interface{})
		params["left"] = node.LeftResult
		params["right"] = node.RightResult

		expr, _ := govaluate.NewEvaluableExpression(fmt.Sprintf("left %s right", node.Logic))
		eval, _ := expr.Evaluate(params)

		return eval.(bool)

	case "&&", "||", "bool":
		var exprStr string
		exprStr += fmt.Sprintf(" %t", node.LeftResult)

		if node.Logic != "bool" {
			exprStr += fmt.Sprintf(" %s", node.Logic)
			exprStr += fmt.Sprintf(" %t", node.RightResult)
		}

		expr, _ := govaluate.NewEvaluableExpression(exprStr)
		eval, _ := expr.Evaluate(nil)

		return eval.(bool)
	}

	return false
}

// 获取每个节点的值
func (dsl *Dsl) getNodeValue(typeNode string, node *Node) {

	newNode := node.Left
	if typeNode == RIGHT {
		newNode = node.Right
	}

	var value interface{}

	switch newNode.(type) {
	case NodeValue: // 具体值
		nodeValue := newNode.(NodeValue)
		value = nodeValue.Value
		break
	case NodeParam: // 参数
		nodeParam := newNode.(NodeParam)
		nodeParam.Value = dsl.getParamListByName(nodeParam.Name)
		value = nodeParam.Value
		break
	case NodeFunc: // 函数
		nodeFunc := newNode.(NodeFunc)
		switch len(nodeFunc.Params) { // TODO 参数个数
		case 1:
			results, _ := callFunc(dsl.FuncList, nodeFunc.Name, nodeFunc.Params[0])
			// TODO 抽离返回值类型
			for _, item := range results {
				switch item.Type().String() {
				case "uint64":
					nodeFunc.Value = item.Uint()
				case "bool":
					nodeFunc.Value = item.Bool()
				}
			}
			value = nodeFunc.Value
		}
		break
	}

	if typeNode == LEFT {
		node.LeftResult = value
	} else {
		node.RightResult = value
	}
}
