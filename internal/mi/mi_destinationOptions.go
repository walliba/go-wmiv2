package mi

type DestinationOptions struct {
	reserved1 uint64
	reserved2 int64 // ptrdiff_t
	ft        *DestinationOptionsFT
}

type DestinationOptionsFT struct {
	Delete                   uintptr
	SetString                uintptr
	SetNumber                uintptr
	AddCredentials           uintptr
	GetString                uintptr
	GetNumber                uintptr
	GetOptionCount           uintptr
	GetOptionAt              uintptr
	GetOption                uintptr
	GetCredentialsCount      uintptr
	GetCredentialsAt         uintptr
	GetCredentialsPasswordAt uintptr
	Clone                    uintptr
	SetInterval              uintptr
	GetInterval              uintptr
}

func (d *DestinationOptions) Delete() { panic("not implemented") }

func (d *DestinationOptions) SetString() { panic("not implemented") }

func (d *DestinationOptions) SetNumber() { panic("not implemented") }

func (d *DestinationOptions) AddCredentials() { panic("not implemented") }

func (d *DestinationOptions) GetString() { panic("not implemented") }

func (d *DestinationOptions) GetNumber() { panic("not implemented") }

func (d *DestinationOptions) GetOptionCount() { panic("not implemented") }

func (d *DestinationOptions) GetOptionAt() { panic("not implemented") }

func (d *DestinationOptions) GetOption() { panic("not implemented") }

func (d *DestinationOptions) GetCredentialsCount() { panic("not implemented") }

func (d *DestinationOptions) GetCredentialsAt() { panic("not implemented") }

func (d *DestinationOptions) GetCredentialsPasswordAt() { panic("not implemented") }

func (d *DestinationOptions) Clone() { panic("not implemented") }

func (d *DestinationOptions) SetInterval() { panic("not implemented") }

func (d *DestinationOptions) GetInterval() { panic("not implemented") }
