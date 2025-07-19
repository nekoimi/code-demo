//
// Created by nekoimi on 2021/4/8.
//
#include <stdio.h>
#include <math.h>
#include "covert_linked_to_number.h"

int getDecimalValue(struct ListNode *head) {
    struct ListNode* fast = head;
    struct ListNode* slow = NULL;
    while (fast != NULL) {
        struct ListNode* next = fast->next;
        fast->next = slow;
        slow = fast;
        fast = next;
    }

    int result = 0;
    int power = 0;
    while (slow != NULL) {
        result += slow->val * pow(2, power);
        slow = slow->next;
        power++;
    }
    return result;
}

int getDecimalValue2(struct ListNode *head) {
    int res = 0;
    struct ListNode* node = head;
    while (node != NULL) {
        res = ( res << 1 ) + node->val;
        node = node->next;
    }
    return res;
}

void runCovertLinkedToNumber() {
    struct ListNode *head = newLinkedList(1);
    struct ListNode *node = head;
    node = insert(node, 0);
    node = insert(node, 1);
//    insert(node, 9);
    printLinkedList(head);

    int r = getDecimalValue(head);
    printf("r: %d \n", r);
}
