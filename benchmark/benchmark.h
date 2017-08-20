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
#include <process.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <assert.h>

int get_os() {
  return WINDOWS_OS;
}

char** str_split(char* a_str, const char a_delim)
{
    char** result    = 0;
    size_t count     = 0;
    char* tmp        = a_str;
    char* last_comma = 0;
    char delim[2];
    delim[0] = a_delim;
    delim[1] = 0;

    /* Count how many elements will be extracted. */
    while (*tmp)
    {
        if (a_delim == *tmp)
        {
            count++;
            last_comma = tmp;
        }
        tmp++;
    }

    /* Add space for trailing token. */
    count += last_comma < (a_str + strlen(a_str) - 1);

    /* Add space for terminating null string so caller
       knows where the list of returned strings ends. */
    count++;

    result = malloc(sizeof(char*) * count);

    if (result)
    {
        size_t idx  = 0;
        char* token = strtok(a_str, delim);

        while (token)
        {
            assert(idx < count);
            *(result + idx++) = strdup(token);
            token = strtok(0, delim);
        }
        assert(idx == count - 1);
        *(result + idx) = 0;
    }

    return result;
}

int sysexec(char const **argv)
{
    int return_code = (int) spawnv(P_WAIT, argv[0], argv);

    if (return_code != 0) {
        return_code = errno;
    }

    return return_code;
}

int execute(char *cmd) {
    char **parts = str_split(cmd, ' ');
    return sysexec(parts);
}

#else
#include <stdlib.h>
#include <sys/types.h>
#include <unistd.h>

int get_os() {
  return LINUX_OS;
}

int execute(char *cmd) {
  return system(cmd);
}

#endif
