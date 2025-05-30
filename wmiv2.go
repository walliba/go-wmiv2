// wmiv2 is the public-facing package that encapsulates access to the underlying MI API.
package wmiv2

import (
	"github.com/walliba/go-wmiv2/internal/mi"
)

var instance Application

// var once sync.Once

type Instance interface {
	GetProperties() any // ???
}

type Session interface {
	Close() error
	Query(namespace string, query string) []*map[string]any
	GetClass(namespace string, class string)
	GetClassNames(namespace string) []string
}

// Application is an interface describing an MI_Application instance. All interaction with MI must begin with this instance.
type Application interface {
	Close() error
	NewSession(destination string) (Session, error)
	Query(namespace string, query string) []*map[string]any
}

// GetApplication returns the MI Application instance.
//
// Per Microsoft, only one instance is recommended.
func GetApplication() Application {
	// once.Do(func() {
	// 	latentInitialize()
	// })
	if instance == nil {
		latentInitialize()
	}

	return instance
}

func latentInitialize() {

	app := new(mi.Application)

	err := mi.MI_Application_Initialize(app)

	if err != mi.RESULT_OK {
		panic("error initializing MI_Application")
	}

	if *app == (mi.Application{}) {
		panic("error: init app is nil")
	}

	instance = &miApplication{raw: app}
}
