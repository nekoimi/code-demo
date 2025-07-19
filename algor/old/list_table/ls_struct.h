//
// Created by nekoimi on 2021/9/1.
//

#ifndef TEST_LS_STRUCT_H
#define TEST_LS_STRUCT_H

struct LsNode {
    // 节点标记：0 - 原子 / 1 - 广义表
    int tag;
    // 量表下一个节点指针
    struct LsNode *tp;
    union {
        // 节点数据
        int data;
        // 下一个子表节点的头指针
        struct LsNode *hp;
    };
};

typedef struct LsNode* ls_node_t;

/**
 * 复制广义表
 * @param src
 * @param target
 */
void copy_ls_node(ls_node_t src, ls_node_t* target);
#endif //TEST_LS_STRUCT_H
