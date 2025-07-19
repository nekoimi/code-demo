//
// Created by nekoimi on 2021/8/31.
//

#include <stdio.h>
#include <stdlib.h>
#include "island_perimeter.h"

int islandPerimeter(int **grid, int gridSize, int *gridColSize) {
    int rows = gridSize;
    int cols = gridColSize[0];
    int total_size = 0;
    for (int i = 0; i < rows; ++i) {
        for (int j = 0; j < cols; ++j) {
            int curr = grid[i][j];
            if (curr == 1) {
                total_size += 4;

                // 当前行的列没有超出
                if (j < cols - 1) {
                    int next_right = grid[i][j + 1];
                    if (next_right == 1) {
                        total_size-=2;
                    }
                }

                // 当前列的行没有超出
                if (i < rows - 1) {
                    int next_bottom = grid[i + 1][j];
                    if (next_bottom == 1) {
                        total_size-=2;
                    }
                }
            }
        }
    }
    return total_size;
}

void print_matrix(int **matrix, int rows, int cols) {
    printf("Matrix(%d, %d) => \n", rows, cols);
    for (int i = 0; i < rows; ++i) {
        for (int j = 0; j < cols; ++j) {
            printf("%d-%d => %d \n", i, j, matrix[i][j]);
        }
    }
    printf("End\n\n");
}

void do_island_perimeter() {
    int matrix[3][4] = {
            {0, 1, 0, 0},
            {1, 0, 0, 0},
            {1, 1, 0, 0}
    };


//    int a[3][4] = {
//            {0, 1, 2,  3},
//            {4, 5, 6,  7},
//            {8, 9, 10, 11}
//    };

    // int (*p)[4] = int a[4]
//    int (*p)[4] = a;
    // p + 1 => 进入二维数组下一行
    // p - 1 => 返回二维数组上一行
    // *p               => a[0][4]
    // *(p + 1)         => a[1][4]
    // *(*(p + 1) + 1)  => a[1][0]

    int *matrix_p[3];
    matrix_p[0] = matrix[0];
    matrix_p[1] = matrix[1];
    matrix_p[2] = matrix[2];
    print_matrix(matrix_p, 3, 4);
}
