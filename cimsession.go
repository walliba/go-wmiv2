package wmiv2

// type SessionOptions struct {
// 	Destination        string
// 	Credential         *mi.UserCredentials
// 	SkipTestConnection bool
// }

// CimSession is a wrapper around an MI_Session instance. This struct is very similar to the object used in PowerShell
// type CimSession struct {
// 	destination        string
// 	session            *mi.Session
// 	credential         *mi.UserCredentials
// 	skipTestConnection bool
// }

// type CimSessionOption func(*CimSession)

// func WithDestination(destination string) CimSessionOption {
// 	return func(c *CimSession) {
// 		c.destination = destination
// 	}
// }

// // TODO: Support different auth types
// func WithCredential(domain string, username string, password string) CimSessionOption {
// 	return func(c *CimSession) {
// 		c.credential = newCredentials(domain, username, password)
// 	}
// }

// // REFACTOR: this like a damn onion with these layers
// func newCredentials(domain string, username string, password string) *mi.UserCredentials {
// 	return mi.NewUserCredentials(mi.AUTH_TYPE_DEFAULT, domain, username, password)
// }

// func NewCimSession(opts ...CimSessionOption) *CimSession {
// 	// define defaults
// 	cfg := &CimSession{
// 		destination:        "localhost",
// 		credential:         nil,
// 		skipTestConnection: false,
// 	}

// 	// apply config
// 	for _, opt := range opts {
// 		opt(cfg)
// 	}

// 	var destOpt *mi.DestinationOptions = nil

// 	destOpt, err := instance.app.NewDestinationOptions()

// 	if err != mi.RESULT_OK {
// 		panic("failure creating destination options")
// 	}

// 	destOpt.AddCredentials(cfg.credential)

// 	cfg.credential.Destroy()

// 	w_dest, _ := syscall.UTF16PtrFromString(cfg.destination)

// 	if w_dest == nil {
// 		panic("failed to convert string to *uint16")
// 	}

// 	session, err := instance.app.NewSession(w_dest, destOpt)

// 	if err != mi.RESULT_OK {
// 		panic(fmt.Sprintf("error NewCimSession: HRESULT = %d", err))
// 	}

// 	cfg.session = session

// 	return cfg
// }

// func NewDefaultSession() *CimSession {}

// func NewCimSession(destination string, domain string, username string, password string) *CimSession {
// func NewCimSession(opts *SessionOptions) *CimSession {
// 	destOpt, err := instance.app.NewDestinationOptions()

// 	if err != mi.RESULT_OK {
// 		panic("failure creating destination options")
// 	}

// 	destOpt.AddCredentials(opts.Credential)

// 	opts.Credential.Destroy()

// 	w_dest := windows.StringToUTF16Ptr(opts.Destination)

// 	session, err := instance.app.NewSession(w_dest, destOpt)

// 	if err != mi.RESULT_OK {
// 		panic(fmt.Sprintf("error NewCimSession: HRESULT = %d", err))
// 	}

// 	return &CimSession{
// 		session:      session,
// 		id:           1,
// 		Name:         "Session1",
// 		ComputerName: opts.Destination,
// 		Protocol:     "Default",
// 	}
// }

// Commit finalizes the CimSession parameters and passes the options to MI for creation
// func (cs *CimSession) Commit() *CimSession {
// 	destOpt, err := instance.app.NewDestinationOptions()

// 	if err != mi.RESULT_OK {
// 		panic("failure creating destination options")
// 	}

// 	destOpt.AddCredentials(cs.Credential)

// 	cs.Credential.Destroy()

// 	w_dest := windows.StringToUTF16Ptr(cs.Destination)

// 	session, err := instance.app.NewSession(w_dest, destOpt)

// 	if err != mi.RESULT_OK {
// 		panic(fmt.Sprintf("error NewCimSession: HRESULT = %d", err))
// 	}

// 	cs.session = session

// 	return cs
// }

// func OldQuery(query string) []map[string]any {

// 	if cs.session == nil {
// 		panic("Underlying session is nil")
// 	}

// 	operation := cs.session.QueryInstances("root\\cimv2", query)

// 	defer func() {
// 		fmt.Println("attempting to close operation")
// 		if err := operation.Close(); err != mi.RESULT_OK {
// 			panic("Failed to close MI_Operation handle")
// 		}
// 	}()

// 	// TODO: use a concurrency-safe map slice or alternative
// 	result := make([]map[string]any, 0)
// 	instanceCount := 0
// 	for moreResults := true; moreResults; {

// 		instance, err := operation.GetInstance(&moreResults)

// 		if err != 0 {
// 			fmt.Println("failed on operation->GetInstance")
// 			continue
// 		}

// 		if instance != nil {
// 			instanceCount++
// 			instanceMap := make(map[string]any)
// 			result = append(result, instanceMap)

// 			var elementCount uint32

// 			err := instance.GetElementCount(&elementCount)

// 			if err != mi.RESULT_OK {
// 				fmt.Println("error getting element count")
// 			}

// 			var i uint32
// 			for i = 0; i < elementCount; i++ {
// 				// MI_Value value;
// 				value := new(mi.Value)
// 				// MI_Type type;
// 				vType := new(mi.Type)
// 				// MI_Uint32 flags;
// 				flags := new(mi.Flag)

// 				name, err := instance.GetElementAt(i, value, vType, flags)

// 				if err != mi.RESULT_OK {
// 					fmt.Fprintf(os.Stderr, "error %d: getting element at index: %d\n", err, i)
// 					continue
// 				}

// 				if flags.HasFlag(mi.FLAG_NULL) {
// 					// Omitting this results in a smaller slice, and still allows for indexing into the result map (returns nil)
// 					// key := windows.UTF16PtrToString(name)
// 					// instanceMap[key] = nil
// 					continue
// 				}

// 				key := util.UTF16PtrToString(name)
// 				instanceMap[key] = value.As(*vType)
// 			}
// 		}
// 	}

// 	return result
// }
