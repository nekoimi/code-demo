//
// Created by nekoimi on 2021/3/30.
//
#include <cstdio>
#include <cstdlib>
#include <iostream>
#include <map>
#include <vector>
#include "ci_binary_tree.h"

using namespace std;

struct BinTreeNode {
    int data;
    struct BinTreeNode *leftChild, *rightChild;
};

void putsListToBinTree(struct BinTreeNode *binTree, const int *dataList, int length) {
    int data;
    while (length > 0) {
        data = dataList[length - 1];
        BinTreeNode *tmpBt = (BinTreeNode *) malloc(sizeof(BinTreeNode));
        tmpBt->leftChild = NULL;
        tmpBt->rightChild = NULL;
        tmpBt->data = data;
        if (data >= 0) {
            binTree->rightChild = tmpBt;
        } else {
            binTree->leftChild = tmpBt;
        }
        binTree = tmpBt;
        length--;
    }
}

// 二叉树中序遍历
void middleOrderTraversal(struct BinTreeNode *binTree) {
    if (binTree == NULL) {
        return;
    }

    /**
     *                                     1
     *                      -------------------------------
     *                     2                               3
     *             ------------------              ------------------
     *            4                  5            6                  7
     */

    middleOrderTraversal(binTree->leftChild);
    printf("data: %d \n", binTree->data);
    middleOrderTraversal(binTree->rightChild);
}

// 打印二叉树
void printBinTree(struct BinTreeNode *binTree) {
    int level = 0;
    while (binTree != NULL) {
        BinTreeNode *left = binTree->leftChild;

        BinTreeNode *right = binTree->rightChild;
    }
}

class Student {
public:
    string username;
    int age;
    float score;

    void showUsername() {
        cout << "My name is " << username << endl;
    }
};

void binary_tree() {
    // pair 模板类型
    // vector 封装了动态大小数组的顺序容器，一个可以存放任意类型的动态数组
    int dataList[] = {1, -2, 3, -4, 5, -6, 7, -8, 9, 0};
    int length = sizeof(dataList) / sizeof(int);
    printf("dataLength: %d \n", length);

    BinTreeNode* binTree;
    if ((binTree = (BinTreeNode *) malloc(sizeof(BinTreeNode))) == NULL) {
        printf("malloc error! \n");
        return;
    }

    binTree->leftChild = NULL;
    binTree->rightChild = NULL;
    putsListToBinTree(binTree, dataList, length);

    middleOrderTraversal(binTree);

    free(binTree);
    binTree = NULL; // fixme 指针被free后要设置为NULL，避免野指针

    // 在 栈 (Stack) 上创建对象，有名字，返回这个对象
    Student student;
    student.username = "张三";
    student.showUsername();

    Student *student_ptr;
//    student_ptr = &student;
    // 在 堆 (Heap) 上创建对象 => 使用 new 关键字创建对象，没有名字，返回的是一个指针
    student_ptr = new Student;
    student_ptr->username = "李四";
    student_ptr->showUsername();

    delete student_ptr;
}
