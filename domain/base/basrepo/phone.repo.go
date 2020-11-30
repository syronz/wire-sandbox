package basrepo

import (
	"omega/domain/base/basmodel"
	"omega/domain/base/message/basterm"
	"omega/internal/core"
	"omega/internal/core/corerr"
	"omega/internal/core/corterm"
	"omega/internal/core/validator"
	"omega/internal/param"
	"omega/internal/types"
	"omega/pkg/dict"
	"omega/pkg/helper"
	"omega/pkg/limberr"
	"reflect"
)

// PhoneRepo for injecting engine
type PhoneRepo struct {
	Engine *core.Engine
	Cols   []string
}

// ProvidePhoneRepo is used in wire and initiate the Cols
func ProvidePhoneRepo(engine *core.Engine) PhoneRepo {
	return PhoneRepo{
		Engine: engine,
		Cols:   helper.TagExtracter(reflect.TypeOf(basmodel.Phone{}), basmodel.PhoneTable),
	}
}

// FindByID finds the phone via its id
func (p *PhoneRepo) FindByID(fix types.FixedNode) (phone basmodel.Phone, err error) {
	err = p.Engine.DB.Table(basmodel.PhoneTable).
		Where("id = ? AND company_id = ? AND node_id = ?", fix.ID.ToUint64(), fix.CompanyID, fix.NodeID).
		First(&phone).Error

	phone.ID = fix.ID
	err = p.dbError(err, "E1057421", phone, corterm.List)

	return
}

// FindByPhone finds the phone via its id
func (p *PhoneRepo) FindByPhone(phoneNumber string) (phone basmodel.Phone, err error) {
	err = p.Engine.DB.Table(basmodel.PhoneTable).
		Where("phone LIKE ?", phoneNumber).First(&phone).Error

	err = p.dbError(err, "E1057421", phone, corterm.List)

	return
}

// List returns an array of phones
func (p *PhoneRepo) List(params param.Param) (phones []basmodel.Phone, err error) {
	var colsStr string
	if colsStr, err = validator.CheckColumns(p.Cols, params.Select); err != nil {
		err = limberr.Take(err, "E1071147").Build()
		return
	}

	var whereStr string
	if whereStr, err = params.ParseWhere(p.Cols); err != nil {
		err = limberr.Take(err, "E1066154").Custom(corerr.ValidationFailedErr).Build()
		return
	}

	err = p.Engine.DB.Table(basmodel.PhoneTable).Select(colsStr).
		Where(whereStr).
		Order(params.Order).
		Limit(params.Limit).
		Offset(params.Offset).
		Find(&phones).Error

	err = p.dbError(err, "E1058608", basmodel.Phone{}, corterm.List)

	return
}

// Count of phones, mainly calls with List
func (p *PhoneRepo) Count(params param.Param) (count uint64, err error) {
	var whereStr string
	if whereStr, err = params.ParseWhere(p.Cols); err != nil {
		err = limberr.Take(err, "E1083854").Custom(corerr.ValidationFailedErr).Build()
		return
	}

	err = p.Engine.DB.Table(basmodel.PhoneTable).
		Where(whereStr).
		Count(&count).Error

	err = p.dbError(err, "E1099536", basmodel.Phone{}, corterm.List)
	return
}

// Save the phone, in case it is not exist create it
func (p *PhoneRepo) Save(phone basmodel.Phone) (u basmodel.Phone, err error) {
	if err = p.Engine.DB.Table(basmodel.PhoneTable).Save(&phone).Error; err != nil {
		err = p.dbError(err, "E1038506", phone, corterm.Updated)
	}

	p.Engine.DB.Table(basmodel.PhoneTable).Where("id = ?", phone.ID).Find(&u)
	return
}

// Create a phone
func (p *PhoneRepo) Create(phone basmodel.Phone) (u basmodel.Phone, err error) {
	if err = p.Engine.DB.Table(basmodel.PhoneTable).Create(&phone).Scan(&u).Error; err != nil {
		err = p.dbError(err, "E1029788", phone, corterm.Created)
	}
	return
}

// Delete the phone
func (p *PhoneRepo) Delete(phone basmodel.Phone) (err error) {
	if err = p.Engine.DB.Table(basmodel.PhoneTable).Delete(&phone).Error; err != nil {
		err = p.dbError(err, "E1099429", phone, corterm.Deleted)
	}
	return
}

// dbError is an internal method for generate proper database error
func (p *PhoneRepo) dbError(err error, code string, phone basmodel.Phone, action string) error {
	switch corerr.ClearDbErr(err) {
	case corerr.Nil:
		err = nil

	case corerr.NotFoundErr:
		err = corerr.RecordNotFoundHelper(err, code, corterm.ID, phone.ID, basterm.Phones)

	case corerr.ForeignErr:
		err = limberr.Take(err, code).
			Message(corerr.SomeVRelatedToThisVSoItIsNotV, dict.R(basterm.Users),
				dict.R(basterm.Phone), dict.R(action)).
			Custom(corerr.ForeignErr).Build()

	case corerr.DuplicateErr:
		err = limberr.Take(err, code).
			Message(corerr.VWithValueVAlreadyExist, dict.R(basterm.Phone), phone.Phone).
			Custom(corerr.DuplicateErr).Build()
		err = limberr.AddInvalidParam(err, "name", corerr.VisAlreadyExist, phone.Phone)

	case corerr.ValidationFailedErr:
		err = corerr.ValidationFailedHelper(err, code)

	default:
		err = limberr.Take(err, code).
			Message(corerr.InternalServerError).
			Custom(corerr.InternalServerErr).Build()
	}

	return err
}