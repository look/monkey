package evaluator

import (
	"monkey/ast"
	"monkey/object"
)

var (
	NULL  = &object.Null{}
	TRUE  = &object.Boolean{Value: true}
	FALSE = &object.Boolean{Value: false}
)

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	// statements
	case *ast.Program:
		return evalStatements(node.Statements)
	case *ast.ExpressionStatement:
		return Eval(node.Expression)
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	case *ast.Boolean:
		return nativeBoolToBooleanObject(node.Value)
	case *ast.PrefixExpression:
		right := Eval(node.Right)
		return evalPrefixExpression(node.Operator, right)
	case *ast.InfixExpression:
		left := Eval(node.Left)
		right := Eval(node.Right)
		return evalInfixExpression(node.Operator, left, right)
	case *ast.BlockStatement:
		return evalStatements(node.Statements)
	case *ast.IfExpression:
		return evalIfExpression(node)
	}

	return nil
}

func evalStatements(statements []ast.Statement) object.Object {
	var result object.Object

	for _, statement := range statements {
		result = Eval(statement)
	}

	return result
}

func evalIfExpression(ifExp *ast.IfExpression) object.Object {
	cond := Eval(ifExp.Condition)
	if isTruthy(cond) {
		return Eval(ifExp.Consequence)
	} else if ifExp.Alternative != nil {
		return Eval(ifExp.Alternative)
	} else {
		return NULL
	}
}

func nativeBoolToBooleanObject(value bool) object.Object {
	if value {
		return TRUE
	}

	return FALSE
}

func evalPrefixExpression(op string, right object.Object) object.Object {
	switch op {
	case "!":
		return evalBangOperator(right)
	case "-":
		return evalMinusPrefixOperator(right)
	default:
		return NULL
	}
}

func evalBangOperator(right object.Object) object.Object {
	if isTruthy(right) {
		return FALSE
	}

	return TRUE
}

func evalMinusPrefixOperator(right object.Object) object.Object {
	if right.Type() != object.INTEGER_OBJ {
		return NULL
	}

	value := right.(*object.Integer).Value
	return &object.Integer{Value: -value}
}

func evalInfixExpression(op string, left, right object.Object) object.Object {
	switch {
	case left.Type() == object.INTEGER_OBJ && right.Type() == object.INTEGER_OBJ:
		return evalIntegerInfixExpression(op, left, right)
	case op == "==":
		return nativeBoolToBooleanObject(left == right)
	case op == "!=":
		return nativeBoolToBooleanObject(left != right)
	default:
		return NULL
	}
}

func evalIntegerInfixExpression(op string, left, right object.Object) object.Object {
	l := left.(*object.Integer).Value
	r := right.(*object.Integer).Value

	switch op {
	case "+":
		return &object.Integer{Value: l + r}
	case "-":
		return &object.Integer{Value: l - r}
	case "*":
		return &object.Integer{Value: l * r}
	case "/":
		return &object.Integer{Value: l / r}
	case "<":
		return nativeBoolToBooleanObject(l < r)
	case ">":
		return nativeBoolToBooleanObject(l > r)
	case "==":
		return nativeBoolToBooleanObject(l == r)
	case "!=":
		return nativeBoolToBooleanObject(l != r)
	default:
		return NULL
	}
}

func isTruthy(obj object.Object) bool {
	switch obj {
	case FALSE, NULL:
		return false
	default:
		return true
	}
}
