package wmiv2

import (
	"fmt"

	"github.com/walliba/go-wmiv2/internal/mi"
)

type miApplication struct {
	raw      *mi.Application
	sessions []*mi.Session
}

func (app *miApplication) NewSession(destination string) (Session, error) {
	session, result := app.raw.NewSession(nil, nil)

	if result != mi.RESULT_OK {
		return nil, fmt.Errorf("error creating session")
	}

	app.sessions = append(app.sessions, session)

	return &miSession{raw: session}, nil
}

func (app *miApplication) Close() error {
	// // This feels like a band-aid solution
	if sessionCount := len(app.sessions); sessionCount > 0 {
		for i := range app.sessions {
			// TODO: handle failure (err)
			app.sessions[i].Close()
			app.sessions[i] = nil
		}

		// allow slice to be garbage collected
		app.sessions = nil
	}

	result := app.raw.Close()
	if result != mi.RESULT_OK {
		return fmt.Errorf("error: application failed to close")
	}

	return nil
}

// Straight forward query with default session options (localhost)
func (app *miApplication) Query(namespace string, query string) []*map[string]any {

	session, _ := app.NewSession("localhost")

	defer session.Close()

	return session.Query(namespace, query)
}
