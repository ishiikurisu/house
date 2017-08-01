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

int get_os() {
  return WINDOWS_OS;
}

int sysexec(char const **argv)
{
    char **arg = argv;
    return spawnv(P_WAIT, argv[1], ++arg);
}


int execute(char *cmd) {
  int return_code = -1;
  char **parts = NULL;
  int no_spaces = 0;
  int i, j, k;

  no_spaces = 0;
  for (i = 0; cmd[i] != '\0'; i++) {
    if (cmd[i] == ' ') {
      no_spaces++;
    }
  }

  // TODO Fix this monster
  parts = (char**) malloc((1+no_spaces) * sizeof(char*));
  for (i = 0, j = 0, k = 0; j <= no_spaces; i++, j++) {
    parts[j] = (char*) malloc(sizeof(char));
    for (k = 0; (cmd[i] != ' ') && (cmd[i] != '\0'); i++, k++) {
      parts[j][k] = cmd[i];
      parts[j] = (char*) realloc(parts[j], (2+k) * sizeof(char));
    }
    parts[k] = '\0';
  }

  return_code = sysexec(parts);
  return return_code;
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
