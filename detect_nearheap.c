#include <stdio.h>
#include <sys/auxv.h>
#include <unistd.h>
#include <stdlib.h>

#include "antideb.h"

static int detect_nearheap(void)
{
    //  GDB relocates heap to the end of the bss section
    static unsigned char bss_section;
    unsigned char *probe = malloc(0x10);

    if (probe - &bss_section > 0x20000)
    {
        return RESULT_NO;
    }
    else
    {
        return RESULT_YES;
    }
}
