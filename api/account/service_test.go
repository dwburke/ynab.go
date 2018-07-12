package account_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/jarcoal/httpmock.v1"

	"bmvs.io/ynab"
	"bmvs.io/ynab/api/account"
)

func TestService_GetAccounts(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://api.youneedabudget.com/v1/budgets/bbdccdb0-9007-42aa-a6fe-02a3e94476be/accounts",
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewStringResponse(200, `{
  "data": {
    "accounts": [
			{
				"id": "aa248caa-eed7-4575-a990-717386438d2c",
				"name": "Test Account 2",
				"type": "savings",
				"on_budget": false,
				"closed": true,
				"note": "omg omg omg",
				"balance": -123930,
				"cleared_balance": -123930,
				"uncleared_balance": 0,
				"deleted": false
			}
    ]
  }
}
		`), nil
		},
	)

	client := ynab.NewClient("")
	accounts, err := client.Account().GetAccounts("bbdccdb0-9007-42aa-a6fe-02a3e94476be")
	assert.NoError(t, err)

	note := "omg omg omg"
	expected := &account.Account{
		ID:               "aa248caa-eed7-4575-a990-717386438d2c",
		Name:             "Test Account 2",
		Type:             account.TypeSavings,
		OnBudget:         false,
		Closed:           true,
		Note:             &note,
		Balance:          int64(-123930),
		ClearedBalance:   int64(-123930),
		UnclearedBalance: int64(0),
		Deleted:          false,
	}
	assert.Equal(t, expected, accounts[0])
}

func TestService_GetAccountByID(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://api.youneedabudget.com/v1/budgets/bbdccdb0-9007-42aa-a6fe-02a3e94476be/accounts/aa248caa-eed7-4575-a990-717386438d2c",
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewStringResponse(200, `{
  "data": {
    "account": {
      "id": "aa248caa-eed7-4575-a990-717386438d2c",
      "name": "Test Account",
      "type": "checking",
      "on_budget": true,
      "closed": true,
			"note": "omg omg omg",
      "balance": 0,
      "cleared_balance": 0,
      "uncleared_balance": 0,
      "deleted": false
    }
  }
}
		`), nil
		},
	)

	client := ynab.NewClient("")
	a, err := client.Account().GetAccountByID(
		"bbdccdb0-9007-42aa-a6fe-02a3e94476be",
		"aa248caa-eed7-4575-a990-717386438d2c",
	)
	assert.NoError(t, err)

	note := "omg omg omg"
	expected := &account.Account{
		ID:               "aa248caa-eed7-4575-a990-717386438d2c",
		Name:             "Test Account",
		Type:             account.TypeChecking,
		OnBudget:         true,
		Note:             &note,
		Closed:           true,
		Balance:          int64(0),
		ClearedBalance:   int64(0),
		UnclearedBalance: int64(0),
		Deleted:          false,
	}
	assert.Equal(t, expected, a)
}