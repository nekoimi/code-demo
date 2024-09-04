//
// Created by nekoimi on 2021/3/24.
//
#include <stdio.h>
#include "ci_queue.h"

// Queue
struct queue {
    int data[100];
    int first;
    int last;
};

void queue_main() {
    int i;
    struct queue q;
    q.first = 1;
    q.last = 1;
    for (int j = 0; j < 9; ++j) {
        q.data[j] = j;
        q.last++;
    }
    while (q.first < q.last) {
        printf("%d ", q.data[q.first]);
        q.first++;
    }
}
