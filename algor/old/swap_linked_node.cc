//
// Created by nekoimi on 2021/4/9.
//

#include <stdio.h>
#include "swap_linked_node.h"

struct ListNode* swapPairs(struct ListNode* head){
    struct ListNode* p = head;
//    struct ListNode* na = NULL;
//    struct ListNode* nb = NULL;
//    int i = 1;
//    while (p != NULL && p->next != NULL) {
//        if (i % 2 > 0) {
//            printf("i: %d, p->val: %d \n", i, p->val);
//            na = p->next;
//            nb = p->next->next;
//            p->next = nb;
//            na->next = p;
//
//            printf("na: %d, na->next: %d \n", na->val, na->next->val);
//        } else {
//            p = p->next;
//        }
//        i ++;
//    }
    return head;
}

void runSwapLinkedNode() {
    struct ListNode *head = newLinkedList(1);
    struct ListNode *node = head;
    node = insert(node, 0);
    node = insert(node, 8);
    node = insert(node, 18);
    insert(node, 9);
    printLinkedList(head);

    head = swapPairs(head);
    printLinkedList(head);
}

