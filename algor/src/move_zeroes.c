//
// Created by nekoimi on 2021/9/18.
//
#include <stdio.h>
#include <stdlib.h>
#include "move_zeroes.h"

void moveZeroes(int* nums, int numsSize) {
    if (numsSize == 0) return;
    int j = 0;
    for (int i = 0; i < numsSize; ++i) {
        if (nums[i] != 0) {
            int tmp = nums[i];
            nums[i] = nums[j];
            nums[j++] = tmp;
        }
    }
}

void do_move_zero() {
    int size = 5;
    int *nums = malloc(sizeof(int) * size);
    nums[0] = 0;
    nums[1] = 1;
    nums[2] = 5;
    nums[3] = 0;
    nums[4] = 9;
    moveZeroes(nums, size);
    for (int i = 0; i < size; ++i) {
        printf("%d, ", nums[i]);
    }
}
