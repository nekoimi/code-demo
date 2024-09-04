//
// Created by nekoimi on 2021/3/17.
//

#include <stdio.h>
#include <stdlib.h>
#include "ci_plus_one.h"

void plus_one() {
    int ints[] = {9};
    int length = sizeof(ints) / sizeof(ints[0]);

//    int *tmp = (int*) malloc(sizeof(int) * (length + 1));

    int need_plus = 0; // default do need plus.
    printf("length: %d \n", length);

    for (int i = length - 1; i >= 0; i--) {

        if (i == length - 1) {
            if (ints[i] != 9) {
                // 不是数字9， 直接加1即可
                ints[i] = ints[i] + 1;
                need_plus = 0;
            } else {
                // 数字为9
                ints[i] = 0;

                // 下一次需要进1
                need_plus = 1;
            }
        } else {
            // 需要在当前位进1
            if (need_plus) {
                if (ints[i] != 9) {
                    // 不是数字9， 直接加1即可
                    ints[i] = ints[i] + 1;
                    need_plus = 0;
                } else {
                    ints[i] = 0;
                    need_plus = 1;
                }
            }
        }

//        tmp[i] = ints[i];
        printf("i#%d -> v:%d, ", i, ints[i]);
    }

    int insert_val = 0;
    int result_length = length;
    if (ints[0] == 0) {
        // todo 需要进位
        insert_val = 1;
        result_length = length + 1;
    }

    printf("\n");

    int *results = (int*) malloc(sizeof(int) * (result_length));

    for (int k = 0; k < length; k++) {
        if (insert_val > 0) {
            results[k + 1] = ints[k];
        } else {
            results[k] = ints[k];
        }
    }

    if (insert_val > 0) {
        results[0] = insert_val;
    }

    printf("\n");
    for (int m = 0; m < result_length; m++) {
        printf("%d, ", results[m]);
    }
}
