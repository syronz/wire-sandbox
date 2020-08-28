package service

import (
	"fmt"
	"net/http"
	"omega/domain/base"
	"omega/domain/base/basmodel"
	"omega/domain/base/basrepo"
	"omega/internal/core"
	"omega/internal/core/action"
	"omega/internal/param"
	"omega/internal/term"
	"omega/internal/types"
	"omega/pkg/glog"
	"omega/utils/password"
)

// BasUserServ for injecting auth basrepo
type BasUserServ struct {
	Repo   basrepo.UserRepo
	Engine *core.Engine
}

// ProvideBasUserService for user is used in wire
func ProvideBasUserService(p basrepo.UserRepo) BasUserServ {
	return BasUserServ{Repo: p, Engine: p.Engine}
}

// FindByID for getting user by it's id
func (p *BasUserServ) FindByID(id types.RowID) (user basmodel.User, err error) {
	if user, err = p.Repo.FindByID(id); err != nil {
		glog.CheckError(err, fmt.Sprintf("User with id %v", id))
		return
	}

	return
}

// FindByUsername find user with username
func (p *BasUserServ) FindByUsername(username string) (user basmodel.User, err error) {
	user, err = p.Repo.FindByUsername(username)
	glog.CheckError(err, fmt.Sprintf("User with username %v", username))

	return
}

// List of users, it support pagination and search and return back count
func (p *BasUserServ) List(params param.Param) (data map[string]interface{}, err error) {

	data = make(map[string]interface{})

	data["list"], err = p.Repo.List(params)
	glog.CheckError(err, "users list")
	if err != nil {
		return
	}

	data["count"], err = p.Repo.Count(params)
	glog.CheckError(err, "users count")

	return
}

func (p *BasUserServ) Create(user basmodel.User,
	params param.Param) (createdUser basmodel.User, err error) {

	if err = user.Validate(action.Create); err != nil {
		glog.CheckError(err, term.Validation_failed)
		return
	}

	fmt.Printf(">>>>>>>>>>>>>>>>>>>>>>>>> %p \n", p.Engine)
	var oo core.Engine
	reg := p.Engine
	oo = *p.Engine
	fmt.Printf(">>>>>>>>>>>>>>>>>>>>>>>>> %p \n", &oo)
	p.Repo.Engine = &oo

	// original := p.Engine.DB
	tx := p.Engine.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			p.Engine = reg
		}
	}()
	// p.Engine.DB = tx
	oo.DB = tx

	// if createdUser, err = p.CreateRollback(user, params); err != nil {
	if createdUser, err = p.Repo.Create(user); err != nil {
		tx.Rollback()
		p.Engine = reg
		return
	}

	// time.Sleep(30 * time.Second)
	// tx.Rollback()

	tx.Commit()
	// p.Engine.DB = original
	p.Engine = reg

	return
}

func (p *BasUserServ) CreateRollback(user basmodel.User,
	params param.Param) (createdUser basmodel.User, err error) {

	if err = user.Validate(action.Create); err != nil {
		glog.CheckError(err, "Failed in validation")
		return
	}

	user.Password, err = password.Hash(user.Password, p.Engine.Envs[base.PasswordSalt])
	glog.CheckError(err, fmt.Sprintf("Hashing password failed for %+v", user))

	if createdUser, err = p.Repo.Create(user); err != nil {
		// tx.Rollback()
		p.Engine.DB.Rollback()
		p.Engine.CheckInfo(err, fmt.Sprintf("Failed in saving user for %+v", user))
	}
	// tx.Commit()
	// p.Engine.DB = original

	createdUser.Password = ""

	return
}

// Save user
func (p *BasUserServ) Save(user basmodel.User) (createdUser basmodel.User, err error) {

	var oldUser basmodel.User
	oldUser, _ = p.FindByID(user.ID)

	if user.ID > 0 {
		if err = user.Validate(action.Update); err != nil {
			return
		}

		if user.Password != "" {
			user.Password, err = password.Hash(user.Password, p.Engine.Envs[base.PasswordSalt])
			glog.CheckError(err, fmt.Sprintf("Hashing password failed for %+v", user))
		} else {
			user.Password = oldUser.Password
		}

	} else {
		if err = user.Validate(action.Create); err != nil {
			return
		}
		user.Password, err = password.Hash(user.Password, p.Engine.Envs[base.PasswordSalt])
		glog.CheckError(err, fmt.Sprintf("Hashing password failed for %+v", user))
	}

	if createdUser, err = p.Repo.Update(user); err != nil {
		p.Engine.CheckInfo(err, fmt.Sprintf("Failed in saving user for %+v", user))
	}

	BasAccessDeleteFromCache(user.ID)

	createdUser.Password = ""

	return
}

// Excel is used for export excel file
func (p *BasUserServ) Excel(params param.Param) (users []basmodel.User, err error) {
	params.Limit = p.Engine.Envs.ToUint64(core.ExcelMaxRows)
	params.Offset = 0
	params.Order = "bas_users.id ASC"

	users, err = p.Repo.List(params)
	glog.CheckError(err, "users excel")

	return
}

// Delete user, it is hard delete, by deleting account related to the user
func (p *BasUserServ) Delete(userID types.RowID, params param.Param) (user basmodel.User, err error) {
	if user, err = p.FindByID(userID); err != nil {
		return user, core.NewErrorWithStatus(err.Error(), http.StatusNotFound)
	}

	if err = p.Repo.Delete(user); err != nil {
		glog.CheckError(err, fmt.Sprintf("error in deleting user %+v", user))
	}

	return
}
