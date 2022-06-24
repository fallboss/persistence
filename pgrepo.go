package persistence

import (
	"fmt"
	"strings"
)

func GetPgRepo(cnt *PgClient) DBQuery {
	return &pgRepo{cnt}
}

type pgRepo struct {
	cnt *PgClient
}

func (r *pgRepo) FindColumnsByConditions(tableName string, columns []string, conditions []DBCondition, response interface{}) error {
	var conditionsQuery = ""
	for _, condition := range conditions {
		if !strings.EqualFold(conditionsQuery, "") {
			conditionsQuery = conditionsQuery + " AND "
		}
		op, err := parseOperator(condition.Operator)
		if err != nil {
			return err
		}
		conditionsQuery = conditionsQuery + fmt.Sprintf("(%s %s '%s')", condition.FieldName, op, condition.Value)
	}
	var query = fmt.Sprintf("SELECT %s FROM %s WHERE %s", strings.Join(columns, ","), tableName, conditionsQuery)
	row, err := r.cnt.RunQueryRows(query)
	if err != nil {
		return err
	}
	return (*row).Scan(&response)
}

func (r *pgRepo) ExecuteQuery(query string, response interface{}) error {
	row, err := r.cnt.RunQueryRows(query)
	if err != nil {
		return err
	}
	return (*row).Scan(&response)
}

func (r *pgRepo) FindAllBy(tableName string, condition DBCondition, response interface{}) error {
	op, err := parseOperator(condition.Operator)
	if err != nil {
		return err
	}
	var query = fmt.Sprintf("SELECT * FROM %s WHERE %s %s '%s'", tableName, condition.FieldName, op, condition.Value)
	row, err := r.cnt.RunQueryRows(query)
	if err != nil {
		return err
	}
	return (*row).Scan(&response)
}

func (r *pgRepo) FindByConditions(tableName string, conditions []DBCondition, response interface{}) error {
	var conditionsQuery = ""
	for _, condition := range conditions {
		if !strings.EqualFold(conditionsQuery, "") {
			conditionsQuery = conditionsQuery + " AND "
		}
		op, err := parseOperator(condition.Operator)
		if err != nil {
			return err
		}
		conditionsQuery = conditionsQuery + fmt.Sprintf("(%s %s '%s')", condition.FieldName, op, condition.Value)
	}
	var query = fmt.Sprintf("SELECT * FROM %s WHERE %s", tableName, conditionsQuery)
	row, err := r.cnt.RunQueryRows(query)
	if err != nil {
		return err
	}
	return (*row).Scan(&response)
}

func (r *pgRepo) ExistsBy(tableName string, condition DBCondition, response interface{}) error {
	op, err := parseOperator(condition.Operator)
	if err != nil {
		return err
	}
	var query = fmt.Sprintf("SELECT 1 CANT FROM %s WHERE %s %s '%s'", tableName, condition.FieldName, op, condition.Value)
	//query = fmt.Sprintf("SELECT exists (%s)", query)
	row, err := r.cnt.RunQueryRows(query)
	if err != nil {
		return err
	}
	for (*row).Next() {
		response, err = (*row).Values()
		return err
	}
	return err
}

func (r *pgRepo) InsertInto(tableName string, columnValues []DBValue, response interface{}) error {
	var columnNames []string
	var values []interface{}
	for _, column := range columnValues {
		columnNames = append(columnNames, column.FieldName)
		values = append(values, column.Value)
	}
	var query = fmt.Sprintf("INSERT INTO %s (%s) VALUES", tableName, strings.Join(columnNames, ","))
	//var queryValues = query + " (%s) RETURNING *"
	row, err := r.cnt.RunQueryArgs(query, values)
	if err != nil {
		return err
	}
	return (*row).Scan(&response)
}

func (r *pgRepo) Update(tableName string, columnValues []DBValue, conditions []DBCondition, response interface{}) error {
	var columnQuery = ""
	for _, column := range columnValues {
		if !strings.EqualFold(columnQuery, "") {
			columnQuery = columnQuery + " , "
		}
		columnQuery = columnQuery + fmt.Sprintf("%s = %s", column.FieldName, column.Value)
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
		conditionsQuery = conditionsQuery + fmt.Sprintf("%s %s %s", condition.FieldName, op, condition.Value)
	}
	var query = fmt.Sprintf("UPDATE %s SET %s WHERE %s ", tableName, columnQuery, conditionsQuery)
	row := r.cnt.RunQueryRow(query)
	return (*row).Scan(&response)
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
