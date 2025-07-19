//
// Created by nekoimi on 2021/9/6.
//
#include "stdio.h"
#include "binary_search.h"

int search(int *nums, int numsSize, int target) {
    int mid = 0, mVal = 0, left = 0, right = numsSize - 1;
    while (left <= right) {
        mid = left + ((right - left) / 2);
        mVal = nums[mid];
        if (mVal == target) {
            return mid;
        }

        if (mVal < target) {
            left = mid + 1;
        } else if (mVal > target) {
            right = mid - 1;
        }
    }
    return -1;
}

void do_search() {
    int nums[] = {2, 5};

    int result = search(nums, 2, 5);

    printf("result: %d \n", result);
}
