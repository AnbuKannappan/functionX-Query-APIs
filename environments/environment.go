package environments

import (
	"github.com/Netflix/go-env"
)

type FXCORE_ENV struct {
	App_Port string `env:"APP_PORT,default=:5000"`
	Extras   env.EnvSet
}

func (fxenv *FXCORE_ENV) Load() error {
	es, err := env.UnmarshalFromEnviron(fxenv)
	if err != nil {
		return err
	}
	fxenv.Extras = es

	return nil
}
