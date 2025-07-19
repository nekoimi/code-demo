//
// Created by nekoimi on 2021/3/18.
//

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "ci_add_binary.h"

char *addBinary(const char *a, const char *b) {
    int length_a = (int )strlen(a);
    int length_b = (int )strlen(b);
    int length_tmp_str = length_a >= length_b ? length_a : length_b;

    int next_plus = 0, tmp = 0, tmp_a = 0, tmp_b = 0;

    char *tmp_str = (char *) malloc(sizeof(char) * (length_tmp_str + 2));
    memset(tmp_str, 0, length_tmp_str + 2);

    int r = 0;
    for (int a_index = length_a - 1, b_index = length_b - 1; a_index >= 0 || b_index >= 0; a_index--, b_index--) {
        if (a_index >= 0) {
            if (a[a_index] == '\0') {
                tmp_a = 0;
            } else {
                tmp_a = a[a_index] - '0';
            }
        } else {
            tmp_a = 0;
        }

        if (b_index >= 0) {
            if (b[b_index] == '\0') {
                tmp_b = 0;
            } else {
                tmp_b = b[b_index] - '0';
            }
        } else {
            tmp_b = 0;
        }

        tmp = tmp_a + tmp_b + next_plus;
        if (tmp == 2) {
            // todo 需要进1， 并且把当前位置为 0
            next_plus = 1;
            tmp = 0;
        } else {
            if (tmp > 2) {
                tmp = tmp - 2;
                next_plus = 1;
            } else {
                // todo 不需要进1, 直接相加
                next_plus = 0;
            }
        }

        tmp_str[r] = tmp + '0';
        r++;
    }


    // todo 进位
    if (next_plus > 0) {
        tmp_str[r] = 1 + '0';
    }


    char* res_str = (char* )malloc(sizeof(char) * (length_tmp_str + 2));
    memset(res_str, 0, (length_tmp_str + 2));

    // 交换
    // 011011 \0
    // \0 110110
    int set_index = 0;
    for (int i = 0; i <= length_tmp_str; i++) {
        if (tmp_str[length_tmp_str - i] == '\0') {
            continue;
        }
        res_str[set_index] = tmp_str[length_tmp_str - i];
        set_index++;
    }
    return res_str;
}

/**
执行用时：4 ms, 在所有 C 提交中击败了56.18%的用户
内存消耗：5.5 MB, 在所有 C 提交中击败了85.13%的用户
 */
void add_binary() {
    char *a = "100";
    char *b = "110010";
    char *r = addBinary(a, b);
    printf("addBin: %s", r);
}
