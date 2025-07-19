//
// Created by nekoimi on 2021/9/18.
//
#include <stdio.h>
#include <stdlib.h>
#include "rotate_array.h"

void rotate(int* nums, int numsSize, int k) {
    if (numsSize == 0 || k == 0 || k % numsSize == 0) return;
    int* res = malloc(sizeof(int) * numsSize);
    for (int i = 0; i < numsSize; ++i) {
        res[i] = nums[i];
    }
    for (int i = 0; i < numsSize; ++i) {
        nums[(i + k) % numsSize] = res[i];
    }
}

void do_rotate() {
    int size = 7;
    int *nums = malloc(sizeof(int) * size);
    for (int i = 0; i < size; ++i) {
        nums[i] = i + 1;
    }
    int k = 3;
    rotate(nums, size, k);
    for (int i = 0; i < size; ++i) {
        printf("%d, ", nums[i]);
    }
}
