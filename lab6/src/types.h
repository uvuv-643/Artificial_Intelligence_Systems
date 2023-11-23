#ifndef TYPES_H
#define TYPES_H

typedef char* row_t;
typedef char* string_t;
typedef char* column_t;

struct dataset {
  string_t* labels;
  row_t* values;
  int row_count;
  int column_count;
};

struct train_test_split_data {
  struct dataset train;
  struct dataset test;
};

struct dt_clf {
  struct dt_node* head;
};

struct dt_node {
  int row_count;
  int column_count;
  int* row_indexes; // row_indexes.length == row_count
  int* column_indexes; // column_indexes.length == column_count
  int child_column;
  float entropy;
  char response; // if response == 0 - not a leaf
  struct dt_node** children;
};

struct dt_entropy_children {
  double information_gain;
  double* entropies;
};

struct dt_params {
  double max_entropy;
  int max_number_to_split;
  char true_class_mark;
  char false_class_mark;
};

#endif