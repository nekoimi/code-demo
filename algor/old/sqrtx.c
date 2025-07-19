//
// Created by nekoimi on 2021/3/19.
//
#include <stdio.h>
#include "ci_sqrtx.h"

/**
 * 求解算术平方根
 *  平方根：又叫二次方根，表示为〔±√￣〕，其中属于非负数的平方根称之为算术平方根
 *         一个正数有两个实平方根，它们互为相反数，负数没有平方根，0的平方根是0。
 *  算术平方根: 一般地说，若一个非负数x的平方等于a，则x叫做a的算术平方根。
 * 思路：一个数的平方根肯定不会超过它自己，
 *      不过直觉还告诉我们，一个数的平方根最多不会超过它的一半
 *      例如 8 的平方根，8 的一半是 4，4^2=16>8
 *      2 =16>8，如果这个数越大越是如此，因此我们要计算一下，这个边界是多少。
 *          不等式求解：
 *              (a / 2) ^ 2 >= a
 */

int mySqrt (int x) {
    if (x <= 1) return x;
    int min = 0;
    int max = x;
    while (max - min > 1) {
        int m = (max + min) / 2;
        if (x / m < m) {
            max = m;
        } else {
            min = m;
        }
    }
    return min;
}

void sqrtx() {
    int r = mySqrt(2147395599);
    printf("result: %d\n", r);
}
