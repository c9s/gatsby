package sqlutils

/*
#include <stdio.h>
#include <errno.h>
#include <string.h>

*/
import "C"

func IndexOfChar(s string, sep string) int {
	return int(C.strcspn(C.CString(s), C.CString(sep)))
}
