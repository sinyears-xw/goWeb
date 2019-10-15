package util

/*
#cgo linux LDFLAGS: -lrt
#include <fcntl.h>
#include <unistd.h>
#include <sys/mman.h>
#define FILE_MODE (S_IRUSR | S_IWUSR | S_IRGRP | S_IROTH)
int my_shm_new(char *name) {
    shm_unlink(name);
    return shm_open(name, O_RDWR|O_CREAT|O_EXCL, FILE_MODE);
}
*/