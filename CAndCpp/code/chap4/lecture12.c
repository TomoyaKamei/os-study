#include <stdio.h>

void main(){
    int a1, a2;
    double b1, b2;

    a1 = 10;
    b1 = 5.32;
    a2 = (int)b1;
    b2 = (double)10;

    printf("a2=%d, b2=%f\n", a2, b2);
}