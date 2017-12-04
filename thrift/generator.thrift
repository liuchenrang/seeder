namespace go generator 

enum ErrorCode {
    UNKNOWN_ERROR = 0,
    TOO_BUSY_ERROR = 1,
}

exception UserException {
    1: required ErrorCode error_code,
    2: required string error_name,
    3: optional string message,
}

exception SystemException {
    1: required ErrorCode error_code,
    2: required string error_name,
    3: optional string message,
}

exception UnknownException {
    1: required ErrorCode error_code,
    2: required string error_name,
    3: required string message,
}


struct TGetIdParams {
    1: required string tag,
    2: required i32 generator_type,
}

service IdGeneratorService {
    string ping()
        throws (1: UserException user_exception,
                2: SystemException system_exception,
                3: UnknownException unknown_exception),

    string getId(1: TGetIdParams params)
        throws (1: UserException user_exception,
                2: SystemException system_exception,
                3: UnknownException unknown_exception),

}
