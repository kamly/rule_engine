package risk_dsl

type Dsl struct {
	Nodes     []*Node
	ParamList map[string]interface{}
	FuncList  map[string]interface{}
}

func New() *Dsl {
	return new(Dsl)
}

func (dsl *Dsl) Parser() interface{} {
	for _, node := range dsl.Nodes {
		result := dsl.getNextNode(node)
		if result {
			return node.Result
		}
	}
	return nil
}

func (dsl *Dsl) SetNodes(nodes []*Node) {
	dsl.Nodes = nodes
}

func (dsl *Dsl) SetParamList(paramList map[string]interface{}) {
	dsl.ParamList = paramList
}

func (dsl *Dsl) getParamListByName(name string) interface{} {
	return dsl.ParamList[name]
}

func (dsl *Dsl) SetFuncList(funcList map[string]interface{}) {
	dsl.FuncList = funcList
}
