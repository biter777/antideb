#ifndef _GDB_DETECT_H
#define _GDB_DETECT_H

#define RESULT_NO 0
#define RESULT_UNK 1
#define RESULT_YES 2

int aslr_active(void);

#if !defined __linux__
#error "Supported only Linux"
#endif

#if !defined __amd64__ &&   \
    !defined __aarch64__ && \
    !defined __arm__ &&     \
    !defined __i386__
#error "Architecture not supported"
#endif

#define ARCH_AMD64 0
#define ARCH_I386 1
#define ARCH_ARM64 2
#define ARCH_ARMV7 3

#endif
