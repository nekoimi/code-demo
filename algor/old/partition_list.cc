//
// Created by nekoimi on 2021/4/12.
//
#include <stdio.h>
#include "partition_list.h"

/**
 * 找到第一个大于 x 的节点 a, 把 x 后面所有小于 x 的节点全部挪到 a 的前面
 */
struct ListNode* partition(struct ListNode* head, int x){
    if (head == NULL || head->next == NULL) {
        return head;
    }
    struct ListNode* node = (struct ListNode*) malloc(sizeof(struct ListNode));
    node->val = -101;
    node->next = head;
    struct ListNode* p = node;
    while (p != NULL && p->next != NULL) {
        if (p->next->val >= x) {
            struct ListNode* sp = p->next;
            struct ListNode* spr = sp;
            while (sp != NULL && sp->next != NULL) {
                if (sp->next->val < x) {
                    p->next = sp->next;
                    p = p->next;
                    sp->next = sp->next->next;
                } else {
                    sp = sp->next;
                }
            }
            p->next = spr;
            break;
        }
        p = p->next;
    }
    return node->next;
}

void runPartitionList() {
    struct ListNode* head = newLinkedList(1);
    struct ListNode* node = insert(head, 4);
    node = insert(node, 3);
    node = insert(node, 0);
    node = insert(node, 2);
    node = insert(node, 5);
    node = insert(node, 2);
    printLinkedList(head);
    head = partition(head, 3);
    printLinkedList(head);
}
