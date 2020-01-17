package auth

import (
	"omega/engine"
	// "omega/utils/password"
)

type Service struct {
	Repo   Repo
	Engine engine.Engine
}

func ProvideService(p Repo) Service {
	return Service{Repo: p, Engine: p.Engine}
}

func (p *Service) Logout(user Auth) error {
	return p.Repo.Logout(user)
}

func (p *Service) Login(auth Auth) (Auth, error) {
	return p.Repo.Login(auth)
}
