//
// Created by nekoimi on 2021/3/16.
//
#include <stdio.h>
#include "ci_maximum_subarray.h"

int rmax(int a, int b) {
    if (a >= b) {
        return a;
    }
    return b;
}

// O(n^2)
// 遍历所有可能的子序列，求解所有子序列的和，取最大
void do_something(const int *sums, int length) {
    int t_val = 0;
    int t_max = sums[0];
    for (int i = 0; i < length; ++i) {
        t_val = 0;
        for (int j = i; j < length; j++) {
            t_val = t_val + sums[j];
            if (t_val > t_max) {
                t_max = t_val;
            }
        }

        printf("(%d)子序列[ ", t_val);
        for (int k = i; k < length; ++k) {
            printf("%d", sums[k]);
            if (k + 1 < length) {
                printf(", ");
            }
        }
        printf(" ]遍历\n");
    }
    printf("max val: %d\n", t_max);
}

void do_something_v2(int *nums, int size) {
    int pre = 0, max = nums[0];
    for (int i = 0; i < size; i++) {
        // 之前所有的数加起来比当前的数小， 肯定选当前的数啊， 不然就把当前的数跟以前的数加起来， 进入下一次
        if (pre + nums[i] < nums[i]) {
            // 子序和小于当前 i 数， 重新开始计算
            pre = nums[i];
        } else {
            pre = pre + nums[i];
        }
        if (pre > max) {
            max = pre;
        }
    }
    printf("max: %d\n", max);
}

// run
void maximum_subarray() {
    int nums[] = {-2, 1, -3, 4, -1, 2, 1, -5, 4};

    int length = sizeof(nums) / sizeof(nums[0]);
    do_something(nums, length);
    do_something_v2(nums, length);
}
