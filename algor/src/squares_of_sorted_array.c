//
// Created by nekoimi on 2021/9/9.
//
#include <stdio.h>
#include <stdlib.h>
#include "squares_of_sorted_array.h"

int* sortedSquares(int* nums, int numsSize, int* returnSize) {
    int *result = malloc(sizeof(int) * numsSize);
    int pos = numsSize - 1;
    int i = 0;
    int j = numsSize - 1;
    while(i <= j) {
        if(nums[i] * nums[i] > nums[j] * nums[j]) {
            result[pos] = nums[i] * nums[i];
            i++;
        } else {
            result[pos] = nums[j] * nums[j];
            j--;
        }
        pos--;
    }
    *returnSize = numsSize;
    return result;
}

void do_sorted_squares() {
    int numbers[3] = {-4, 2, 9};
    int returnSize = 0;
    int* result = sortedSquares(numbers, 3, &returnSize);
    for (int i = 0; i < 3; ++i) {
        printf("%d, ", result[i]);
    }
}
