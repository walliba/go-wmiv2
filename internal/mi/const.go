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

// type MI_ENUMERATION

// type MI_ClassDecl struct {
// 	flags uint32
// 	code  uint32
// 	name  *uint16
// }
