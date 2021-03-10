package rule_engine

import (
	"testing"

	"github.com/likexian/gokit/assert"
)

const REJECT = "REJECT"
const Demo1 = 100
const Demo2 = true

func TestNew(t *testing.T) {

	paramList := map[string]interface{}{
		"tmpNum": 100,
	}

	funcList := map[string]interface{}{
		"requestDemo1": requestDemo1,
		"requestDemo2": requestDemo2,
	}

	nodes := make([]*Node, 0)

	// if ( requestDemo1("demo_1") > 90 || tmpNum < 1000  ) &&  requestDemo2("demo_2") {
	//   return REJECT
	// }
	tmpNode := &Node{
		Type:   TypeNode,
		Logic:  "&&",
		Result: REJECT,
		Left: &Node{
			Type:  TypeNode,
			Logic: "||",
			Left: &Node{
				Type:  TypeNode,
				Logic: ">",
				Left: &Node{
					Type:   TypeFunc,
					Name:   "requestDemo1",
					Params: []interface{}{"demo_1"},
				},
				Right: &Node{
					Type:  TypeValue,
					Value: 90,
				},
			},
			Right: &Node{
				Type:  TypeNode,
				Logic: "<",
				Left: &Node{
					Type: TypeParam,
					Name: "tmpNum",
				},
				Right: &Node{
					Type:  TypeValue,
					Value: 1000,
				},
			},
		},
		Right: &Node{
			Type:  TypeNode,
			Logic: "bool",
			Left: &Node{
				Type:   TypeFunc,
				Name:   "requestDemo2",
				Params: []interface{}{"demo_2"},
			},
		},
	}

	nodes = append(nodes, tmpNode)

	rule := New()
	rule.SetParamList(paramList)
	rule.SetFuncList(funcList)
	rule.SetNodes(nodes)
	result := rule.Parser()
	//fmt.Println(result)

	if result == nil {
		result = ""
	}

	assert.Equal(t, result.(string), REJECT)
}

// mock1
func requestDemo1(sql string) uint64 {
	// 以下是 mock
	switch sql {
	case "demo_1":
		return mockDemo1()
	default:
		break
	}
	return 0
}

// mock1
func mockDemo1() uint64 {
	return Demo1
}

// mock2
func requestDemo2(command string) bool {
	switch command {
	case "demo_2":
		return mockDemo2()
	default:
		break
	}
	return false
}

// mock2
func mockDemo2() bool {
	return Demo2
}
