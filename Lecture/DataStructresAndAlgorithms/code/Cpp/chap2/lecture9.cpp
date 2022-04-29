#include <iostream>
#include <vector>
using namespace std;

int main(){
    vector<vector<int>> arr = {
        {1, 2, 3},
        {4, 5, 6},
        {7, 8, 9, 10},
        {11, 12}
    };

    for (vector<int> vec : arr){
        for (int num: vec){
            cout << num << ",";
        }
        cout << endl;
    }
}