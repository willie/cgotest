package main

// typedef int (*writeFunc) (void* writeFuncData, void* ptr, int length);
// int bridge_write_func(writeFunc f, void* writeFuncData, void* ptr, int length);
import "C"
import (
	"io"
	"net/http"
	"unsafe"
)

// HTTP code

type HTTPFnWriter struct {
	Fn     C.writeFunc
	FnData unsafe.Pointer
}

func (h *HTTPFnWriter) Write(p []byte) (n int, err error) {
	l := C.int(len(p))
	n = int(C.bridge_write_func(h.Fn, h.FnData, unsafe.Pointer(&p[0]), l))
	return n, nil
}

//export HTTPGet
func HTTPGet(url *C.char, writeFunc C.writeFunc, writeFuncData unsafe.Pointer) (read int64) {
	req, _ := http.NewRequest("GET", C.GoString(url), nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}

	dest := &HTTPFnWriter{Fn: writeFunc, FnData: writeFuncData}
	read, _ = io.Copy(dest, resp.Body)

	resp.Body.Close()
	return
}

func main() {}
