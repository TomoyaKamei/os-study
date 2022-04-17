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