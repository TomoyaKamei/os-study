#include<iostream>
#include<vector>
using namespace std;

int mountain_MyAns(vector<int> arr){
    int result = 0;
    int mountainLenTmp = 1;
    bool ascFlag = true;

    for (int i=0; i < arr.size()-1; i++){
        int j = arr[i] - arr[i+1];

        if (ascFlag == true){
            mountainLenTmp += 1;

            if(j >= 0){
                ascFlag = false;
            }
        }else if (ascFlag == false){
            mountainLenTmp += 1;

            if (j < 0){
                if (result < mountainLenTmp){
                    result = mountainLenTmp;
                }
                cout << i << " " << mountainLenTmp << endl;
                mountainLenTmp = 1;
                ascFlag = true;
            }
        }
    }

    if (result < mountainLenTmp){
        result = mountainLenTmp;
    }

    return result;
}

int triplets_Ans(vector<int> arr){
    int n = arr.size();

    int largest = 0;

    for (int i=0;i<=n-2;){
        if(arr[i]>arr[i-1] && arr[i]>arr[i+1]){
            int cnt = 1;
            int j = i;

            while(j>=1 && arr[j]>arr[j-1]){
                j--;
                cnt++;
            }

            while(i<=1 && arr[i]>arr[i+1]){
                i++;
                cnt++;
            }

            if (cnt > largest){
                cout << i << " " << cnt << endl;
                largest = cnt;
            }
        }else{
            i++;
        }
    }
    return largest;
}

int main(){
vector<int> arr{3, -5, 5, 6, 1, 2, 3, 4, 5, 5, 3, 2, 0, 1, 2, 3, -2, 4};

    int p = mountain_MyAns(arr);
    cout << "result: " << p << endl;

    int q = mountain_MyAns(arr);
    cout << "result: " << q << endl;

    return 0;
}