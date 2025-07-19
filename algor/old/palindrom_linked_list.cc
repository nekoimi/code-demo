//
// Created by nekoimi on 2021/4/2.
//
#include <stdio.h>
#include <vector>
#include "palindrome_linked_list.h"

bool isPalindrome(struct ListNode* head){
    std::vector<int> nodes;
    struct ListNode* tmp = head;
    while (tmp != nullptr) {
        nodes.push_back(tmp->val);
        tmp = tmp->next;
    }
    int i = nodes.size() - 1;
    while (head != nullptr && !nodes.empty()) {
        printf("head->val: %d / nodes->val: %d \n", head->val, nodes[i]);
        if (head->val != nodes[i]) {
            return false;
        }
        head = head->next;
        i--;
    }
    return true;
}

bool isPalindrome2(struct ListNode* head) {
    int nums[50001];
    int i = 0, sum = 0;
    while (head != NULL) {
        sum++;
        nums[i++] = head->val;
        head = head->next;
    }
    for (int j = 0, k = sum - 1; j < sum; ++j, k--) {
        if (nums[j] != nums[k]) {
            return false;
        }
    }
    return true;
}

//bool isPalindrome3(struct ListNode* head) {
//    struct ListNode* fast = head;
//    struct ListNode* slow = head;
//    struct ListNode* tmp = NULL;
//    struct ListNode* rev = NULL;
//    int flag = 0;
//    while (slow != NULL) {
//        tmp->next = slow;
//
//
//        slow = slow->next;
//        if (fast->next != NULL) {
//            fast = fast->next->next;
//        }
//
//        if (fast == NULL || fast->next == NULL) {
//            flag = 1;
//            fast = head;
//        }
//
//        if (flag) {
//
//        }
//    }
//}

void runIsPalindrome() {
    int nums[] = {2, 2, 1};
    int length = sizeof(nums) / sizeof(int);
    struct ListNode *node = newLinkedList(1);
    struct ListNode *head = node;
    for (int i = 0; i < length; ++i) {
        node = insert(node, nums[i]);
    }
    if (isPalindrome2(head)) {
        printf("True");
    } else {
        printf("False");
    }
}