#include<iostream>
#include<vector>
using namespace std;

int main(){
    // // 空のベクタ生成
    // vector<int> arr;       
    
    // // 要素を含むベクタの初期化
    // vector<int> arr = {1, 2, 12, 15, 17};

    // FillConstructorはベクトル名(要素数、要素の値)で、要素数分だけ要素の値が詰まったベクトルを初期化する。
    vector<int> arr(10, 7);

    // // ベクトルの末尾の要素削除
    // arr.pop_back();

    // ベクトルの末尾に要素追加
    arr.push_back(23);

    // iostreamを用いて値を出力するには、cout<<値<<値2 ... を用いる。
    // ベクトルの要素数を出力
    cout << arr.size() <<endl;

    // ベクトルの容量を出力
    // ベクトルの容量は可変であり、初期化した要素数が最初の容量数となる。
    cout << arr.capacity() << endl;

    // ベクトルの全ての要素の出力
    for (int i=0; i<arr.size(); i++){
        cout << arr[i] << endl;
    }

    return 0;    
}