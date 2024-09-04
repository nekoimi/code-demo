//
// Created by nekoimi on 2021/3/31.
//
#include <stdio.h>
#include <stdlib.h>
#include "merge_two_sorted_link_list.h"

//if (l1 == NULL) {
//return l2;
//}
//
//if (l2 == NULL) {
//return l1;
//}
//
//if (l1->val < l2->val) {
//l1->next = mergeTwoLists(l1->next, l2);
//return l1;
//} else {
//l2->next = mergeTwoLists(l1, l2->next);
//return l2;
//}

// 思路：先合并两个链表公共部分，剩下哪个不为空直接接在后面
struct ListNode* mergeTwoLists(struct ListNode* l1, struct ListNode* l2) {
    struct ListNode* result = (ListNode*) malloc(sizeof(ListNode));
    struct ListNode* head = result;
    result->val = 0;
    result->next = NULL;
    while (l1->next != NULL && l2->next != NULL) {
        if (l1->val < l2->val) {
            result->next = l1->next;
            l1 = l1->next;
        } else {
            result->next = l2->next;
            l2 = l2->next;
        }
    }

    if (l1->next != NULL) {
        result->next = l1;
    }

    if (l2->next != NULL) {
        result->next = l2;
    }

    return head;
}
