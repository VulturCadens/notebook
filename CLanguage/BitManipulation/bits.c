#include <stdbool.h>
#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>

#define B_0000_0000 0x00
#define B_0000_0001 0x01
#define B_0000_0010 0x02
#define B_0000_0100 0x04
#define B_0000_1000 0x08
#define B_0001_0000 0x10
#define B_0010_0000 0x20
#define B_0100_0000 0x40
#define B_1000_0000 0x80

uint8_t setBit(uint8_t b, uint8_t m)
{
    b |= m;
    return b;
}

uint8_t clearBit(uint8_t b, uint8_t m)
{
    b &= ~m;
    return b;
}

bool testBit(uint8_t b, uint8_t m)
{
    return (b & m) != 0;
}

const char* representation[] = {
    [0x00] = "0000",
    [0x01] = "0001",
    [0x02] = "0010",
    [0x03] = "0011",
    [0x04] = "0100",
    [0x05] = "0101",
    [0x06] = "0110",
    [0x07] = "0111",
    [0x08] = "1000",
    [0x09] = "1001",
    [0x0A] = "1010",
    [0x0B] = "1011",
    [0x0C] = "1100",
    [0x0D] = "1101",
    [0x0E] = "1110",
    [0x0F] = "1111",
};

char* get_binary_string(uint8_t byte)
{
    char* s = (char*)calloc(10, sizeof(char));

    if (s == NULL) {
        fprintf(stderr, "The memory cannot be allocated\n");
        exit(EXIT_FAILURE);
    }

    sprintf(s, "%s %s", representation[byte >> 4], representation[byte & 0x0F]);

    return s;
}

int main(void)
{
    char* str;
    uint8_t b = B_0000_0000;

    b = setBit(b, B_0000_0001 | B_0000_0100);

    str = get_binary_string(b);
    printf("Set first and third bit -> byte is %s\n", str);
    free(str);

    if (testBit(b, B_0000_0100))
        printf("Test third bit -> True\n");
    else
        printf("Test third bit -> False\n");

    b = clearBit(b, B_0000_0100);

    str = get_binary_string(b);
    printf("Clear third bit -> byte is %s\n", str);
    free(str);

    if (testBit(b, B_0000_0100))
        printf("Test third bit -> True\n");
    else
        printf("Test third bit -> False\n");

    return EXIT_SUCCESS;
}
