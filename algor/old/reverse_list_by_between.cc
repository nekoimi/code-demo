//
// Created by nekoimi on 2021/4/12.
//
#include <stdio.h>
#include "reverse_list_by_between.h"

struct ListNode *reverseBetween(struct ListNode *head, int left, int right) {
    if (head == NULL || head->next == NULL) {
        return head;
    }

    if (left >= right) {
        return head;
    }

    int step = 1;
    if (left == 1) {
        struct ListNode* fast = head;
        struct ListNode* slow = NULL;
        struct ListNode* spr = NULL;
        struct ListNode* next = NULL;
        while (fast != NULL && step <= right) {
            next = fast->next;
            fast->next = slow;
            slow = fast;
            if (spr == NULL) {
                spr = slow;
            }
            fast = next;
            step++;
            if (spr->next != NULL) {
                spr = spr->next;
            }
        }
        spr->next = next;
        return slow;
    }

    struct ListNode* p = head;
    while (p != NULL) {
        if (step == left - 1) {
            struct ListNode* fast = p->next;
            struct ListNode* slow = NULL;
            struct ListNode* spr = NULL;
            struct ListNode* next = NULL;
            while (fast != NULL && step < right) {
                next = fast->next;
                fast->next = slow;
                slow = fast;
                if (spr == NULL) {
                    spr = slow;
                }
                fast = next;
                step++;
                if (spr->next != NULL) {
                    spr = spr->next;
                }
            }
            spr->next = next;
            p->next = slow;
            break;
        } else {
            step++;
            p = p->next;
        }
    }
    return head;
}

void runReverseListByBetween() {
    struct ListNode *head = newLinkedList(1);
    struct ListNode *node = insert(head, 2);
    node = insert(node, 3);
    node = insert(node, 4);
    node = insert(node, 5);
    printLinkedList(head);
    head = reverseBetween(head, 1, 2);
    printLinkedList(head);
}
