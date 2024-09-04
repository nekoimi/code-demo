//
// Created by nekoimi on 2021/4/8.
//
#include <stdio.h>
#include "delete_current_node.h"

void deleteNode1(struct ListNode *node) {
    while (node != NULL) {
        node->val = node->next->val;
        if (node->next->next == NULL) {
            node->next = NULL;
            break;
        }
        node = node->next;
    }
}

struct ListNode *deleteNode2(struct ListNode *head, int val) {
    struct ListNode *node = head;
    while (node != NULL) {
        if (node->val == val) {
            break;
        }
        if (node->next->next == NULL) {
            node->next = NULL;
            return head;
        }
        node = node->next;
    }
    while (node != NULL) {
        node->val = node->next->val;
        if (node->next->next == NULL) {
            node->next = NULL;
            break;
        }
        node = node->next;
    }
    return head;
}

struct ListNode* deleteNode3(struct ListNode* head, int val) {
    struct ListNode* p;
    struct ListNode* q;
    p = head; // 慢指针
    q = head->next; // 快指针
    // 判断第一个节点是不是要被删除的节点
    if (p->val == val) {
        return q;
    }
    // 判断后续节点
    while (q != NULL) {
        if (q->val == val) {
            p->next = q->next;
        } else {
            p = p->next;
        }
        q = q->next;
    }
    return head;
}

struct ListNode* middleNode(struct ListNode* head){
    struct ListNode* p;
    struct ListNode* q;
    p = head;
    q = p;
    while (q && q->next) {
        p = p->next;
        q = q->next->next;
    }
    return p;
}

struct ListNode* removeElements(struct ListNode* head, int val){
    if (head == NULL) {
        return head;
    }

    struct ListNode* p = head;
    struct ListNode* q = head->next;
    while (q != NULL) {
        if (q->val == val) {
            p->next = q->next;
        } else {
            p = p->next;
        }
        q = q->next;
    }

    if (head->val == val) {
        head = head->next;
    }

    return head;
}

void runDeleteNode() {
    struct ListNode *node = newLinkedList(3);
    struct ListNode *head = node;
    node = insert(node, 3);
    node = insert(node, 3);
    node = insert(node, 3);
    insert(node, 3);

    // deleteNode1(head->next);
    // head = deleteNode2(head, 9);
    // head = deleteNode3(head, 9);

    printLinkedList(head);

    head = removeElements(head, 3);

    printLinkedList(head);
}
