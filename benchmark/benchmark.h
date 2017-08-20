#pragma once
#include <stdio.h>
#include <time.h>
#define LINUX_OS 1
#define WINDOWS_OS 2

void write_results(clock_t what)
{
    printf("... # %.3fs\n", (float) what/CLOCKS_PER_SEC);
}

#ifdef _WIN32

int get_os() {
  return WINDOWS_OS;
}

#else

int get_os() {
  return LINUX_OS;
}

#endif
