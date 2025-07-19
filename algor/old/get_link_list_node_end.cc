//
// Created by nekoimi on 2021/3/31.
//
#include <stdio.h>
#include "common_linked_list.h"
#include "get_link_node_end.h"

//struct ListNode {
//    int val;
//    struct ListNode* next;
//};

struct ListNode* getKthFromEnd(struct ListNode* head, int k){
    struct ListNode* tmp = head;
    while(k -- > 1) {
        tmp = tmp->next;
    }
    while(tmp->next != NULL) {
        tmp = tmp->next;
        head = head->next;
    }
    return head;
}

void getLinkListNodeEnd () {
}
