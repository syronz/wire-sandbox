package service

import (
	"fmt"
	"omega/domain/base/basmodel"
	"omega/domain/base/basrepo"
	"omega/internal/core"
	"omega/internal/core/coract"
	"omega/internal/core/corerr"
	"omega/internal/param"
	"omega/internal/types"
	"omega/pkg/glog"
)

// BasPhoneServ for injecting auth basrepo
type BasPhoneServ struct {
	Repo   basrepo.PhoneRepo
	Engine *core.Engine
}

// ProvideBasPhoneService for phone is used in wire
func ProvideBasPhoneService(p basrepo.PhoneRepo) BasPhoneServ {
	return BasPhoneServ{
		Repo:   p,
		Engine: p.Engine,
	}
}

// FindByID for getting phone by it's id
func (p *BasPhoneServ) FindByID(fix types.FixedNode) (phone basmodel.Phone, err error) {
	if phone, err = p.Repo.FindByID(fix); err != nil {
		err = corerr.Tick(err, "E1057387", "can't fetch the phone", fix.ID, fix.CompanyID, fix.NodeID)
		return
	}

	return
}

// FindByPhone for getting phone by it's id
func (p *BasPhoneServ) FindByPhone(phoneNumber string) (phone basmodel.Phone, err error) {
	if phone, err = p.Repo.FindByPhone(phoneNumber); err != nil {
		err = corerr.Tick(err, "E1048291", "can't fetch the phone by phone-number", phoneNumber)
		return
	}

	return
}

// List of phones, it support pagination and search and return back count
func (p *BasPhoneServ) List(params param.Param) (phones []basmodel.Phone,
	count uint64, err error) {

	if phones, err = p.Repo.List(params); err != nil {
		glog.CheckError(err, "error in phones list")
		return
	}

	if count, err = p.Repo.Count(params); err != nil {
		glog.CheckError(err, "error in phones count")
	}

	return
}

// Create a phone
func (p *BasPhoneServ) Create(phone basmodel.Phone) (createdPhone basmodel.Phone, err error) {

	if err = phone.Validate(coract.Save); err != nil {
		err = corerr.TickValidate(err, "E1067746", "validation failed in creating the phone", phone)
		return
	}

	if createdPhone, err = p.Repo.Create(phone); err != nil {
		err = corerr.Tick(err, "E1091571", "phone not created", phone)
		return
	}

	return
}

// Save a phone, if it is exist update it, if not create it
func (p *BasPhoneServ) Save(phone basmodel.Phone) (savedPhone basmodel.Phone, err error) {
	if err = phone.Validate(coract.Save); err != nil {
		err = corerr.TickValidate(err, "E1031666", corerr.ValidationFailed, phone)
		return
	}

	if savedPhone, err = p.Repo.Save(phone); err != nil {
		err = corerr.Tick(err, "E1031295", "phone not saved")
		return
	}

	return
}

// Delete phone, it is soft delete
func (p *BasPhoneServ) Delete(fix types.FixedNode) (phone basmodel.Phone, err error) {
	if phone, err = p.FindByID(fix); err != nil {
		err = corerr.Tick(err, "E1044187", "phone not found for deleting")
		return
	}

	if err = p.Repo.Delete(phone); err != nil {
		err = corerr.Tick(err, "E1032085", "phone not deleted")
		return
	}

	return
}

// Excel is used for export excel file
func (p *BasPhoneServ) Excel(params param.Param) (phones []basmodel.Phone, err error) {
	params.Limit = p.Engine.Envs.ToUint64(core.ExcelMaxRows)
	params.Offset = 0
	params.Order = fmt.Sprintf("%v.id ASC", basmodel.PhoneTable)

	if phones, err = p.Repo.List(params); err != nil {
		err = corerr.Tick(err, "E1066621", "cant generate the excel list for phones")
		return
	}

	return
}