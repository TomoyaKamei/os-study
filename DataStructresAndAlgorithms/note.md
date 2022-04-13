# Data Structres & Algorithms 



## Section1. Welcome to Course


### Lecture7. Introduction
- 配列は、メモリに連続して配置された値の集合である。
    - 配列は可変長ではなく、最初に要素数を確定させておく必要がある。
```cpp
int a[100];

int a[100] = {0};
```
- ベクトルは、可変長の配列である。
    - ベクトルは規定のサイズを超えると倍のサイズを予約し、元の配列をコピーする。
    - また、ベクトルを使用するためには値による参照を実施する必要があり、ポインタにより参照を使用すると値が変更されない可能性がある。


### Lecture8. Vector STL - Demo
```cpp
#include<iostream>
#include<vector>
using namespace std;

int main(){
    // 空のベクタ生成
    vector<int> arr;       
    
    // 要素を含むベクタの初期化
    vector<int> arr = {1, 2, 12, 15, 17};

    // FillConstructorはベクトル名(要素数、要素の値)で、要素数分だけ要素の値が詰まったベクトルを初期化する。
    vector<int> arr(10, 7);

    // ベクトルの末尾の要素削除
    arr.pop_back();

    // ベクトルの末尾に要素追加
    arr.push_back(23)

    // iostreamを用いて値を出力するには、cout<<値<<値2 ... を用いる。
    // ベクトルの要素数を出力
    cout<<arr.size() <<endl;

    // ベクトルの容量を出力
    // ベクトルの容量は可変であり、初期化した要素数が最初の容量数となる。
    cout << arr.capacity() << endl;

    // ベクトルの全ての要素の出力
    for (int i=0; i<arr.size(); i++){
        cout << arr[i] << endl;
    }

    return 0;    
}

```

### Lecture9. Vector of Vector - Demo

### Lecture10. How to submit Coding exercises?

### Lecture11. Coding Exercise1. FizzBuzz Test

### Lecture12. Helpful Webinars

### Lecture13. Pairs

### Lecture14. Pairs Code

### Lecture15. Triplets

### Lecture16. Triplets Code

### Lecture17. Mountain

### Lecture18. Mountain Code

### Lecrure19. Longest Band

### Lecture20. Longest Band Code


## Section2. Arrays & Vectors
## Section3. String Problems
## Section4. Sliding Window Problems
## Section5. Sorting & Searching
## Section6. Binary Search
## Section7. Recursion Problems
## Section8. Linked Lists Problems
## Section9. Stacks & Queue Problems
## Section10. Binary Trees Problems
## Section11. BST Problems
## Section12. Priority Queue Problems
## Section13. Hashing Problems
## Section14. Tries & Pattern Matching Problems
## Section15. Graphs Problem Solving
## Section16. Dynamic Programming(1D)
## Section17. Dynamic Programming(2D)
## Section18. DSA Project - LRU Cache


## 調査
- 参照渡しとポインタ渡しの違いについて
    - 参照渡しとポインタ渡しは、根本的に変数の型情報と変数のアドレスを渡す事で引数を関数側に渡す処理を指す。
        - 根本的な違いは、参照渡しではNullポインタを渡す事が出来ない事が挙げられる。
        - Cのポインタ渡し例)  ```bool twice( int* a )```
        - Cppの参照渡し例)　```void twice( int& a )```
    - 参照
        - [C++ 値渡し、ポインタ渡し、参照渡しを使い分けよう](https://qiita.com/agate-pris/items/05948b7d33f3e88b8967)


## 参考サイト
    - [江添亮のC++入門](https://cpp.rainy.me/021-three-virtues-of-a-programmer.html)