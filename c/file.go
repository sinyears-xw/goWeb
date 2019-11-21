package c

/*
#include <sys/stat.h>
#include <fcntl.h>
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <error.h>

void *mmap(char *path) {
	int fd;
	struct stat sb;
	char *mp;

	if (fd = open(path, O_))

}
*/
import (
	"C"
)

func Test() {
	C.test()
}