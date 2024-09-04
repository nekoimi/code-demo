//
// Created by nekoimi on 2021/3/22.
//

#include <stdio.h>
#include <stdint.h>
#include "ci_number_of_bites.h"

int hammingWeight(uint32_t n) {
    int r = 0;
    for (int i = 0; i < 32; ++i) {
        if ((n & 1) == 1) {
            r++;
        }
        n >>= 1;
    }
    return r;
}

void number_of_bytes() {
    int res = hammingWeight(110);
    printf("res: %d \n",res);
}
