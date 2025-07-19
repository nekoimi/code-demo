//
// Created by nekoimi on 2021/10/9.
//

#include <stdio.h>
#include <stdlib.h>
#include "contains_duplicate.h"

#define true 1
#define false 0

int cmp(const void* a, const void* b) {
    return (*(int*)a - *(int*)b);
}

bool containsDuplicate(int *nums, int numsSize) {
    qsort(nums, numsSize, sizeof(int), cmp);
    for (int i = 0; i < numsSize - 1; ++i) {
        if (nums[i] == nums[i + 1]) {
            return true;
        }
    }
    return false;
}

void do_contains_duplicate() {
    int numsSize = 4;
    int *array = malloc(sizeof(int) * numsSize);
    array[0] = 1;
    array[1] = 2;
    array[2] = 3;
    array[3] = 9;
    bool result = containsDuplicate(array, 4);
    printf("result: %d \n", result);
}