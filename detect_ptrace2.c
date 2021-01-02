#include <unistd.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/ptrace.h>

static int detect_ptrace2()
{
    int ret = 0;
    asm("xorl %ebx, %ebx\n"     // %ebx = 0
        "movl %ebx, %ecx\n"     // movl %ebx, %ecx
        "movl %ebx, %edx\n"     // %edx = 0
        "incl %ecx\n"           // %ecx = 1
        "movl $0x1A, %eax\n"    // ptrace
        "int $0x80\n"           // s y s c a l l
        "movl %eax,-4(%ebp )\n" // addr ret
    );

    return ret;
}