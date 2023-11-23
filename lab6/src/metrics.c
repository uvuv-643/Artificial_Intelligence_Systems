#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "defs.h"
#include "types.h"
#include "metrics.h"

char true_class = 'e';
char false_class = 'p';

struct confusion_matrix create_confusion_matrix(int size, column_t expected, column_t predicted) {
  struct confusion_matrix matrix = {0};
  for (int i = 0; i < size; i++) {
    if (expected[i] == true_class) {
      if (predicted[i] == true_class) matrix.tp++;
      if (predicted[i] == false_class) matrix.fp++;
    } else if (expected[i] == false_class) {
      if (predicted[i] == true_class) matrix.fn++;
      if (predicted[i] == false_class) matrix.tn++;
    }
  }
  return matrix;
}

// json encoded array will be written to file with (FP, TP) points
void create_file_with_roc(int size, column_t expected, column_t predicted, FILE* fptr) {
  struct confusion_matrix matrix = create_confusion_matrix(size, expected, predicted);
  struct confusion_matrix dynamic_matrix = {0};
  fprintf(fptr, "[");
  for (int i = 0; i < size; i++) {
    if (expected[i] == true_class) {
      if (predicted[i] == true_class) dynamic_matrix.tp++;
      if (predicted[i] == false_class) dynamic_matrix.fp++;
    } else if (expected[i] == false_class) {
      if (predicted[i] == true_class) dynamic_matrix.fn++;
      if (predicted[i] == false_class) dynamic_matrix.tn++;
    }
    double current_fp_rate = (double)dynamic_matrix.fp / matrix.fp;
    double current_tp_rate = (double)dynamic_matrix.tp / matrix.tp;
    fprintf(fptr, "[%.5f, %.5f]", current_fp_rate, current_tp_rate);
    if (i != size - 1) fprintf(fptr, ", ");
  }
  fprintf(fptr, "]");
}

void create_file_with_pr(int size, column_t expected, column_t predicted, FILE* fptr) {
  struct confusion_matrix dynamic_matrix = {0};
  fprintf(fptr, "[");
  for (int i = 0; i < size; i++) {
    if (expected[i] == true_class) {
      if (predicted[i] == true_class) dynamic_matrix.tp++;
      if (predicted[i] == false_class) dynamic_matrix.fp++;
    } else if (expected[i] == false_class) {
      if (predicted[i] == true_class) dynamic_matrix.fn++;
      if (predicted[i] == false_class) dynamic_matrix.tn++;
    }
    if ((dynamic_matrix.tp + dynamic_matrix.fp) && (dynamic_matrix.tp + dynamic_matrix.fn)) {
      double current_precision = (double)dynamic_matrix.tp / (dynamic_matrix.tp + dynamic_matrix.fp);
      double current_recall = (double)dynamic_matrix.tp / (dynamic_matrix.tp + dynamic_matrix.fn);
      fprintf(fptr, "[%.5f, %.5f]", current_precision, current_recall / current_recall);
      if (i != size - 1) fprintf(fptr, ", ");
    }
  }
  fprintf(fptr, "]");
}

double precision(int size, column_t expected, column_t predicted) {
  struct confusion_matrix matrix = create_confusion_matrix(size, expected, predicted);
  return (double)matrix.tp / (matrix.tp + matrix.fp);
}

double recall(int size, column_t expected, column_t predicted) {
  struct confusion_matrix matrix = create_confusion_matrix(size, expected, predicted);
  return (double)matrix.tp / (matrix.tp + matrix.fn);
}

double f1(int size, column_t expected, column_t predicted) {
  struct confusion_matrix matrix = create_confusion_matrix(size, expected, predicted);
  double precision = (double)matrix.tp / (matrix.tp + matrix.fp);
  double recall = (double)matrix.tp / (matrix.tp + matrix.fn);
  return 2 * precision * recall / (precision + recall);
}

double accuracy(int size, column_t expected, column_t predicted) {
  struct confusion_matrix matrix = create_confusion_matrix(size, expected, predicted);
  return (double)(matrix.tp + matrix.tn) / size;
}


