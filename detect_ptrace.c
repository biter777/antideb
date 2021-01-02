#include <stdio.h>
#include <sys/ptrace.h>

#include "antideb.h"

static int detect_ptrace(void)
{
    if (ptrace(PTRACE_TRACEME, 0, NULL, NULL) == -1)
    {
        return RESULT_YES;
    }

    ptrace(PTRACE_DETACH, 0, NULL, NULL);
    return RESULT_NO;
}
