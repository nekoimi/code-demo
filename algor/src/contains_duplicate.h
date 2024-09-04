//
// Created by nekoimi on 2021/10/9.
//

#ifndef TEST_CONTAINS_DUPLICATE_H
#define TEST_CONTAINS_DUPLICATE_H

typedef int bool;

/**
 * 给定一个整数数组，判断是否存在重复元素。
 * 如果存在一值在数组中出现至少两次，函数返回 true 。
 * 如果数组中每个元素都不相同，则返回 false 。
 *
 * @param nums
 * @param numsSize
 * @return
 * @see {@link https://leetcode-cn.com/problems/contains-duplicate/}
 */
bool containsDuplicate(int* nums, int numsSize);

void do_contains_duplicate();

#endif //TEST_CONTAINS_DUPLICATE_H
