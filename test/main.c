#include <stdio.h>
// #include <stdlib.h>
// #include <string.h>

#include "cgotest.h"

static int printResult (void* data, void* ptr, int length) {
	printf("%.*s\n", length, ptr);
	return length;
}

int main() {
	char* url = "https://gist.githubusercontent.com/willie/e5b5ffdc5f106ebc2ca5/raw/0f18c778e3c4b1762421a252628f9ef096321e3c/gistfile1.txt";
	HTTPGet(url, printResult, NULL);

	return 0;
}
