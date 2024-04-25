// crypto.c

#include "crypto.h"
#include "crypto_rust.h"

const char* generate_key_pair() {
    char* sk;
    char* pk;
    generate_key_pair_rust(&sk, &pk);
    return sk;
}
