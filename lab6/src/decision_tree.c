#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <math.h>
#include "defs.h"
#include "types.h"
#include "files.h"
#include "decision_tree.h"

// default values
int max_numbers_to_split = 10; 
double max_entropy = 0.05;
char true_class_mark = 'e';
char false_class_mark = 'p';

double entropy(double p) {
  if (p == 0) return 0;
  return -p * log2(p);
}

double H(struct dataset dataset, column_t target) {
  double result = 0;
  int count[DATASET_UNIQUE_VALUES] = {0};
  for (int i = 0; i < dataset.row_count; i++) {
    count[target[i] - 'a']++;
  }
  for (int i = 0; i < DATASET_UNIQUE_VALUES; i++) {
    result += entropy((double)count[i] / dataset.row_count);
  }
  return result;
}

struct dt_entropy_children calculate_IG(struct dataset dataset, struct dt_node* node, int col, column_t target) {
  double currCH = 0;
  int true_class[DATASET_UNIQUE_VALUES] = {0};
  int false_class[DATASET_UNIQUE_VALUES] = {0};
  int true_class_cnt = 0; int false_class_cnt = 0; 
  for (int i = 0; i < node->row_count; i++) {
    int row = node->row_indexes[i];
    if (target[row] == true_class_mark) {
      true_class[dataset.values[row][col] - 'a']++; 
      true_class_cnt++;
    } else {
      false_class[dataset.values[row][col] - 'a']++;
      false_class_cnt++;
    }
  }
  double entropy_true = 0; double entropy_false = 0;
  for (int i = 0; i < DATASET_UNIQUE_VALUES; i++) {
    entropy_true += entropy((double)true_class[i] / dataset.row_count);
    entropy_false += entropy((double)false_class[i] / dataset.row_count);
  }
  double* entropies = (double*) malloc(sizeof(double) * DATASET_UNIQUE_VALUES);
  for (int i = 0; i < DATASET_UNIQUE_VALUES; i++) {
    double currH = entropy((double)false_class[i] / dataset.row_count) + entropy((double)true_class[i] / dataset.row_count);
    entropies[i] = currH;
    currCH += currH * (true_class[i] + false_class[i]);
  }
  currCH /= dataset.row_count;
  return (struct dt_entropy_children) {
    node->entropy - currCH,
    entropies
  };
}

void set_responses(struct dataset dataset, column_t target, struct dt_node* node) {
  if (node == NULL) return;
  if (node->children) {
    int child_cnt = 0;
    for (int i = 0; i < DATASET_UNIQUE_VALUES; i++) {
      if (node->children[i] != NULL) {
        set_responses(dataset, target, node->children[i]);
        child_cnt++;
      }
    }
    if (child_cnt) return;
  }
  int true_class_cnt = 0; int false_class_cnt = 0;
  for (int i = 0; i < node->row_count; i++) {
    int row = node->row_indexes[i];
    if (target[row] == true_class_mark) true_class_cnt++;
    else false_class_cnt++;
  }
  if (true_class_cnt > false_class_cnt) node->response = true_class_mark;
  else node->response = false_class_mark;
}

void create_new_nodes_for_column(struct dataset dataset, column_t target, struct dt_node* parent, int col_index, double* entropies) {  
  int count[DATASET_UNIQUE_VALUES] = {0};
  int col = parent->column_indexes[col_index];
  for (int i = 0; i < parent->row_count; i++) {
    int row = parent->row_indexes[i];
    count[dataset.values[row][col] - 'a']++;
  }
  int count_indexes[DATASET_UNIQUE_VALUES] = {0};
  int* row_elements[DATASET_UNIQUE_VALUES] = {0};
  for (int i = 0; i < DATASET_UNIQUE_VALUES; i++) {
    if (!count[i]) continue;
    row_elements[i] = (int*)malloc(sizeof(int) * count[i]);
  }
  for (int i = 0; i < parent->row_count; i++) {
    int row = parent->row_indexes[i];
    int row_value_index = dataset.values[row][col] - 'a';
    row_elements[row_value_index][count_indexes[row_value_index]++] = row;
  }
  int* column_element = (int*)malloc(sizeof(int) * (parent->column_count - 1));
  int column_count = 0;
  for (int i = 0; i < parent->column_count; i++) {
    int c_col = parent->column_indexes[i];
    if (c_col == col) continue;
    column_element[column_count++] = c_col;
  }
  struct dt_node** children = (struct dt_node**)malloc(sizeof(struct dt_node*) * DATASET_UNIQUE_VALUES);
  for (int i = 0; i < DATASET_UNIQUE_VALUES; i++) {
    if (count[i] == 0) continue;
    struct dt_node* new_node = (struct dt_node*)malloc(sizeof(struct dt_node));
    memset(new_node, 0, sizeof(struct dt_node));
    new_node->entropy = entropies[i];
    new_node->row_count = count[i];
    new_node->column_count = parent->column_count - 1;
    new_node->row_indexes = row_elements[i];
    new_node->column_indexes = column_element;
    children[i] = new_node;
    if (new_node->entropy > max_entropy && new_node->row_count > max_numbers_to_split) {
      fill_node(dataset, target, new_node);
    } else {
      new_node->children = NULL;
    }
  }
  parent->children = children;
  
}

void fill_node(struct dataset dataset, column_t target, struct dt_node* node) {
  double max_IG = 0;
  double IGs[DATASET_COLUMN_COUNT] = {0};
  double* entropies[DATASET_COLUMN_COUNT] = {0};
  for (int i = 0; i < node->column_count; i++) {
    int col = node->column_indexes[i];
    struct dt_entropy_children entropy_data = calculate_IG(dataset, node, col, target);
    IGs[i] = entropy_data.information_gain;
    entropies[i] = entropy_data.entropies;
    max_IG = fmax(max_IG, IGs[i]);
  }
  for (int i = 0; i < node->column_count; i++) {
    if (max_IG == IGs[i]) {
      int col = node->column_indexes[i];
      node->child_column = col;
      create_new_nodes_for_column(dataset, target, node, i, entropies[i]);  
      return;  
    }
  }
}

struct dt_clf fit(struct dataset dataset, column_t target, struct dt_params params) {
  true_class_mark = params.true_class_mark;
  false_class_mark = params.false_class_mark;
  max_entropy = params.max_entropy;
  max_numbers_to_split = params.max_number_to_split;
  struct dt_node* head = (struct dt_node*)malloc(sizeof(struct dt_node));
  memset(head, 0, sizeof(struct dt_node));
  head->row_count = dataset.row_count;
  head->column_count = dataset.column_count - 1;
  head->row_indexes = range(0, dataset.row_count, 1);
  head->column_indexes = range(0, dataset.column_count - 1, 1);
  head->entropy = H(dataset, target);
  fill_node(dataset, target, head);
  set_responses(dataset, target, head);
  return (struct dt_clf) {
    head
  };
}

char predict_single(struct dt_clf clf, row_t row) {
  struct dt_node* head = clf.head;
  struct dt_node* node = head;
  while (!node->response) {
    char target_col_value = row[node->child_column] - 'a';
    if (node->children[target_col_value]) {
      node = node->children[target_col_value];
    } else {
      // new value -> pseudo-random choise
      for (int i = 0; i < DATASET_UNIQUE_VALUES; i++) {
        if (node->children[i]) {
          node = node->children[i];
          break;
        }
      }
    }
  }
  return node->response;
}

column_t predict(struct dt_clf clf, struct dataset dataset) {
  char* predicted = (char*)malloc(dataset.row_count);
  for (int i = 0; i < dataset.row_count; i++) {
    predicted[i] = predict_single(clf, dataset.values[i]);
  }
  return predicted;
}

void print_node_entropy(struct dt_node* root, int depth, char parent_char) {
  if (root == NULL) return;
  for (int i = 0; i < depth; i++) printf("|   ");
  if (root->response) {
    printf("|-- [%d, %.2f %c %c]\n", root->row_count, root->entropy, parent_char, root->response);
  } else {
    printf("|-- [%d, %.2f %c %d]\n", root->row_count, root->entropy, parent_char, root->child_column);
  }
  if (root->children != NULL) {
    for (int i = 0; i < DATASET_UNIQUE_VALUES; i++) {
      print_node_entropy(root->children[i], depth + 1, i + 'a');
    }
  }
}