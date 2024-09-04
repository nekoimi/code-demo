//
// Created by nekoimi on 2021/3/31.
//

#ifndef XX1_COMMON_LINKED_LIST_H
#define XX1_COMMON_LINKED_LIST_H

#include <stdlib.h>

struct ListNode {
    int val;
    struct ListNode *next;
};

struct ListNode *newLinkedList(int val);
struct ListNode *insert(struct ListNode *node, int val);
struct ListNode *reversal_iteration(struct ListNode *head);
struct ListNode *reverse_one_each(struct ListNode *head);
struct ListNode* reverse_recursion(struct ListNode* node);
void printLinkedList(struct ListNode *head);


#endif //XX1_COMMON_LINKED_LIST_H
