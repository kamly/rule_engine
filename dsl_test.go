package dsl_engine

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
		Logic:  "&&",
		Result: REJECT,
		Left: &Node{
			Logic: "||",
			Left: &Node{
				Logic: ">",
				Left: NodeFunc{
					Name:   "requestDemo1",
					Params: []interface{}{"demo_1"},
				},
				Right: NodeValue{
					Value: 90,
				},
			},
			Right: &Node{
				Logic: "<",
				Left: NodeParam{
					Name: "tmpNum",
				},
				Right: NodeValue{
					Value: 1000,
				},
			},
		},
		Right: &Node{
			Logic: "bool",
			Left: NodeFunc{
				Name:   "requestDemo2",
				Params: []interface{}{"demo_2"},
			},
		},
	}

	nodes = append(nodes, tmpNode)

	dsl := New()
	dsl.SetParamList(paramList)
	dsl.SetFuncList(funcList)
	dsl.SetNodes(nodes)
	result := dsl.Parser()
	//fmt.Println(result)

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
