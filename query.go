package persistence

type DBQuery interface {
	ExecuteQuery(query string, response *interface{}) error
	FindAllBy(tableName string, condition DBCondition, response *interface{}) error
	FindByConditions(tableName string, conditions []DBCondition, response *interface{}) error
	FindColumnsByConditions(tableName string, columns []string, conditions []DBCondition, response *interface{})
	ExistsBy(tableName string, condition DBCondition, response *bool) error
	InsertInto(tableName string, columnValues []DBValue, response *interface{}) error
	Update(tableName string, columnValues []DBValue, conditions []DBCondition, response *interface{}) error
}

type DBCondition struct {
	FieldName string
	Operator  DBOperator
	Value     interface{}
}
type DBValue struct {
	FieldName string
	Value     interface{}
}

type DBOperator string

const (
	EQUAL     DBOperator = "eq"
	NOTEQUAL  DBOperator = "neq"
	GREATER   DBOperator = "gt"
	LESS      DBOperator = "lt"
	GREATERE  DBOperator = "gte"
	LESSE     DBOperator = "lse"
	LIKE      DBOperator = "lk"
	NOTLIKE   DBOperator = "nlk"
	IN        DBOperator = "in"
	NOTIN     DBOperator = "nin"
	ISNULL    DBOperator = "null"
	ISNOTNULL DBOperator = "nonull"
)
