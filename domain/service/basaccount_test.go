package service

import (
	"encoding/json"
	"fmt"
	"omega/domain/base/basmodel"
	"omega/domain/base/enum/accounttype"
	"omega/internal/types"
	"omega/pkg/helper"
	"testing"
)

func TestChartOfAccounts(t *testing.T) {
	accounts := []basmodel.Account{
		{
			FixedNode: types.FixedNode{
				ID: 1,
			},
			Code: helper.StrPointer("1"),
			Name: "Asset",
			Type: accounttype.Asset,
		},
		{
			FixedNode: types.FixedNode{
				ID: 2,
			},
			ParentID: types.RowIDPointer(1),
			Code:     helper.StrPointer("11"),
			Name:     "Cash USD",
			Type:     accounttype.Cash,
		},
		{
			FixedNode: types.FixedNode{
				ID: 3,
			},
			ParentID: types.RowIDPointer(1),
			Code:     helper.StrPointer("12"),
			Name:     "Cash IQD",
			Type:     accounttype.Cash,
		},
		{
			FixedNode: types.FixedNode{
				ID: 4,
			},
			Code: helper.StrPointer("3"),
			Name: "Expense",
			Type: accounttype.Expense,
		},
		{
			FixedNode: types.FixedNode{
				ID: 5,
			},
			ParentID: types.RowIDPointer(4),
			Code:     helper.StrPointer("31"),
			Name:     "Building",
			Type:     accounttype.Expense,
		},
		{
			FixedNode: types.FixedNode{
				ID: 6,
			},
			ParentID: types.RowIDPointer(1),
			Code:     helper.StrPointer("311"),
			Name:     "HQ",
			Type:     accounttype.Expense,
		},
	}

	for _, v := range accounts {
		fmt.Println(v.Name, v.ID)
	}

	root := treeChartOfAccounts(accounts)

	b, err := json.MarshalIndent(root, "", "    ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))

}
