package mi

type OperationOptions struct {
	reserved1 uint64
	reserved2 int64 // ptrdiff_t
	ft        *operationOptionsFT
}

type operationOptionsFT struct {
	delete             uintptr
	setString          uintptr
	setNumber          uintptr
	setCustomOption    uintptr
	getString          uintptr
	getNumber          uintptr
	getOptionCount     uintptr
	getOptionAt        uintptr
	getOption          uintptr
	getEnabledChannels uintptr
	clone              uintptr
	setInterval        uintptr
	getInterval        uintptr
}

func (o *OperationOptions) Delete() { panic("not implemented") }

func (o *OperationOptions) SetString() { panic("not implemented") }

func (o *OperationOptions) SetNumber() { panic("not implemented") }

func (o *OperationOptions) SetCustomOption() { panic("not implemented") }

func (o *OperationOptions) GetString() { panic("not implemented") }

func (o *OperationOptions) GetNumber() { panic("not implemented") }

func (o *OperationOptions) GetOptionCount() { panic("not implemented") }

func (o *OperationOptions) GetOptionAt() { panic("not implemented") }

func (o *OperationOptions) GetOption() { panic("not implemented") }

func (o *OperationOptions) GetEnabledChannels() { panic("not implemented") }

func (o *OperationOptions) Clone() { panic("not implemented") }

func (o *OperationOptions) SetInterval() { panic("not implemented") }

func (o *OperationOptions) GetInterval() { panic("not implemented") }

func (o *OperationOptions) SetWriteErrorMode() { panic("not implemented") }

func (o *OperationOptions) GetWriteErrorMode() { panic("not implemented") }

func (o *OperationOptions) SetPromptUserMode() { panic("not implemented") }

func (o *OperationOptions) GetPromptUserMode() { panic("not implemented") }

func (o *OperationOptions) SetPromptUserRegularMode() { panic("not implemented") }

func (o *OperationOptions) GetPromptUserRegularMode() { panic("not implemented") }

func (o *OperationOptions) SetProviderArchitecture() { panic("not implemented") }

func (o *OperationOptions) GetProviderArchitecture() { panic("not implemented") }

func (o *OperationOptions) EnableChannel() { panic("not implemented") }

func (o *OperationOptions) DisableChannel() { panic("not implemented") }

func (o *OperationOptions) SetTimeout() { panic("not implemented") }

func (o *OperationOptions) GetTimeout() { panic("not implemented") }

func (o *OperationOptions) SetResourceUriPrefix() { panic("not implemented") }

func (o *OperationOptions) GetResourceUriPrefix() { panic("not implemented") }

func (o *OperationOptions) SetResourceUri() { panic("not implemented") }

func (o *OperationOptions) GetResourceUri() { panic("not implemented") }

func (o *OperationOptions) SetUseMachineID() { panic("not implemented") }

func (o *OperationOptions) GetUseMachineID() { panic("not implemented") }
