/*
https://atcoder.jp/contests/abc418/tasks/abc418_c
*/

using namespace std;

#include <iostream> // For input/output operations
#include <vector>   // For dynamic arrays
#include <algorithm> // For various algorithms like sort, min, max
#include <string>    // For string manipulation
#include <cmath>     // For mathematical functions

// Optional: Include other commonly used headers
// #include <map>
// #include <set>
// #include <queue>

// Define common macros for convenience
#define ll long long

// Fast I/O
void fast_io() {
    std::ios_base::sync_with_stdio(false);
    std::cin.tie(NULL);
    std::cout.tie(NULL);
}

int main() {
    fast_io(); // Enable fast I/O

    // Step 1: Read in Input!!
    // n = number of types of tea
    // q = number of queries
    int n, q;
    cin >> n >> q;
    vector<int> teaAmts(n); // amounts of each type of tea
    vector<int> bVals(q); // game difficulties, b, for each query

    for (int i = 0; i < n; i++) {
        cin >> teaAmts[i];
    }
    for (int i = 0; i < q; i++) {
        cin >> bVals[i];
    }

    // Step 2:
    /*
    From the dealer perspective, it is optimal
    to inflate player's hand with teabags of types that have
    less amount in existence because they cannot help you win.
    So sort teabags by amount, ascending
    */
    sort(teaAmts.begin(), teaAmts.end());

    // Step 3:
    /*
    Create a vector that stores sum of elements in teaAmts
    sums[i] = sum of elements 0 to i in teaAmts
    */
    vector<ll> sums(n);
    sums[0] = teaAmts[0];
    for (int i = 1; i < n; i++) {
        sums[i] = sums[i-1] + teaAmts[i];
    }

    for (int query = 0; query < q; query++) {
        int b = bVals[query];
        // Step 4:
        /*
        Use binary search to find the first tea type with at least b amount
        */
        auto iter = lower_bound(teaAmts.begin(), teaAmts.end(), b);
        int index = iter - teaAmts.begin();
        // You cannot win if none of the teas have at least b amount
        if (index >= n) {
            cout << -1 << "\n";
            continue;
        }

        // index contains the first tea type you can win with
        /* x = total number tea with type amount less than b
            + b-1*(tea types with amount at least b)
            + 1
        */
        ll x = sums[index] - teaAmts[index];
        x += ((ll)b-1)*((ll)n - (ll)index);
        x += 1;
        cout << x << "\n";
    }
    return 0;
}