//
// Created by nekoimi on 2021/9/1.
//

#ifndef TEST_TREE_STRUCT_H
#define TEST_TREE_STRUCT_H

#define DATA_LENGTH 1024

/**
 * 二叉树
 */
struct BinaryTreeNode {
    // 节点存储的数据
    int data;
    // 左孩子
    struct BinaryTreeNode* leftChild;
    // 右孩子
    struct BinaryTreeNode* rightChild;
};

typedef struct BinaryTreeNode binaryTreeNode;

/**
 * 保存标记的节点，入栈保存，用于二叉树后序遍历
 */
struct TagNode {
    binaryTreeNode* node;
    // 0 - 节点右边没有遍历; 1 - 左右孩子都已经遍历完全
    int tag;
};

typedef struct TagNode tagNode;

/**
 * 栈
 */
 typedef struct Stack {
     int index;
     // 这里使用数组保存, 默认长度为 1024, 不考虑溢出
     binaryTreeNode* data[DATA_LENGTH];
 } stack;

typedef struct TagStack {
    int index;
    // 这里使用数组保存，默认长度为 1024, 不考虑溢出
    tagNode data[DATA_LENGTH];
} tagStack;

/**
 * 队列
 */
typedef struct Queue {
    // 头指针
    int front;
    // 尾指针
    int rear;
    // 使用数组实现, 默认长度为 1024, 不考虑溢出
    // 因为是数组实现链式存储，当 front 和 rear 的值相等时队列为空
    binaryTreeNode* data[DATA_LENGTH];
} queue;

/**
 * 二叉树节点入队
 * @param q
 * @param node
 */
void queue_put(queue* q, binaryTreeNode* node);

/**
 * 二叉树节点出队
 * @param q
 */
binaryTreeNode* queue_get(queue* q);

/**
 * 队列是否为空
 * @param q
 * @return
 */
int queue_is_empty(queue* q);

/**
 * 元素入栈
 * @param s
 * @param tree
 */
void stack_push(stack* s, binaryTreeNode* tree);

/**
 * 元素入栈
 * @param s
 * @param tgn
 */
void stack_push_tag(tagStack * s, tagNode tgn);

/**
 * 元素出栈
 * @param s
 */
void stack_pop(stack* s);

/**
 * 元素出栈
 * @param s
 */
void stack_pop_tag(tagStack * s);

/**
 * 获取栈顶元素
 * @param s
 * @return
 */
binaryTreeNode* stack_top(stack* s);

/**
 * 获取栈顶元素
 * @param s
 * @return
 */
tagNode stack_top_tag(tagStack * s);

/**
 * 二叉树前序遍历
 *
 * 根 -> 左 -> 右
 * @param tree
 */
void each_one(struct BinaryTreeNode* tree);

/**
 * 二叉树前序遍历
 *
 * 根 -> 左 -> 右
 * @param tree
 */
void each_one_by_stack(struct BinaryTreeNode* tree);

/**
 * 二叉树中序遍历
 *
 * 左 -> 根 -> 右
 * @param tree
 */
void each_two(struct BinaryTreeNode* tree);

/**
 * 二叉树中序遍历
 *
 * 左 -> 根 -> 右
 * @param tree
 */
void each_two_by_stack(struct BinaryTreeNode* tree);

/**
 * 二叉树后序遍历
 *
 * 左 -> 右 -> 根
 * @param tree
 */
void each_three(struct BinaryTreeNode* tree);

/**
 * 二叉树后序遍历
 *
 * 左 -> 右 -> 根
 * @param tree
 */
void each_three_by_stack(struct BinaryTreeNode* tree);

/**
 * 二叉树的层序遍历
 * @param node
 */
void each_tree_by_arrangement(binaryTreeNode* node);

void binary_tree_node();

#endif //TEST_TREE_STRUCT_H
