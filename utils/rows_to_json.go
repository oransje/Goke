package utils

import (
	"database/sql"
	"encoding/json"
	"strconv"
)

func Jsonify(rows *sql.Rows) []map[string]interface{} {

	columns, err := rows.Columns()

	HandleError(err, "Error while database reading columns")

	values := make([]interface{}, len(columns))

	scanArgs := make([]interface{}, len(values))

	for i := range values {
		scanArgs[i] = &values[i]
	}

	results := make(map[string]interface{})
	var data []map[string]interface{}

	for rows.Next() {

		err = rows.Scan(scanArgs...)
		HandleError(err, "Error while database scanning rows")

		for i, value := range values {
			switch value.(type) {
			case nil:
				results[columns[i]] = nil

			case []byte:
				s := string(value.([]byte))
				x, err := strconv.Atoi(s)

				if err != nil {
					results[columns[i]] = s
				} else {
					results[columns[i]] = x
				}

			default:
				results[columns[i]] = value
			}
		}

		b, _ := json.Marshal(results)
		var m map[string]interface{}
		json.Unmarshal(b, &m)
		data = append(data, m)

	}

	return data
}
