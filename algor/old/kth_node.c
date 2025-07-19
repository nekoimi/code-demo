//
// Created by nekoimi on 2021/3/24.
//
#include <stdio.h>
#include "ci_kth_node.h"

struct KTH_ListNode {
    int val;
    struct KTH_ListNode *next;
};

int kthToLast(struct KTH_ListNode *head, int k) {
    struct KTH_ListNode *a = head;
    // a 向前走 k 步
    while (k-- > 1) {
        a = a->next;
    }

    while (a->next != NULL) {
        a = a->next;
        head = head->next;
    }

    //

    return head->val;
}


void kth_node() {

}

