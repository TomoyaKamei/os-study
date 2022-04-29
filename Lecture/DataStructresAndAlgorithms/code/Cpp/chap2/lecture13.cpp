#include<iostream>
#include<vector>
#include<unordered_set>
using namespace std;

vector<int> pairSum_MyAns(vector<int>arr, int S){
    vector<int> result;

    for (int i=0; i < arr.size(); i++){
        for (int j=0; j < arr.size(); j++){
            if (arr[i] + arr[j] == S && i != j){
                result.push_back(arr[i]);
                result.push_back(arr[j]);
                return result;
            }
        }
    }

    return {};
}

vector<int> pairSum_Ans(vector<int>arr, int S){
    unordered_set<int> s;
    vector<int> result;
    
    for (int i = 0; i < arr.size(); i++){
        int x = S - arr[i];

        cout << x << "=" << S << "-" << arr[i] << endl;

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
    if(p.size()==0){
        cout << "No such pair";
    }else{
        cout << p[0] << "," << p[1] << endl;
    }

    return 0;
}