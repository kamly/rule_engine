package rule_engine

type Rule struct {
	Nodes     []*Node
	ParamList map[string]interface{}
	FuncList  map[string]interface{}
}

func New() *Rule {
	return new(Rule)
}

func (rule *Rule) Parser() interface{} {
	for _, node := range rule.Nodes {
		result := rule.getNextNode(node)
		if result {
			return node.Result
		}
	}
	return nil
}

func (rule *Rule) SetNodes(nodes []*Node) {
	rule.Nodes = nodes
}

func (rule *Rule) SetParamList(paramList map[string]interface{}) {
	rule.ParamList = paramList
}

func (rule *Rule) getParamListByName(name string) interface{} {
	return rule.ParamList[name]
}

func (rule *Rule) SetFuncList(funcList map[string]interface{}) {
	rule.FuncList = funcList
}
