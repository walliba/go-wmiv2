package wmiv2

import (
	"fmt"
	"runtime"

	"github.com/walliba/go-wmiv2/internal/mi"
)

type miApplication struct {
	raw *mi.Application
}

func (app *miApplication) NewSession(destination string) (Session, error) {
	session, result := app.raw.NewSession(nil, nil)

	if result != mi.RESULT_OK {
		return nil, fmt.Errorf("error creating session")
	}

	s := &miSession{raw: session}

	runtime.AddCleanup(s, func(r *mi.Session) {
		r.Close()
	}, s.raw)

	return s, nil
}

func (app *miApplication) Close() error {
	result := app.raw.Close()
	if result != mi.RESULT_OK {
		return fmt.Errorf("error: application failed to close")
	}

	instance = nil

	return nil
}

// Straight forward query with default session options (localhost)
func (app *miApplication) Query(namespace string, query string) []*map[string]any {

	session, err := app.NewSession("localhost")

	if err != nil {
		// TODO: return error type
		fmt.Println("Error on session creation")
	}

	defer session.Close()

	return session.Query(namespace, query)
}
