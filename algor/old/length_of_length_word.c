//
// Created by nekoimi on 2021/3/17.
//

#include <stdio.h>

void length_of_last_word() {
    char *s = "b   a    ";
    int i = 0, pre_length = 0, length = 0;
    while (1) {
        if (s[i] == '\0') {
            if (length == 0) {
                length = pre_length;
            }
            break;
        } else {
            printf("%c, ", s[i]);
            if (s[i] == ' ') {
                if (length > 0) {
                    pre_length = length;
                }
                length = 0;
            } else {
                length ++;
            }
        }
        i ++;
    }
    printf("length: %d \n", length);
}
