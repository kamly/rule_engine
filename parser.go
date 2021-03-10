package rule_engine

import (
	"fmt"

	"github.com/Knetic/govaluate"
)

func (rule *Rule) getNextNode(node *Node) bool {
	// 先计算 left
	if node.Left != nil {
		if node.Left.Type == TypeNode {
			node.LeftResult = rule.getNextNode(node.Left)
		} else {
			rule.getNodeValue(LEFT, node)
		}
	}

	if node.Right != nil {
		if node.Right.Type == TypeNode {
			node.RightResult = rule.getNextNode(node.Right)
		} else {
			rule.getNodeValue(RIGHT, node)
		}
	}

	// 进行逻辑判断 > < <= >= == || &&
	return rule.calculate(node)
}

// 进行逻辑判断 > < <= >= == || &&
func (rule *Rule) calculate(node *Node) bool {
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
func (rule *Rule) getNodeValue(typeNode string, node *Node) {

	newNode := node.Left
	if typeNode == RIGHT {
		newNode = node.Right
	}

	var value interface{}

	switch newNode.Type {
	case TypeValue: // 具体值
		value = newNode.Value
		break
	case TypeParam: // 参数
		newNode.Value = rule.getParamListByName(newNode.Name)
		value = newNode.Value
		break
	case TypeFunc: // 函数
		switch len(newNode.Params) { // TODO 参数个数
		case 1:
			results, _ := callFunc(rule.FuncList, newNode.Name, newNode.Params[0])
			// TODO 抽离返回值类型
			for _, item := range results {
				switch item.Type().String() {
				case "uint64":
					newNode.Value = item.Uint()
				case "bool":
					newNode.Value = item.Bool()
				}
			}
			value = newNode.Value
		}
		break
	}

	if typeNode == LEFT {
		node.LeftResult = value
	} else {
		node.RightResult = value
	}
}
