package datasource

import (
	"fmt"
	"strings"
	"testing"
)

func Test_GetKeys(t *testing.T) {
	filter := make([]DBFilter, 3)

	for i := 0; i < 3; i++ {

		filter[i].Key = fmt.Sprintf("Key:%d", i)
		filter[i].Value = fmt.Sprintf("value:%d", i)
		filter[i].Opr = EQUALS
	}

	keysStr := strings.Join(*GetKeys(&filter), " ")
	if keysStr != "Key:0 Key:1 Key:2" {
		t.Errorf("\ngot: %s, wanted: %s", keysStr, "Key:0 Key:1 Key:2")
	}

}

func Test_GetValues(t *testing.T) {
	filter := make([]DBFilter, 3)

	for i := 0; i < 3; i++ {

		filter[i].Key = fmt.Sprintf("Key:%d", i)
		filter[i].Value = fmt.Sprintf("value:%d", i)
		filter[i].Opr = EQUALS
	}

	valuesStr := fmt.Sprintf("%v %v %v", *GetValues(&filter)...)
	if valuesStr != "value:0 value:1 value:2" {
		t.Errorf("\ngot: %s, wanted: %s", valuesStr, "value:0 value:1 value:2")
	}
}

func Test_Get2Values_with_int(t *testing.T) {
	filter := make([]DBFilter, 3)

	for i := 0; i < 3; i++ {

		filter[i].Key = fmt.Sprintf("Key:%d", i)
		filter[i].Value = i
		filter[i].Opr = EQUALS
	}

	valuesStr := fmt.Sprintf("%v %v %v", *GetValues(&filter)...)
	if valuesStr != "0 1 2" {
		t.Errorf("\ngot: %s, wanted: %s", valuesStr, "0 1 2")
	}
}

func Test_JoinQuery(t *testing.T) {
	filter := make([]DBFilter, 3)
	opr := []Operation{EQUALS, BIGGERTHAN, SMALLERTHAN}

	for i := 0; i < 3; i++ {

		filter[i].Key = fmt.Sprintf("Key:%d", i)
		filter[i].Value = fmt.Sprintf("value:%d", i)
		filter[i].Opr = opr[i]
	}

	query := JoinWhereClauses(&filter)

	if query != "Key:0 = $1 AND Key:1 > $2 AND Key:2 < $3" {
		t.Errorf("\ngot: %s, wanted: %s", query, "Key:0 = $1 AND Key:1 > $2 AND Key:2 < $3")

	}
}

func Test_Filter(t *testing.T) {
	filter := make([]DBFilter, 3)
	opr := []Operation{EQUALS, BIGGERTHAN, SMALLERTHAN}
	table := "my_table"
	for i := 0; i < 3; i++ {

		filter[i].Key = fmt.Sprintf("Key:%d", i)
		filter[i].Value = fmt.Sprintf("value:%d", i)
		filter[i].Opr = opr[i]
	}

	query, _, _ := Filter(table, &filter)

	if query != "SELECT * FROM my_table WHERE Key:0 = $1 AND Key:1 > $2 AND Key:2 < $3" {
		t.Errorf("\ngot: %s, wanted: %s", query, "SELECT * FROM my_table WHERE Key:0 = $1 AND Key:1 > $2 AND Key:2 < $3")

	}

}

func Test_Filter_null(t *testing.T) {

	table := "my_table"
	query, value, _ := Filter(table, nil)
	if query != "SELECT * FROM my_table" {
		t.Errorf("\ngot: %s, wanted: %s", query, "SELECT * FROM my_table")

	}

	if value != nil {
		t.Error("Expected value to be nil")
	}

}
