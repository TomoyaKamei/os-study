#include <stdio.h>

void main(){
    int a;
    printf("数値を入力:");
    scanf("%d", &a);

    if (a % 2 == 0){
        printf("入力した値は、偶数です。\n");
    }else if (a % 2 != 0){
        printf("入力した値は、奇数です。\n");
    }
}