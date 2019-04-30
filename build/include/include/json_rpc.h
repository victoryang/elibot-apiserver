#ifndef MCSERVER_JSON_RPC_H
#define MCSERVER_JSON_RPC_H

#define JRPC_PARSE_ERROR -32700
#define JRPC_INVALID_REQUEST -32600
#define JRPC_METHOD_NOT_FOUND -32601
#define JRPC_INVALID_PARAMS -32603
#define JRPC_INTERNAL_ERROR -32693

// Json RPC Context
typedef struct {
    void *data;
    int error_code;
    char *error_message;
} jrpc_context;

#endif //MCSERVER_JSON_RPC_H
