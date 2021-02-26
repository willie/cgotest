#include <stdio.h>

#include "cgotest.h"

static int printResult (void* data, void* ptr, int length) {
	printf("%.*s\n", length, ptr);
	return length;
}

int main() {
	char* url = "https://raw.githubusercontent.com/willie/cgotest/master/test/result.txt";
	HTTPGet(url, printResult, NULL);

	return 0;
}
