package wmiv2

import (
	"fmt"

	"github.com/walliba/go-wmiv2/internal/mi"
)

type miApplication struct {
	raw      *mi.Application
	sessions []Session
}

func (app *miApplication) NewSession(destination string) (Session, error) {
	session, result := app.raw.NewSession(nil, nil)

	if result != mi.RESULT_OK {
		return nil, fmt.Errorf("error creating session")
	}

	return &miSession{raw: session}, nil
}

func (app *miApplication) Close() error {
	// // This feels like a band-aid solution
	// if sessionCount := len(app.sessions); sessionCount > 0 {
	// 	for i := range app.sessions {
	// 		// TODO: handle failure (err)
	// 		app.sessions[i].Close()
	// 	}

	// 	// allow slice to be garbage collected
	// 	app.sessions = nil
	// }

	result := app.raw.Close()
	if result != mi.RESULT_OK {
		return fmt.Errorf("error: application failed to close")
	}

	return nil
}

// Straight forward query with defaults (just avoids using NewSession)
func (app *miApplication) Query(query string) []Instance {
	return nil
}
