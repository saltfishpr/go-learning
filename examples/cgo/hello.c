#include <stdio.h>
#include <stddef.h>
#include <stdlib.h>

long readFile(const char *filename, char **buffer)
{
    FILE *infile;
    long numbytes;

    infile = fopen(filename, "r");
    if (infile == NULL)
        return -1;
    printf("Opened file %s\n", filename);

    fseek(infile, 0L, SEEK_END);
    numbytes = ftell(infile);
    fseek(infile, 0L, SEEK_SET);

    *buffer = (char *)calloc(numbytes, sizeof(char));
    if (*buffer == NULL)
        return -1;
    printf("Allocated %ld bytes\n", numbytes);
    long length = fread(*buffer, sizeof(char), numbytes, infile);
    fclose(infile);
    return length;
}
