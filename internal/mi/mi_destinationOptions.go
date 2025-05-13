package mi

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

type DestinationOption = *uint16

// NOTE: This may cause concurrency issues later on?
var (
	TIMEOUT                 DestinationOption = windows.StringToUTF16Ptr("__MI_DESTINATIONOPTIONS_TIMEOUT")
	CERT_CA_CHECK           DestinationOption = windows.StringToUTF16Ptr("__MI_DESTINATIONOPTIONS_CERT_CA_CHECK")
	CERT_CN_CHECK           DestinationOption = windows.StringToUTF16Ptr("__MI_DESTINATIONOPTIONS_CERT_CN_CHECK")
	CERT_REVOCATION_CHECK   DestinationOption = windows.StringToUTF16Ptr("__MI_DESTINATIONOPTIONS_CERT_REVOCATION_CHECK")
	PACKET_PRVIACY          DestinationOption = windows.StringToUTF16Ptr("__MI_DESTINATIONOPTIONS_PACKET_PRIVACY")
	PACKET_INTEGRITY        DestinationOption = windows.StringToUTF16Ptr("__MI_DESTINATIONOPTIONS_PACKET_INTEGRITY")
	PACKET_ENCODING         DestinationOption = windows.StringToUTF16Ptr("__MI_DESTINATIONOPTIONS_PACKET_ENCODING")
	DATA_LOCALE             DestinationOption = windows.StringToUTF16Ptr("__MI_DESTINATIONOPTIONS_DATA_LOCALE")
	UI_LOCALE               DestinationOption = windows.StringToUTF16Ptr("__MI_DESTINATIONOPTIONS_UI_LOCALE")
	MAX_ENVELOPE_SIZE       DestinationOption = windows.StringToUTF16Ptr("__MI_DESTINATIONOPTIONS_MAX_ENVELOPE_SIZE")
	ENCODE_PORT_IN_SPN      DestinationOption = windows.StringToUTF16Ptr("__MI_DESTINATIONOPTIONS_ENCODE_PORT_IN_SPN")
	HTTP_URL_PREFIX         DestinationOption = windows.StringToUTF16Ptr("__MI_DESTINATIONOPTIONS_HTTP_URL_PREFIX")
	DESTINATION_PORT        DestinationOption = windows.StringToUTF16Ptr("__MI_DESTINATIONOPTIONS_DESTINATION_PORT")
	TRANSPORT               DestinationOption = windows.StringToUTF16Ptr("__MI_DESTINATIONOPTIONS_TRANSPORT")
	PROXY_TYPE              DestinationOption = windows.StringToUTF16Ptr("__MI_DESTINATIONOPTIONS_PROXY_TYPE")
	PROXY_CREDENTIALS       DestinationOption = windows.StringToUTF16Ptr("__MI_DESTINATIONOPTIONS_PROXY_CREDENTIALS")
	DESTINATION_CREDENTIALS DestinationOption = windows.StringToUTF16Ptr("__MI_DESTINATIONOPTIONS_DESTINATION_CREDENTIALS")
	IMPERSONATION_TYPE      DestinationOption = windows.StringToUTF16Ptr("__MI_DESTINATIONOPTIONS_IMPERSONATION_TYPE")
)

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

func (do *DestinationOptions) isValid() bool {
	return do != nil && do.ft != nil
}

// Delete deletes the destination options structure created by using Application.NewDestinationOptions or DestinationOptions.Clone
func (do *DestinationOptions) Delete() {
	if do.isValid() {
		_, _, _ = syscall.SyscallN(do.ft.Delete,
			uintptr(unsafe.Pointer(do)),
		)
	}
}

func (do *DestinationOptions) SetString() { panic("not implemented") }

func (do *DestinationOptions) SetNumber() { panic("not implemented") }

func (do *DestinationOptions) AddCredentials(credentials *UserCredentials) Result {
	if !do.isValid() {
		return RESULT_INVALID_PARAMETER
	}

	r0, _, _ := syscall.SyscallN(do.ft.AddCredentials,
		uintptr(unsafe.Pointer(do)),                      // options *DestinationOptions
		uintptr(unsafe.Pointer(DESTINATION_CREDENTIALS)), // optionName *uint16
		uintptr(unsafe.Pointer(credentials)),             // credentials *UserCredentials
		0,                                                // flags uint32
	)

	return Result(r0)
}

func (do *DestinationOptions) GetString() { panic("not implemented") }

func (do *DestinationOptions) GetNumber() { panic("not implemented") }

func (do *DestinationOptions) GetOptionCount() { panic("not implemented") }

func (do *DestinationOptions) GetOptionAt() { panic("not implemented") }

func (do *DestinationOptions) GetOption() { panic("not implemented") }

func (do *DestinationOptions) GetCredentialsCount() { panic("not implemented") }

func (do *DestinationOptions) GetCredentialsAt() { panic("not implemented") }

func (do *DestinationOptions) GetCredentialsPasswordAt() { panic("not implemented") }

func (do *DestinationOptions) Clone() { panic("not implemented") }

func (do *DestinationOptions) SetInterval() { panic("not implemented") }

func (do *DestinationOptions) GetInterval() { panic("not implemented") }
