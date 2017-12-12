namespace cpp gateway

enum ResultCode {
    Okay = 0,
    Err_Start = 1,
    Err_MaxConn = 2
}

struct Result {
    1:ResultCode ret
    2:i32 conn
}


service Gateway {
    Result start()
}