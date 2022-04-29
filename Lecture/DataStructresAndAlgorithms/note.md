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
- for文の定義の仕方として```for (型 変数名: 配列)```の方法もある。
```cpp
#include <iostream>
#include <vector>
using namespace std;

int main(){
    vector<vector<int>> arr = {
        {1, 2, 3},
        {4, 5, 6},
        {7, 8, 9, 10},
        {11, 12}
    }

    for (vector<int> vec : arr){
        for (int num; vec){
            cout << num << ",";
        }
        cout << endl;
    }
}
```
 
### Lecture10. How to submit Coding exercises?

### Lecture11. Coding Exercise1. FizzBuzz Test
```cpp
#include<vector>
#include <iostream>
#include<string>
using namespace std;

//Complete this method, don't write main
vector<string> fizzbuzz(int n){
    vector<string> arr;
    
    for (int i = 0; i <= n; i++){
        if (i % 3 == 0 && i % 5 == 0){
            arr.push_back("FizzBuzz");    
        }else if(i % 5 == 0){
            arr.push_back("Buzz");
        }else if(i % 3 == 0){
            arr.push_back("Fizz");
        }else{
            arr.push_back(std::to_string(i));
        }
    }
    return arr;
}

int main(){
    vector<string> arr = fizzbuzz(15);

    for (string s : arr){
        cout << s << endl;
    }

    return 0;
}
```

### Lecture12. Helpful Webinars

### Lecture13. Pairs
- pairSum問題は、リストarrに合計がvalueに等しい2つの異なる数値が含まれている場合、その組み合わせを返す。
```cpp
#include<iostream>
#include<vector>
using namespace std;

vector<int> pairSum_MyAns(vector<int>arr, int S){
    vector<int> result;

    for (int i; i < arr.Size(); i++){
        for (int j; j < arr.Size(); j++){
            if (i + j == S){
                result.push_back(x);
                result.push_back(arr[i]);
                return result;
            } 
        }
    }

    return {}
}

vector<int> pairSum_Ans(vector<int>arr, int S){
    unordered_set<int> s;
    vector<int> result;
    
    for (int i; i < arr.Size(); i++){
        int x = S - arr[i];

        // xが集合sの中に存在した場合
        if (s.find(x) != s.end()){
            result.push_back(x);
            result.push_back(arr[i]);
            return result;
        }
        s.insert(arr[i]);
    }
    return {};
}

int main(){
    vector<int> arr{10, 5, 2, 3, -6, 9, 11};
    int S = 4;

    auto p = pairSum_MyAns(arr, S);
    if(p.Size()==0){
        cout << "No such pair";
    }else{
        cout << "," << p[1] << endl;
    }

    return 0
}
```

### Lecture14. Pairs Code

### Lecture15. Triplets
- Triplet
    - N個の要素を含む配列と目標となる合計の整数Sが与えられる。
    - 配列から任意の要素を取り出し、合計の整数Sとなる配列を作る。
    - 回答の配列Aは、昇順に並べられる必要があり、回答がない場合は、空の配列となる。

### Lecture16. Triplets Code
```cpp
#include<iostream>
#include<vector>
using namespace std;

vector<int> triplets_MyAns(vector<int>arr, int S){
    int totalSum;
    vector<int> result;

    for (int i; i < arr.Size(); i++){
        totalSum += i
        s.insert(arr[i]);
        if (totalSum == S){
            return arr
        }else if (totalSum > S){
            int diff = totalSum - S
            vector<int> tmp = triplets_MyAns(result, diff).Size() 
            if (tmp.Size() != 0){
                result.insert(result.end(), tmp.begin(), tmp.end())
                return result
            }else{
                return {}
            }
        }
    }
}

vector<int> triplets_Ans(vector<int>arr, int S){
    vector<vector <int>> result;
    int n = arr.size();
    sort(arr.begin(), arr.end());

    for (int i=0; i<=n-3; i++){
        int j = i + 1;
        int k = n - 1;

        while(j < k){
            int currentSum = arr[i] + arr[j] + arr[k];
            
            if (currentSum == S){
                result.push_back({arr[i], arr[j], arr[k]});
                j++;
                k--;
            }else if (currentSum > S){
                k--;
            }else{
                j++;
            }
        }
    }

    return result;
}

int main(){
    vector<int> arr{2, 4, 6, 8, 13};
    int S = 15;

    auto p = triplet_MyAns(arr, S);
    if(p.Size()==0){
        cout << "No such pair";
    }else{
        cout << "," << p[1] << endl;
    }

    return 0
}
```

### Lecture17. Mountain
- Mountain
    - Mountainは、異なる整数の列を入力として、山となっている配列の長さを出力とする。

### Lecture18. Mountain Code
```cpp
#include<iostream>
#include<vector>
using namespace std;

int triplets_MyAns(vector<int>arr){
    int result = 0;
    int mountainLenTmp = 0;
    bool ascFlag = True;

    for (int i = 0; i < arr.Size()-2; i++){
        int j = arr[i] - arr[i+1]

        if (ascFlag == True){
            if (j < 0){
                moutainLenTmp += 1;
            }else if(j >= 0){
                ascFlag = False;
            }
        }else if (ascFlag == False){
            if (j > 0){
                moutainLenTmp += 1;
            }else{
                if (result < moutainLenTmp){
                    result = moutainLenTmp;
                }
                moutainLenTmp = 0;
                ascFlag = True;
            }
        }
    }

    return result;
}

vector<int> triplets_Ans(vector<int>arr, int S){
    vector<vector <int>> result;
    int n = arr.size();
    sort(arr.begin(), arr.end());

    for (int i=0; i<=n-3; i++){
        int j = i + 1;
        int k = n - 1;

        while(j < k){
            int currentSum = arr[i] + arr[j] + arr[k];
            
            if (currentSum == S){
                result.push_back({arr[i], arr[j], arr[k]});
                j++;
                k--;
            }else if (currentSum > S){
                k--;
            }else{
                j++;
            }
        }
    }

    return result;
}

int main(){
    vector<int> arr{2, 4, 6, 8, 13};
    int S = 15;

    auto p = triplet_MyAns(arr, S);
    if(p.Size()==0){
        cout << "No such pair";
    }else{
        cout << "," << p[1] << endl;
    }

    return 0
}
```

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