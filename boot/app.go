package boot

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

type InitCallback func()

type StartCallback func()

type StopCallbacks func()

type appInf interface {
	Init(callbacks ...InitCallback)
	Start(callbacks ...StartCallback)
	Stop(callbacks ...StopCallbacks)
}

type app struct {
	osSigChan chan os.Signal
}

var appInstance appInf

func Application() appInf {
	if appInstance == nil {
		fmt.Println("New application instance created")
		appStruct := app{osSigChan: make(chan os.Signal, 1)}
		signal.Notify(appStruct.osSigChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
		appInstance = &appStruct
	}
	return appInstance
}

func (a *app) Init(callbacks ...InitCallback) {
	fmt.Println("Application initializing...")
	for _, callback := range callbacks {
		callback()
	}
}

func (a *app) Start(callbacks ...StartCallback) {
	fmt.Println("Application starting...")
	for _, callback := range callbacks {
		callback()
	}
}

func (a *app) Stop(callbacks ...StopCallbacks) {
	<-a.osSigChan
	fmt.Println("Application stopping...")
	for _, callback := range callbacks {
		callback()
	}
}
