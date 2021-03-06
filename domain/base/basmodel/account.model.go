package basmodel

import (
	"omega/domain/base/enum/accounttype"
	"omega/internal/core/coract"
	"omega/internal/core/corerr"
	"omega/internal/core/corterm"
	"omega/internal/types"
	"omega/pkg/dict"
	"omega/pkg/helper"
	"omega/pkg/limberr"
)

// AccountTable is used inside the repo layer
const (
	AccountTable = "bas_accounts"
)

// Account model
type Account struct {
	types.FixedNode
	ParentID  *types.RowID `json:"parent_id"`
	Code      string       `gorm:"unique" json:"code"`
	Name      string       `gorm:"not null;unique" json:"name,omitempty"`
	Type      types.Enum   `json:"type,omitempty"`
	Status    types.Enum   `json:"status,omitempty"`
	Phones    []Phone      `gorm:"-" json:"phones" table:"-"`
	Childrens []Account    `gorm:"-" json:"childrens" table:"-"`
}

// Validate check the type of fields
func (p *Account) Validate(act coract.Action) (err error) {

	// switch act {
	// case coract.Save:

	// 	if len(p.Name) < 5 {
	// 		err = limberr.AddInvalidParam(err, "name",
	// 			corerr.MinimumAcceptedCharacterForVisV,
	// 			dict.R(corterm.Name), 5)
	// 	}

	// 	if len(p.Name) > 255 {
	// 		err = limberr.AddInvalidParam(err, "name",
	// 			corerr.MaximumAcceptedCharacterForVisV,
	// 			dict.R(corterm.Name), 255)
	// 	}

	if p.Code == "" {
		err = limberr.AddInvalidParam(err, "code",
			corerr.VisRequired, dict.R(corterm.Code))
	}

	// 	if len(p.Description) > 255 {
	// 		err = limberr.AddInvalidParam(err, "description",
	// 			corerr.MaximumAcceptedCharacterForVisV,
	// 			dict.R(corterm.Description), 255)
	// 	}
	// }

	// TODO: it should be checked after API has been created
	if ok, _ := helper.Includes(accounttype.List, p.Type); !ok {
		return limberr.AddInvalidParam(err, "type",
			corerr.AcceptedValueForVareV, dict.R(corterm.Type),
			accounttype.Join())
	}

	return err
}
