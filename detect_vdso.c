#include <stdio.h>
#include <stdlib.h>
#include <sys/auxv.h>
#include <string.h>

#include "antideb.h"

static int detect_vdso(void)
{
    unsigned long tos;
    unsigned long vdso = getauxval(AT_SYSINFO_EHDR);

    if (!vdso)
    {
        /* Auxiliary vector does not contain vdso entry. Unknown result. */
        return RESULT_UNK;
    }

    if (!(aslr_active() == RESULT_YES))
    {
        /* No ASLR on this machine. Unknown result */
        return RESULT_UNK;
    }

    if ((unsigned long)&tos > vdso)
        return RESULT_YES;
    else
        return RESULT_NO;
}
