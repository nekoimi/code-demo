//
// Created by nekoimi on 2021/9/1.
//

#include <stdlib.h>
#include "ls_struct.h"

void copy_ls_node(ls_node_t src, ls_node_t *target) {
    if (!src) {
        *target = NULL;
    } else {
        *target = malloc(sizeof(ls_node_t));
        // 复制 tag
        (*target)->tag = src->tag;
        // 根据 tag 类型来复制其他
        if (src->tag == 0) {
            // 原子
            (*target)->data = src->data;
        } else {
            // 子表
            // 复制表头
            copy_ls_node(src->hp, &((*target)->hp));
            // 复制表尾
            copy_ls_node(src->tp, &((*target)->tp));
        }
    }
}
