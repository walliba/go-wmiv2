package mi

type Result uint32

const (
	RESULT_OK                                  Result = iota // 0
	RESULT_FAILED                                            // 1
	RESULT_ACCESS_DENIED                                     // 2
	RESULT_INVALID_NAMESPACE                                 // 3
	RESULT_INVALID_PARAMETER                                 // 4
	RESULT_INVALID_CLASS                                     // 5
	RESULT_NOT_FOUND                                         // 6
	RESULT_NOT_SUPPORTED                                     // 7
	RESULT_CLASS_HAS_CHILDREN                                // 8
	RESULT_CLASS_HAS_INSTANCES                               // 9
	RESULT_INVALID_SUPERCLASS                                // 10
	RESULT_ALREADY_EXISTS                                    // 11
	RESULT_NO_SUCH_PROPERTY                                  // 12
	RESULT_TYPE_MISMATCH                                     // 13
	RESULT_QUERY_LANGUAGE_NOT_SUPPORTED                      // 14
	RESULT_INVALID_QUERY                                     // 15
	RESULT_METHOD_NOT_AVAILABLE                              // 16
	RESULT_METHOD_NOT_FOUND                                  // 17
	RESULT_NAMESPACE_NOT_EMPTY                               // 20
	RESULT_INVALID_ENUMERATION_CONTEXT                       // 21
	RESULT_INVALID_OPERATION_TIMEOUT                         // 22
	RESULT_PULL_HAS_BEEN_ABANDONED                           // 23
	RESULT_PULL_CANNOT_BE_ABANDONED                          // 24
	RESULT_FILTERED_ENUMERATION_NOT_SUPPORTED                // 25
	RESULT_CONTINUATION_ON_ERROR_NOT_SUPPORTED               // 26
	RESULT_SERVER_LIMITS_EXCEEDED                            // 27
	RESULT_SERVER_IS_SHUTTING_DOWN                           // 28
)

type Type uint32

const (
	MI_BOOLEAN    Type = iota
	MI_UINT8                // 1
	MI_SINT8                // 2
	MI_UINT16               // 3
	MI_SINT16               // 4
	MI_UINT32               // 5
	MI_SINT32               // 6
	MI_UINT64               // 7
	MI_SINT64               // 8
	MI_REAL32               // 9
	MI_REAL64               // 10
	MI_CHAR16               // 11
	MI_DATETIME             // 12
	MI_STRING               // 13
	MI_REFERENCE            // 14
	MI_INSTANCE             // 15
	MI_BOOLEANA             // 16
	MI_UINT8A               // 17
	MI_SINT8A               // 18
	MI_UINT16A              // 19
	MI_SINT16A              // 20
	MI_UINT32A              // 21
	MI_SINT32A              // 22
	MI_UINT64A              // 23
	MI_SINT64A              // 24
	MI_REAL32A              // 25
	MI_REAL64A              // 26
	MI_CHAR16A              // 27
	MI_DATETIMEA            // 28
	MI_STRINGA              // 29
	MI_REFERENCEA           // 30
	MI_INSTANCEA            // 31
	MI_ARRAY      Type = 16 // 16
)

// type MI_ENUMERATION

// type MI_ClassDecl struct {
// 	flags uint32
// 	code  uint32
// 	name  *uint16
// }
