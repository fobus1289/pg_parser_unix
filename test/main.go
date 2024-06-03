package main

import (
	"encoding/json"
	"fmt"
)

// Структура для представления узла JSON
type Node struct {
	AExpr *AExprNode `json:"AExpr"`
}

// Структура для представления узла AExpr
type AExprNode struct {
	Kind     int        `json:"kind"`
	Name     []NameNode `json:"name"`
	Lexpr    *Node      `json:"lexpr"`
	Rexpr    *Node      `json:"rexpr"`
	Location int        `json:"location"`
}

// Структура для представления узла Name
type NameNode struct {
	Node *StringNode `json:"Node"`
}

// Структура для представления узла String
type StringNode struct {
	Sval string `json:"sval"`
}

// Функция для преобразования узлов в математическое выражение
func nodeToMathExpression(node *Node) string {

	if node == nil || node.AExpr == nil {
		return ""
	}

	// Рекурсивно обрабатываем левое и правое подвыражение
	leftExpr := nodeToMathExpression(node.AExpr.Lexpr)
	rightExpr := nodeToMathExpression(node.AExpr.Rexpr)

	// Строим математическое выражение на основе текущего узла
	expression := ""
	if leftExpr != "" {
		expression += fmt.Sprintf("(%s", leftExpr)
	} else {
		expression += "("
	}

	expression += fmt.Sprintf(" %s ", getOperator(node.AExpr.Name))

	if rightExpr != "" {
		expression += fmt.Sprintf("%s)", rightExpr)
	} else {
		expression += ")"
	}

	return expression
}

// Функция для получения оператора из узла Name
func getOperator(nameNodes []NameNode) string {
	if len(nameNodes) > 0 && nameNodes[0].Node != nil {
		return nameNodes[0].Node.Sval
	}
	return ""
}

func main() {
	// JSON-подобная строка с вашей структурой данных
	jsonData := `{"Node":{"AExpr":{"kind":1,"name":[{"Node":{"String_":{"sval":"+"}}}],"lexpr":{"Node":{"AExpr":{"kind":1,"name":[{"Node":{"String_":{"sval":"+"}}}],"lexpr":{"Node":{"AExpr":{"kind":1,"name":[{"Node":{"String_":{"sval":"+"}}}],"lexpr":{"Node":{"AConst":{"Val":{"Ival":{"ival":1}},"location":393}}},"rexpr":{"Node":{"AConst":{"Val":{"Ival":{"ival":2}},"location":395}}},"location":394}}},"rexpr":{"Node":{"AExpr":{"kind":1,"name":[{"Node":{"String_":{"sval":"+"}}}],"lexpr":{"Node":{"AConst":{"Val":{"Ival":{"ival":1}},"location":398}}},"rexpr":{"Node":{"AConst":{"Val":{"Ival":{"ival":1}},"location":400}}},"location":399}}},"location":396}}},"rexpr":{"Node":{"AExpr":{"kind":1,"name":[{"Node":{"String_":{"sval":"+"}}}],"lexpr":{"Node":{"AConst":{"Val":{"Ival":{"ival":3}},"location":405}}},"rexpr":{"Node":{"AExpr":{"kind":1,"name":[{"Node":{"String_":{"sval":"*"}}}],"lexpr":{"Node":{"AConst":{"Val":{"Ival":{"ival":4}},"location":407}}},"rexpr":{"Node":{"AConst":{"Val":{"Ival":{"ival":5}},"location":409}}},"location":408}}},"location":406}}},"location":403}}}`

	// Разбор JSON данных в структуру Node
	var nodeData map[string]Node
	err := json.Unmarshal([]byte(jsonData), &nodeData)
	if err != nil {
		fmt.Printf("Error parsing JSON: %v\n", err)
		return
	}
	// log.Println(jsonData)
	// Получение корневого узла Node
	rootNode := nodeData["Node"]

	// Вызов функции для преобразования в математическое выражение
	mathExpression := nodeToMathExpression(&rootNode)

	// Вывод математического выражения
	fmt.Printf("Математическое выражение: %s\n", mathExpression)
}
