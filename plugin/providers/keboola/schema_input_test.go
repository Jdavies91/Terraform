package keboola

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMappingFromInputToSchema(t *testing.T) {
	inputs := make([]interface{}, 0, 1)

	testInput := map[string]interface{}{
		"source":        "test source",
		"destination":   "test destination",
		"whereOperator": "test operator",
		"whereColumn":   "test column",
		"datatypes": map[string]interface{}{
			"foo": "bar",
		},
		"days": 2,
	}

	inputs = append(inputs, testInput)
	result := mapInputSchemaToModel(inputs)

	assert.Equal(t, testInput["source"].(string), result[0].Source, "Original source and mapped source should match")
	assert.Equal(t, testInput["destination"].(string), result[0].Destination, "Original destination and mapped destination should match")
	assert.Equal(t, testInput["whereOperator"].(string), result[0].WhereOperator, "Original whereOperator and mapped whereOperator should match")
	assert.Equal(t, testInput["whereColumn"].(string), result[0].WhereColumn, "Original whereColumn and mapped whereColumn should match")
	assert.Equal(t, testInput["datatypes"].(map[string]interface{}), result[0].DataTypes, "Original datatypes and mapped datatypes should match")
	assert.Equal(t, testInput["days"].(int), result[0].Days, "Original days and mapped days should match")
}