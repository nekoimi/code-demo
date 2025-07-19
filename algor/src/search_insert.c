//
// Created by nekoimi on 2021/9/6.
//

#include "search_insert.h"

int searchInsert(int* nums, int numsSize, int target) {
    if(numsSize <= 0 || target <= nums[0]) return 0;

    int i;
    for (i = 0; i < numsSize; ++i) {
        if (target <= nums[i]) {
            return i;
        }
    }
    return i;
}

void do_search_insert() {

}
