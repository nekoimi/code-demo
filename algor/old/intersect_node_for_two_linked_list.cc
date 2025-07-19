//
// Created by nekoimi on 2021/4/8.
//
#include <stdio.h>
#include "intersect_node_for_two_linked_list.h"


struct ListNode *getIntersectionNode(struct ListNode *headA, struct ListNode *headB) {
    struct ListNode *n1 = headA;
    struct ListNode *n2 = headB;
    while (n1 != n2) {
        n1 = n1 == NULL ? headB : n1->next;
        n2 = n2 == NULL ? headA : n2->next;
    }
    return n1;
}

struct ListNode *reverseList(struct ListNode *head) {
    struct ListNode* fast = head;
    struct ListNode* slow = NULL;
    while (fast != NULL) {
        struct ListNode* next = fast->next;
        fast->next = slow;
        slow = fast;
        fast = next;
    }
    return slow;
}

struct ListNode *reverseList2(struct ListNode *head) {
    if (head == NULL || head->next == NULL) {
        return head;
    }

    struct ListNode* res = reverseList2(head->next);
    head->next->next = head;
    head->next = NULL;
    return res;
}

void runIntersectLinkedList() {
    struct ListNode *node = newLinkedList(1);
    struct ListNode *head = node;
    node = insert(node, 2);
    node = insert(node, 3);
    node = insert(node, 4);
    insert(node, 5);

    printLinkedList(head);
    head = reverseList2(head);
    printLinkedList(head);
}
