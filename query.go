package persistence

type DBQuery interface {
	//ExecuteQuery response Param must be a pointer to Slice of struct
	ExecuteQuery(query string, response interface{}) error
	//FindAllBy response Param must be a pointer to Slice of struct
	FindAllBy(tableName string, condition DBCondition, response interface{}) error
	//FindByConditions response Param must be a pointer to Slice of struct
	FindByConditions(tableName string, conditions []DBCondition, response interface{}) error
	//FindColumnsByConditions response Param must be a pointer to Slice of struct
	FindColumnsByConditions(tableName string, columns []string, conditions []DBCondition, response interface{}) error
	ExistsBy(tableName string, condition DBCondition) (bool, error)
	InsertInto(tableName string, columnValues []DBValue) error
	Update(tableName string, columnValues []DBValue, conditions []DBCondition) error
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
