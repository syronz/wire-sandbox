// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package server

import (
	"omega/domain/base/basapi"
	"omega/domain/base/basrepo"
	"omega/domain/eaccounting/eacapi"
	"omega/domain/eaccounting/eacrepo"
	"omega/domain/service"
	"omega/internal/core"
)

// Injectors from wire.go:

func initSettingAPI(e *core.Engine) basapi.SettingAPI {
	settingRepo := basrepo.ProvideSettingRepo(e)
	basSettingServ := service.ProvideBasSettingService(settingRepo)
	settingAPI := basapi.ProvideSettingAPI(basSettingServ)
	return settingAPI
}

func initRoleAPI(e *core.Engine) basapi.RoleAPI {
	roleRepo := basrepo.ProvideRoleRepo(e)
	basRoleServ := service.ProvideBasRoleService(roleRepo)
	roleAPI := basapi.ProvideRoleAPI(basRoleServ)
	return roleAPI
}

func initUserAPI(engine *core.Engine) basapi.UserAPI {
	userRepo := basrepo.ProvideUserRepo(engine)
	basUserServ := service.ProvideBasUserService(userRepo)
	userAPI := basapi.ProvideUserAPI(basUserServ)
	return userAPI
}

func initAuthAPI(e *core.Engine) basapi.AuthAPI {
	basAuthServ := service.ProvideBasAuthService(e)
	authAPI := basapi.ProvideAuthAPI(basAuthServ)
	return authAPI
}

func initActivityAPI(engine *core.Engine) basapi.ActivityAPI {
	activityRepo := basrepo.ProvideActivityRepo(engine)
	basActivityServ := service.ProvideBasActivityService(activityRepo)
	activityAPI := basapi.ProvideActivityAPI(basActivityServ)
	return activityAPI
}

func initAccountAPI(e *core.Engine) basapi.AccountAPI {
	accountRepo := basrepo.ProvideAccountRepo(e)
	basAccountServ := service.ProvideBasAccountService(accountRepo)
	accountAPI := basapi.ProvideAccountAPI(basAccountServ)
	return accountAPI
}

func initCurrencyAPI(e *core.Engine) eacapi.CurrencyAPI {
	currencyRepo := eacrepo.ProvideCurrencyRepo(e)
	eacCurrencyServ := service.ProvideEacCurrencyService(currencyRepo)
	currencyAPI := eacapi.ProvideCurrencyAPI(eacCurrencyServ)
	return currencyAPI
}

func initTransactionAPI(e *core.Engine, slotServ service.EacSlotServ) eacapi.TransactionAPI {
	transactionRepo := eacrepo.ProvideTransactionRepo(e)
	eacTransactionServ := service.ProvideEacTransactionService(transactionRepo, slotServ)
	transactionAPI := eacapi.ProvideTransactionAPI(eacTransactionServ)
	return transactionAPI
}

func initSlotAPI(e *core.Engine, currencyServ service.EacCurrencyServ, accountServ service.BasAccountServ) eacapi.SlotAPI {
	slotRepo := eacrepo.ProvideSlotRepo(e)
	eacSlotServ := service.ProvideEacSlotService(slotRepo, currencyServ, accountServ)
	slotAPI := eacapi.ProvideSlotAPI(eacSlotServ)
	return slotAPI
}
