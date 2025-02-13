package main


/*
#include <stdint.h>

struct myStruct {
    unsigned char a;
    signed char b;
    int c;
    unsigned int d;
    char e[10];
};

// Define a function to create a new struct (makeStruct)
struct myStruct makeStruct() {
    struct myStruct s;
    s.a = 'A';
    s.b = 'Z';
    s.c = 100;
    s.d = 10;
    s.e[0] = 'h';
    s.e[1] = 'e';
    s.e[2] = 'l';
    s.e[3] = 'l';
    s.e[4] = 'o';
    return s;
}
*/
import "C"
import (
	"bytes"
	"encoding/binary"
	"log"
	"unsafe"
)

// Go equivalent struct for unpacking
type myStruct struct {
	a uint8
	b int8
	c int32
	d uint32
	e [10]uint8
}

// Unpack function to convert C struct to Go struct
func unpack(i *C.struct_myStruct) (m myStruct) {
	// Use Go's unsafe.Sizeof to get the size of the struct
	b := bytes.NewBuffer(C.GoBytes(unsafe.Pointer(i), C.int(unsafe.Sizeof(*i))))
	for _, v := range []interface{}{&m.a, &m.b, &m.c, &m.d, &m.e} {
		binary.Read(b, binary.LittleEndian, v)
	}
	return
}

func main() {
	// Create an instance of the C struct using the correct type for each field
	v := C.struct_myStruct{
		a: C.uchar('A'),    // unsigned char
		b: C.schar('Z'),    // signed char (use C.schar instead of C.char)
		c: C.int(100),
		d: C.uint(10),
		e: [10]C.char{'h', 'e', 'l', 'l', 'o'},
	}
	log.Printf("C struct: %#v", v)

	// Call the C function to create the struct
	v = C.makeStruct()
	log.Printf("C struct created by makeStruct: %#v", v)

	// Unpack the C struct into a Go struct
	m := unpack(&v)
	log.Printf("Unpacked Go struct: %#v", m)
}