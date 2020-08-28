// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package server

import (
	"omega/domain/base/basapi"
	"omega/domain/base/basrepo"
	"omega/domain/service"
	"omega/internal/core"
)

// Injectors from wire.go:

func initSettingAPI(e *core.Engine) basapi.SettingAPI {
	basSettingRepo := basrepo.ProvideSettingRepo(e)
	basBasSettingServ := service.ProvideBasSettingService(basSettingRepo)
	basSettingAPI := basapi.ProvideSettingAPI(basBasSettingServ)
	return basSettingAPI
}

func initRoleAPI(e *core.Engine) basapi.RoleAPI {
	basRoleRepo := basrepo.ProvideRoleRepo(e)
	basBasRoleServ := service.ProvideBasRoleService(basRoleRepo)
	basRoleAPI := basapi.ProvideRoleAPI(basBasRoleServ)
	return basRoleAPI
}

func initUserAPI(engine *core.Engine) basapi.UserAPI {
	basUserRepo := basrepo.ProvideUserRepo(engine)
	basBasUserServ := service.ProvideBasUserService(basUserRepo)
	basUserAPI := basapi.ProvideUserAPI(basBasUserServ)
	return basUserAPI
}

func initAuthAPI(e *core.Engine) basapi.AuthAPI {
	basBasAuthServ := service.ProvideBasAuthService(e)
	basAuthAPI := basapi.ProvideAuthAPI(basBasAuthServ)
	return basAuthAPI
}

func initActivityAPI(engine *core.Engine) basapi.ActivityAPI {
	basActivityRepo := basrepo.ProvideActivityRepo(engine)
	basBasActivityServ := service.ProvideBasActivityService(basActivityRepo)
	basActivityAPI := basapi.ProvideActivityAPI(basBasActivityServ)
	return basActivityAPI
}