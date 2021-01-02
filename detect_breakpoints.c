#include <stdio.h>
#include <stdlib.h>
extern unsigned char *_start;
extern unsigned char *__etext;

static int detect_breakpoints()
{
    int count = 0;
    char *start = (char *)&_start;
    char *end = (char *)&__etext;
    // printf("_start @ %p\n", start);
    // printf("__etext @ %p\n", end);
    while (start != end)
    {
        if (((*(volatile unsigned *)start) & 0xFF) == 0xCC)
        {
            ++count;
            // printf("Breakpoint at %p: (%x)\n", start, *start);
        }
        ++start;
    }
    return count;
}