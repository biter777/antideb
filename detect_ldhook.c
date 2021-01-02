#include <stdio.h>
#include <sys/auxv.h>
#include <limits.h>
#include <stdlib.h>
#include <string.h>

#include "antideb.h"

extern long _r_debug;

static const unsigned char *arch_ret_ldhook[] = {
    [ARCH_AMD64] = (unsigned char *)"\xf3\xc3",         /* rep ret */
    [ARCH_I386] = (unsigned char *)"\xf3\xc3",          /* rep ret */
    [ARCH_ARM64] = (unsigned char *)"\xc0\x03\x5f\xd6", /* ret     */
    [ARCH_ARMV7] = (unsigned char *)"\x1e\xff\x2f\xe1", /* bx lr   */
};

static const size_t arch_ret_len_ldhook[] = {
    [ARCH_AMD64] = 2,
    [ARCH_I386] = 2,
    [ARCH_ARM64] = 4,
    [ARCH_ARMV7] = 4,
};

static unsigned int get_arch_copy(void)
{
    const char *arch_strings[] = {
        [ARCH_AMD64] = "x86_64",
        [ARCH_I386] = "i686",
        [ARCH_ARM64] = "aarch64",
        [ARCH_ARMV7] = "v7l",
    };

    for (unsigned int i = 0; i < sizeof(arch_strings) / sizeof(arch_strings[0]); i++)
    {
        if (!strcmp((const char *)getauxval(AT_PLATFORM), arch_strings[i]))
            return i;
    }

    return UINT_MAX;
}

static int detect_ldhook(void)
{
    unsigned int this_arch = get_arch_copy();

    if (memcmp((void *)*(&_r_debug + 2), arch_ret_ldhook[this_arch],
               arch_ret_len_ldhook[this_arch]))
        return RESULT_YES;
    else
        return RESULT_NO;
}
