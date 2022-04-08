package datasource

import (
	"fmt"
	"strings"
)

type Operation int

const (
	EQUALS Operation = iota
	SMALLERTHAN
	BIGGERTHAN
)

func (o Operation) String() string {
	switch o {
	case EQUALS:
		return "="
	case SMALLERTHAN:
		return "<"
	case BIGGERTHAN:
		return ">"
	}
	return "unknown"
}

type DBFilter struct {
	Key   string
	Value any
	Opr   Operation
}

func Filter(table string, filters *[]DBFilter) (string, *[]any, error) {
	query := fmt.Sprintf("SELECT * FROM %s", table)
	var values *[]any
	if filters != nil {
		values = GetValues(filters)

		query = fmt.Sprintf("%s WHERE %s", query, JoinWhereClauses(filters))
	}

	return query, values, nil
}

func GetKeys(filters *[]DBFilter) *[]string {
	keys := make([]string, len(*filters))

	i := 0
	for _, k := range *filters {
		keys[i] = k.Key
		i++
	}

	return &keys
}

func GetValues(filters *[]DBFilter) *[]any {
	values := make([]any, len(*filters))
	i := 0
	for k := range *filters {
		values[i] = (*filters)[k].Value
		i++
	}

	return &values
}

func JoinWhereClauses(filter *[]DBFilter) string {
	whereClauses := make([]string, len(*filter))

	for i := range *filter {
		whereClauses[i] = fmt.Sprintf("%s %s $%d", (*filter)[i].Key, (*filter)[i].Opr.String(), i+1)
	}
	return strings.Join(whereClauses, " AND ")
}
