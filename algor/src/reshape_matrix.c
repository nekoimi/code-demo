//
// Created by nekoimi on 2021/9/1.
//
#include <stdlib.h>
#include "reshape_matrix.h"

int** matrixReshape(int** mat, int matSize, int* matColSize, int r, int c, int* returnSize, int** returnColumnSizes) {
    int rows = matSize;
    int cols = matColSize[0];
    if (rows * cols != r * c) {
        *returnSize = rows;
        *returnColumnSizes = matColSize;
        return mat;
    }

    *returnSize = r;
    *returnColumnSizes = malloc(sizeof(int) * r);
    int **result = malloc(sizeof(int *) * r);
    for (int i = 0; i < r; ++i) {
        (*returnColumnSizes)[i] = c;
        result[i] = malloc(sizeof(int) * c);
    }
    for (int x = 0; x < rows * cols; ++x) {
        result[x / c][x % c] = mat[x / cols][x % cols];
    }
    return result;
}
