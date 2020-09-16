package core

import (
	"fmt"

	"github.com/saurabh2013/microservice-basic/internal/config"
	"github.com/saurabh2013/microservice-basic/internal/constant"
	"github.com/saurabh2013/microservice-basic/internal/logger"
)

//Core global objects need to initalize one time.
type Core struct {
	logger *logger.Log
}

//InitService initialize all service related configurations
func InitService(env string) (*Core, error) {
	core := new(Core)
	// init config
	if err := config.Init(env); err != nil {
		return nil, fmt.Errorf("Failed to initialize config, %v", err)
	}

	// init logger
	lglv, err := config.Get(constant.LogLevel)
	if err != nil {
		return nil, err
	}

	// set log ooutput or default
	lgOut, _ := config.Get(constant.LogOutput)
	core.logger, err = logger.New(lglv, lgOut)
	if err != nil {
		return nil, err
	}

	// init db connection/pool here if any

	// init queue/sqs here if any

	return core, nil
}

//Auth

// // Setup Auth
// authority, err := config.GetEnv(config.AdfsAuthority)
// if err != nil {
// 	log.Error("AdfsAuthority error:"+err.Error(), nil)
// 	panic(err)
// }
// audience, err := config.GetEnv(config.AdfsAudience)
// if err != nil {
// 	log.Error("AdfsAudience error:"+err.Error(), nil)
// 	panic(err)
// }
// if err = auth.NewAuthenticator(authority, audience); err != nil {
// 	log.Error("Authenticator error:"+err.Error(), nil)
// 	panic(err)
// }
