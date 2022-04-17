# 一週間で身につくC/C++言語


## Section1. はじめに

### Lecture1. はじめに


## Section2. VisualStudio2022の使い方

### Lecture2. ダウンロードとインストール

### Lecture3. プロジェクトの作成

### Lecture4. プログラムの入力と実行

### Lecture5. 補足


## Section3. C言語入門1 はじめの一歩

### Lecture6. Hello, world!
```C
#include <stdio.h>

void main(){
    printf("Hello World.\n");
}
```

### Lecture7. printf関数のさまざまな使い方
```C
#include <stdio.h>

void main(){
    printf("My name is %s.\n", "Tomoya Kamei");
    printf("I'm %d years old.\n", 26);
    prinf("%f + %f = %f\n", 1.2, 1.3, 2.5);
}
```

### Lecture8. まとめ


## Section4. C言語入門2 演算と変数


### Lecture9. 演算の基本
```c
#include <stdio.h>

voin main(){
    printf("%d + %d = %d\n", 5, 2, 5 + 2);
    printf("%d - %d = %d\n", 5, 2, 5 - 2);
    printf("%d * %d = %d\n", 5, 2, 5 * 2);
    printf("%d / %d = %d\n", 5, 2, 5 / 2);
    printf("%d % %d = %d\n", 5, 2, 5 % 2);
}
```

### Lecture10. 変数の基本
```c
#include <stdio.h>

void main(){
    int a;
    int b = 3;
    int add, sub;
    double avg;

    a = 6;
    add = a + b;
    sub = a - b;
    avg = (a + b)/2.0;
    printf("%d + %d = %d\n", a, b, add);
    printf("%d - %d = %d\n", a, b, sub);
    printf("avg(%d, %d) = %f\n", a, b, avg);
}
```

### Lecture11. 代入演算子
```c
#include <stdio.h>

void main(){
    int a1 = 2, b1 = 2, c1 = 2, d1 = 2;
    int a2 = 2, b2 = 2, c2 = 2, d2 = 2;

    a1 += 1;
    b1 -= 1;
    c1 *= 2;
    d1 /= 2;

    printf("a1=%d, b1=%d, c1=%d, d1=%d\n", a1, b1, c1, d1);
    printf("a2=%d, b2=%d, c2=%d, d2=%d\n", a2, b2, c2, d2);
}
```

### Lecture12. データ型とキャスト
```c
#include <stdio.h>

void main(){
    int a1, a2;
    double b1, b2;

    a1 = 10;
    b1 = 5.32
    a2 = (int)b1;
    b2 = (double)10;

    printf("a2=%d, b2=%f\n", a2, b2);
}
```

### Lecture13. まとめ


## Section5. C言語入門3 条件の分岐

### Lecture14. if文

### Lecture15. if～else文

### Lecture16. if～else if～else

### Lecture17. 複雑な条件分岐

### Lecture18. Switch文

### Lecture19. まとめ