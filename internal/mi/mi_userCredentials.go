package mi

import (
	"syscall"

	"github.com/walliba/go-wmiv2/internal/mi/util"
	"golang.org/x/sys/windows"
)

type AuthType = *uint16

var (
	AUTH_TYPE_DEFAULT AuthType = windows.StringToUTF16Ptr("Default")
	AUTH_TYPE_NON     AuthType = windows.StringToUTF16Ptr("None")
)

type UsernamePasswordCreds struct {
	domain   *uint16
	username *uint16
	password *uint16
}

type UserCredentials struct {
	authenticationType    AuthType // *uint16
	UsernamePasswordCreds          // FIX: technically this is a union with a *uint16 alternative for certificateThumbprint
}

func NewUserCredentials(authType AuthType, domain string, username string, password string) *UserCredentials {
	w_domain, _ := syscall.UTF16PtrFromString(domain)
	w_user, _ := syscall.UTF16PtrFromString(username)
	w_pass, _ := syscall.UTF16PtrFromString(password)

	return &UserCredentials{
		authenticationType: authType,
		UsernamePasswordCreds: UsernamePasswordCreds{
			domain:   w_domain,
			username: w_user,
			password: w_pass,
		},
	}
}

func (uc *UserCredentials) Destroy() {
	util.UTF16PtrZero(uc.password)
	util.UTF16PtrZero(uc.username)
	util.UTF16PtrZero(uc.domain)
}
