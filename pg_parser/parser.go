package pgparser

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	pgQuery "github.com/pganalyze/pg_query_go/v5"
)

type (
	Node                     = pgQuery.Node
	NodeVariableShowStmt     = pgQuery.Node_VariableShowStmt
	NodeCreateStmt           = pgQuery.Node_CreateStmt
	NodeConstraint           = pgQuery.Node_Constraint
	NodeCreateTableSpaceStmt = pgQuery.Node_CreateTableSpaceStmt
	NodeDropTableSpaceStmt   = pgQuery.Node_DropTableSpaceStmt
	NodeString               = pgQuery.Node_String_
)

type Table struct {
	Name    string `json:"name"`
	Columns []Column
}

type Column struct {
	Name        string     `json:"name"`
	DataType    PgDataType `json:"dataType"`
	Meta        string     `json:"meta"`
	Nullable    *bool      `json:"nullable"`
	PrimaryKey  *bool      `json:"primaryKey"`
	Constraints []string
}

func ParseToStruct(sql string) ([]Table, error) {
	tree, err := pgQuery.Parse(sql)
	if err != nil {
		return nil, err
	}

	tables := walk(tree)

	// log.Println(pgQuery.Deparse(tree))

	result, _ := pgQuery.Scan(sql)
	Start := result.Tokens[0].Start
	End := result.Tokens[0].End
	log.Println(sql[Start:End])
	log.Println(result.Tokens[0].GetEnd())

	if len(tables) == 0 {
		return nil, errors.New("unknown error")
	}

	return tables, nil
}

func ParseToJson(sql string) ([]byte, error) {
	tables, err := ParseToStruct(sql)
	if err != nil {
		return nil, err
	}

	data, err := json.MarshalIndent(tables, "", strings.Repeat(" ", 4))
	if err != nil {
		return nil, err
	}

	return data, nil
}

func walk(tree *pgQuery.ParseResult) []Table {
	stmts := tree.GetStmts()
	{
		if len(stmts) == 0 {
			return nil
		}
	}

	var tables []Table

	for _, stmt := range stmts {

		rawStmt := stmt.GetStmt()
		{
			if rawStmt == nil {
				continue
			}
		}

		createStml := rawStmt.GetCreateStmt()
		{
			if createStml == nil {
				continue
			}
		}

		var (
			columns   []Column
			relation  = createStml.GetRelation()
			tablename = relation.GetRelname()
		)
		_ = tablename
		_ = columns
		columnsDef := tableColumns(createStml)
		{
			for _, col := range columnsDef {
				constraintsColumn(col)
			}
		}

		// tables = append(tables, Table{
		// 	Name:    tablename,
		// 	Columns: columns,
		// })
	}

	return tables
}

func tableColumns(createStmt *pgQuery.CreateStmt) []*pgQuery.ColumnDef {
	var (
		cols       = createStmt.GetTableElts()
		columnsDef []*pgQuery.ColumnDef
	)

	for _, columnDef := range cols {
		def := columnDef.GetColumnDef()
		{
			if def != nil {
				columnsDef = append(columnsDef, def)
			}
		}
	}

	return columnsDef
}

func constraintsColumn(columnDef *pgQuery.ColumnDef) []string {
	nodes := columnDef.GetConstraints()
	{
		if len(nodes) == 0 {
			return nil
		}
	}

	var constraints []string

	for _, node := range nodes {
		constraint := node.GetConstraint()
		{
			if constraint == nil {
				continue
			}
		}

		switch ConstrType(constraint.Contype) {
		case CONSTR_PRIMARY:
			constraints = append(constraints, "PRIMARY KEY")
		case CONSTR_NULL:
			constraints = append(constraints, "NULL")
		case CONSTR_NOTNULL:
			constraints = append(constraints, "NOT NULL")
		case CONSTR_UNIQUE:
			constraints = append(constraints, "UNIQUE")
		case CONSTR_DEFAULT:
			data, _ := json.Marshal(constraint.GetRawExpr())
			log.Println(string(data))
			expr := constraint.GetRawExpr()
			{
				if expr == nil {
					continue
				}
			}
			// data, _ := json.Marshal(expr)

			// log.Println(string(data))
			_default := defaultValue(expr)
			{
				if len(_default) > 0 {
					constraints = append(constraints, "DEFAULT "+_default)
					log.Println(constraints)
				}
			}

		case CONSTR_CHECK:

		}
	}

	return constraints
}

func defaultValue(node *pgQuery.Node) string {
	switch exp := node.Node.(type) {
	case *pgQuery.Node_AConst:
		switch e := exp.AConst.Val.(type) {
		case *pgQuery.A_Const_Fval:
			return e.Fval.GetFval()
		case *pgQuery.A_Const_Boolval:
			return strconv.FormatBool(e.Boolval.GetBoolval())
		case *pgQuery.A_Const_Bsval:
			return e.Bsval.GetBsval()
		case *pgQuery.A_Const_Ival:
			return strconv.FormatInt(int64(e.Ival.GetIval()), 10)
		case *pgQuery.A_Const_Sval:
			return fmt.Sprintf("'%s'", e.Sval.GetSval())
		}
	case *pgQuery.Node_FuncCall:
		funcCall := exp.FuncCall
		funcnames := funcCall.GetFuncname()[0]
		funcname := funcnames.GetString_().Sval

		var st []string

		for _, arg := range funcCall.GetArgs() {
			st = append(st, defaultValue(arg))
		}

		return fmt.Sprintf("%s(%s)", funcname, strings.Join(st, " "))
	case *pgQuery.Node_AExpr:
		aexpr := exp.AExpr

		op := aexpr.Name[0].GetString_().Sval

		lexpr := aexpr.Lexpr

		rexpr := aexpr.Rexpr

		return fmt.Sprintf("(%s%s%s)", defaultValue(lexpr), op, defaultValue(rexpr))
	case *pgQuery.Node_ColumnRef:

		node := exp.ColumnRef.Fields[0]

		return node.GetString_().Sval
	}

	return ""
}
