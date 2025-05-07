package mi

type Flag uint32

const (
	FLAG_CLASS           Flag = 1 << 0                                                                                                          // CIM meta type (or qualifier scope)
	FLAG_METHOD          Flag = 1 << 1                                                                                                          // CIM meta type (or qualifier scope)
	FLAG_PROPERTY        Flag = 1 << 2                                                                                                          // CIM meta type (or qualifier scope)
	FLAG_PARAMETER       Flag = 1 << 3                                                                                                          // CIM meta type (or qualifier scope)
	FLAG_ASSOCIATION     Flag = 1 << 4                                                                                                          // CIM meta type (or qualifier scope)
	FLAG_INDICATION      Flag = 1 << 5                                                                                                          // CIM meta type (or qualifier scope)
	FLAG_REFERENCE       Flag = 1 << 6                                                                                                          // CIM meta type (or qualifier scope)
	FLAG_ANY             Flag = FLAG_CLASS | FLAG_METHOD | FLAG_PROPERTY | FLAG_PARAMETER | FLAG_ASSOCIATION | FLAG_INDICATION | FLAG_REFERENCE //
	FLAG_ENABLEOVERRIDE  Flag = 1 << 7                                                                                                          // qualifier flavors
	FLAG_DISABLEOVERRIDE Flag = 1 << 8                                                                                                          // qualifier flavors
	FLAG_RESTRICTED      Flag = 1 << 9                                                                                                          // qualifier flavors
	FLAG_TOSUBCLASS      Flag = 1 << 10                                                                                                         // qualifier flavors
	FLAG_TRANSLATABLE    Flag = 1 << 11                                                                                                         // qualifier flavors
	FLAG_KEY             Flag = 1 << 12                                                                                                         // property qualifier
	FLAG_IN              Flag = 1 << 13                                                                                                         // property qualifier
	FLAG_OUT             Flag = 1 << 14                                                                                                         // property qualifier
	FLAG_REQUIRED        Flag = 1 << 15                                                                                                         // property qualifier
	FLAG_STATIC          Flag = 1 << 16                                                                                                         //
	FLAG_ABSTRACT        Flag = 1 << 17                                                                                                         //
	FLAG_TERMINAL        Flag = 1 << 18                                                                                                         //
	FLAG_EXPENSIVE       Flag = 1 << 19                                                                                                         //
	FLAG_STREAM          Flag = 1 << 20                                                                                                         //
	FLAG_READONLY        Flag = 1 << 21                                                                                                         //
	FLAG_EXTENDED        Flag = 1 << 12                                                                                                         // Special flag; same as FLAG_KEY
	FLAG_NOT_MODIFIED    Flag = 1 << 25                                                                                                         // indicates that the property is not modified
	FLAG_VERSION         Flag = (1 << 26) | (1 << 27) | (1 << 28)                                                                               //
	FLAG_NULL            Flag = 1 << 29                                                                                                         // null value
	FLAG_BORROW          Flag = 1 << 30                                                                                                         // memory management
	FLAG_ADOPT           Flag = (1 << 31)                                                                                                       // memory management
)

var miFlags = map[Flag]string{
	FLAG_CLASS:           "CLASS",
	FLAG_METHOD:          "METHOD",
	FLAG_PROPERTY:        "PROPERTY",
	FLAG_PARAMETER:       "PARAMETER",
	FLAG_ASSOCIATION:     "ASSOCIATION",
	FLAG_INDICATION:      "INDICATION",
	FLAG_REFERENCE:       "REFERENCE",
	FLAG_ENABLEOVERRIDE:  "ENABLEOVERRIDE",
	FLAG_DISABLEOVERRIDE: "DISABLEOVERRIDE",
	FLAG_RESTRICTED:      "RESTRICTED",
	FLAG_TOSUBCLASS:      "TOSUBCLASS",
	FLAG_TRANSLATABLE:    "TRANSLATABLE",
	FLAG_KEY:             "KEY",
	FLAG_IN:              "IN",
	FLAG_OUT:             "OUT",
	FLAG_REQUIRED:        "REQUIRED",
	FLAG_STATIC:          "STATIC",
	FLAG_ABSTRACT:        "ABSTRACT",
	FLAG_TERMINAL:        "TERMINAL",
	FLAG_EXPENSIVE:       "EXPENSIVE",
	FLAG_STREAM:          "STREAM",
	FLAG_READONLY:        "READONLY",
	FLAG_NOT_MODIFIED:    "NOT_MODIFIED",
	FLAG_NULL:            "NULL",
	FLAG_BORROW:          "BORROW",
	FLAG_ADOPT:           "ADOPT",
}

func (f Flag) HasFlag(flag Flag) bool {
	return f&flag != 0
}

func (f Flag) GetFlags() []string {
	var result []string
	for bit, name := range miFlags {
		if f&bit != 0 {
			result = append(result, name)
		}
	}
	return result
}
