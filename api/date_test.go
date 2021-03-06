package api_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"go.bmvs.io/ynab/api"
)

func TestDate_UnmarshalJSON(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		wrapper := struct {
			Name string
			Date *api.Date
		}{}

		err := json.Unmarshal([]byte(`{"Name": "YNAB", "Date": "2009-01-29"}`), &wrapper)
		assert.NoError(t, err)
		assert.Equal(t, "2009-01-29 00:00:00 +0000 UTC", wrapper.Date.String())
	})

	t.Run("json string without Date field", func(t *testing.T) {
		wrapper := struct {
			Name string
			Date *api.Date
		}{}

		err := json.Unmarshal([]byte(`{"Name": "YNAB"}`), &wrapper)
		assert.NoError(t, err)
		assert.Nil(t, wrapper.Date)
	})

	t.Run("json string with null Date field", func(t *testing.T) {
		wrapper := struct {
			Name string
			Date *api.Date
		}{}

		err := json.Unmarshal([]byte(`{"Name": "YNAB", "Date": null}`), &wrapper)
		assert.NoError(t, err)
		assert.Nil(t, wrapper.Date)
	})
}

func TestDate_MarshalJSON(t *testing.T) {
	date, err := api.DateFromString("2020-01-20")
	assert.NoError(t, err)

	wrapper := struct {
		Date api.Date
	}{
		date,
	}

	buf, err := json.Marshal(&wrapper)
	assert.NoError(t, err)
	assert.Equal(t, `{"Date":"2020-01-20"}`, string(buf))
}

func TestNewDateFromString(t *testing.T) {
	table := []struct {
		InputDate          string
		OutputDateToString string
		OutputError        bool
	}{
		{"2018-02-01", "2018-02-01 00:00:00 +0000 UTC", false},
		{"2018-13-01", "0001-01-01 00:00:00 +0000 UTC", true},
	}

	for _, test := range table {
		date, err := api.DateFromString(test.InputDate)
		assert.Equal(t, test.OutputError, err != nil)
		assert.Equal(t, test.OutputDateToString, date.String())
	}
}
