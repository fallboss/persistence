package persistence

import (
	"fmt"
	"github.com/georgysavva/scany/pgxscan"
	"reflect"
	"strings"
)

func GetPgRepo(cnt *PgClient) DBQuery {
	return &pgRepo{cnt}
}

type pgRepo struct {
	cnt *PgClient
}

const queryExists = "SELECT EXISTS (SELECT 1 FROM %s WHERE %s %s $1)"

func (r *pgRepo) FindColumnsByConditions(tableName string, columns []string, conditions []DBCondition, response interface{}) error {
	var conditionsQuery = ""

	var args []interface{}
	for i, condition := range conditions {
		if !strings.EqualFold(conditionsQuery, "") {
			conditionsQuery = conditionsQuery + " AND "
		}
		op, err := parseOperator(condition.Operator)
		if err != nil {
			return err
		}
		values := fmt.Sprintf("$%d", i+1)
		conditionsQuery = conditionsQuery + fmt.Sprintf("%s %s %s", condition.FieldName, op, values)
		args = append(args, condition.Value)
	}
	var query = fmt.Sprintf("SELECT %s FROM %s WHERE %s", strings.Join(columns, ","), tableName, conditionsQuery)
	rows, err := r.cnt.RunQueryArgs(query, args)
	if err != nil {
		return err
	}
	return pgxscan.ScanAll(response, rows)
}

func (r *pgRepo) ExecuteQuery(query string, response interface{}) error {
	rows, err := r.cnt.RunQueryRows(query)
	if err != nil {
		return err
	}
	return pgxscan.ScanAll(response, rows)
}

func (r *pgRepo) FindAllBy(tableName string, condition DBCondition, response interface{}) error {
	op, err := parseOperator(condition.Operator)
	if err != nil {
		return err
	}
	var query = fmt.Sprintf("SELECT * FROM %s WHERE %s %s $1", tableName, condition.FieldName, op)
	args := []interface{}{condition.Value}
	rows, err := r.cnt.RunQueryArgs(query, args)
	if err != nil {
		return err
	}
	return pgxscan.ScanAll(response, rows)
}

func (r *pgRepo) FindByConditions(tableName string, conditions []DBCondition, response interface{}) error {
	var conditionsQuery = ""
	var args []interface{}
	for i, condition := range conditions {
		if !strings.EqualFold(conditionsQuery, "") {
			conditionsQuery = conditionsQuery + " AND "
		}
		op, err := parseOperator(condition.Operator)
		if err != nil {
			return err
		}
		value := fmt.Sprintf("$%d", i+1)
		conditionsQuery = conditionsQuery + fmt.Sprintf("%s %s %s", condition.FieldName, op, value)
		args = append(args, condition.Value)
	}
	var query = fmt.Sprintf("SELECT * FROM %s WHERE %s", tableName, conditionsQuery)
	rows, err := r.cnt.RunQueryArgs(query, args)
	if err != nil {
		return err
	}
	return pgxscan.ScanAll(response, rows)
}

func (r *pgRepo) ExistsBy(tableName string, condition DBCondition) (bool, error) {
	op, err := parseOperator(condition.Operator)
	if err != nil {
		return false, err
	}
	var query = fmt.Sprintf(queryExists, tableName, condition.FieldName, op)
	args := []interface{}{condition.Value}
	rows, err := r.cnt.RunQueryArgs(query, args)
	if err != nil {
		return false, err
	}
	var exists = struct {
		Exists bool
	}{}
	err = pgxscan.ScanOne(&exists, rows)
	return exists.Exists, err
}

func (r *pgRepo) InsertInto(tableName string, columnValues []DBValue) error {
	var columnNames []string
	var values []string
	var args []interface{}
	for i, column := range columnValues {
		columnNames = append(columnNames, column.FieldName)
		values = append(values, fmt.Sprintf("$%d", i+1))
		args = append(args, column.Value)
	}
	var columns = strings.Join(columnNames, ",")
	var query = fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s) ", tableName, columns, strings.Join(values, ","))
	rows, err := r.cnt.RunQueryArgs(query, args)
	rows.Close()
	return err
}

func (r *pgRepo) Update(tableName string, columnValues []DBValue, conditions []DBCondition) error {
	var columnQuery = ""
	var args []interface{}
	index := 0
	for _, column := range columnValues {
		if !strings.EqualFold(columnQuery, "") {
			columnQuery = columnQuery + " , "
		}
		index = index + 1
		columnQuery = columnQuery + fmt.Sprintf("%s = $%d", column.FieldName, index)
		args = append(args, column.Value)
	}
	var conditionsQuery = ""
	for _, condition := range conditions {
		if !strings.EqualFold(conditionsQuery, "") {
			conditionsQuery = conditionsQuery + " AND "
		}
		op, err := parseOperator(condition.Operator)
		if err != nil {
			return err
		}
		if condition.Operator == ISNULL || condition.Operator == ISNOTNULL {
			conditionsQuery = conditionsQuery + fmt.Sprintf("%s %s", condition.FieldName, op)
		} else if condition.Value == nil || (reflect.ValueOf(condition.Value).Kind() == reflect.Ptr && reflect.ValueOf(condition.Value).IsNil()) {
			conditionsQuery = conditionsQuery + fmt.Sprintf("%s IS NULL", condition.FieldName)
		} else {
			index = index + 1
			conditionsQuery = conditionsQuery + fmt.Sprintf("%s %s $%d", condition.FieldName, op, index)
			args = append(args, condition.Value)
		}
	}
	var query = fmt.Sprintf("UPDATE %s SET %s WHERE %s", tableName, columnQuery, conditionsQuery)
	rows, err := r.cnt.RunQueryArgs(query, args)
	rows.Close()
	return err
}

func parseOperator(op DBOperator) (string, error) {
	switch op {
	case EQUAL:
		return "=", nil
	case NOTEQUAL:
		return "<>", nil
	case GREATER:
		return ">", nil
	case LESS:
		return "<", nil
	case GREATERE:
		return ">=", nil
	case LESSE:
		return "<=", nil
	case LIKE:
		return "LIKE", nil
	case NOTLIKE:
		return "NOT LIKE", nil
	case IN:
		return "IN", nil
	case NOTIN:
		return "NOT IN", nil
	case ISNULL:
		return "IS NULL", nil
	case ISNOTNULL:
		return "IS NOT NULL", nil
	default:
		return "", fmt.Errorf("operator '%s' not suported", op)
	}
}
