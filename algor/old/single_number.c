//
// Created by nekoimi on 2021/3/22.
//

#include <stdio.h>
#include "ci_single_number.h"

int singleNumber(int* nums, int numsSize) {
    int res = 0;
    for (int i = 0; i < numsSize; i++) {
        res ^= nums[i];
    }
    return res;
}

void single_number() {
    int sums[] = {1, 3, 3, 5, 1};
    int sum_length = sizeof(sums) / sizeof(int);
    int res = singleNumber(sums, sum_length);
    printf("res: %d \n", res);
}
