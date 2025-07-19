//
// Created by nekoimi on 2021/3/24.
//
#include <stdio.h>
#include "ci_quick_sort.h"

void print_arr(int nums[], int numsSize) {
    printf("Array: [ ");
    for (int i = 0; i < numsSize; ++i) {
        if (i + 1 == numsSize) {
            printf("%d ", nums[i]);
        } else {
            printf("%d, ", nums[i]);
        }
    }
    printf("] \n");
}

void quick_sort(int nums[], int start, int end) {
    int i, j, stand;
    if (start > end) {
        return;
    }

    // 默认取第一个值为标准
    stand = nums[start];
    i = start;
    j = end;
    // 循环查找 i -> , <- j 从两个方向开始查找
    while (i != j) {

        // fixme 先要从右边开始找
        while (i < j && nums[j] <= stand) {
            printf("j--, %d \n", j);
            j--;
        }

        while (i < j && nums[i] >= stand) {
            printf("i++, %d \n", i);
            i++;
        }

        if (i < j) {
            printf("交换i, j的值, nums[i] -> %d, nums[j] -> %d\n", nums[i], nums[j]);
            nums[i] = nums[i] ^ nums[j];
            nums[j] = nums[i] ^ nums[j];
            nums[i] = nums[i] ^ nums[j];
        }

    }

    nums[start] = nums[i];
    nums[i] = stand;

    printf("next \n");
    quick_sort(nums, start, i - 1);
    quick_sort(nums, i + 1, end);
}

void quick_do() {
    int nums[] = {3, 4, 1, 9, 6, 2, 10, 100, 24, 34, 24, 24, 24, 32, 58, 99, 4, 71, 54};
    int length = sizeof(nums) / sizeof(int);
    quick_sort(nums, 0, length - 1);
    print_arr(nums, length);
}

