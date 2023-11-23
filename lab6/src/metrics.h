#ifndef METRICS_H
#define METRICS_H

struct confusion_matrix {
  int tp;
  int fp;
  int tn;
  int fn;
};

void create_file_with_roc(int size, column_t expected, column_t predicted, FILE* fptr);
void create_file_with_pr(int size, column_t expected, column_t predicted, FILE* fptr);
double precision(int size, column_t expected, column_t predicted);
double recall(int size, column_t expected, column_t predicted);
double f1(int size, column_t expected, column_t predicted);
double accuracy(int size, column_t expected, column_t predicted);

#endif