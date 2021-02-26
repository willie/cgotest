package main

// typedef int (*writeFunc) (void* writeFuncData, void* ptr, int length);
// int bridge_write_func(writeFunc fn, void* writeFuncData, void* ptr, int length) {
//	return fn(writeFuncData, ptr, length);
// }
import "C"
