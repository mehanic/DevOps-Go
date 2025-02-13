package main

/*
#include <stdint.h>

struct myStruct {
    unsigned char a;
    char b;
    int c;
    unsigned int d;
    char e[10];
};
*/
import "C"
import "log"

func main() {
    // Creating a C struct and initializing the fields
    v := C.struct_myStruct{
        a: C.uchar('A'),
        b: C.char('Z'),
        c: C.int(100),
        d: C.uint(10),
        e: [10]C.char{'h', 'e', 'l', 'l', 'o', ' ', ' ', ' ', ' ', ' '},
    }

    log.Printf("%#v", v)
}

//2025/02/13 15:25:39 main._Ctype_struct_myStruct{a:0x41, b:90, c:100, d:0xa, e:[10]main._Ctype_char{104, 101, 108, 108, 111, 32, 32, 32, 32, 32}, _:[2]uint8{0x0, 0x0}}
