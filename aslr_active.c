#include <stdio.h>

#include "antideb.h"

int aslr_active(void)
{
    char aslr_state[2] = {0};
    int res = 0;
    FILE *fp = fopen("/proc/sys/kernel/randomize_va_space", "r");

    if (!fp)
        return RESULT_UNK;

    if (fread((void *)aslr_state, 1, sizeof(aslr_state) - 1, fp) != sizeof(aslr_state) - 1)
    {
        fclose(fp);
        return RESULT_UNK;
    }

    if (aslr_state[0] != '0')
        res = RESULT_YES;
    else
        res = RESULT_NO;

    fclose(fp);
    return res;
}