#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <sys/wait.h>
#include "defs.h"
#include "types.h"
#include "files.h"
#include "data.h"
#include "decision_tree.h"
#include "metrics.h"

struct thread_data {
    int row_count;
    column_t test_target;
    column_t test_predicted;
};

int main() {

  struct dataset dataset = read_dataset(DATASET_DATA_PATH, DATASET_LABELS_PATH);
  struct dataset new_dataset = drop_random_columns(dataset);
  struct train_test_split_data train_test_data = train_test_split(new_dataset, RANDOM_TEST_RATE);
  struct dataset train = train_test_data.train;
  struct dataset test = train_test_data.test;
  fill_empty_values(train);
  fill_empty_values(test);
  
  struct dt_clf clf = fit(train, get_target_field(train), (struct dt_params) {
    0.01,
    10,
    'e', 'p'
  });

  // print_node_entropy(clf.head, 0, NULL);

  column_t test_target = get_target_field(test);
  column_t test_predicted = predict(clf, test);

  struct thread_data* model_result = (struct thread_data*)malloc(sizeof(struct thread_data));
  model_result->row_count = test.row_count;
  model_result->test_target = test_target;
  model_result->test_predicted = test_predicted;

  printf("\n");
  printf("Accuracy: %.3f\n", accuracy(test.row_count, test_target, test_predicted));
  printf("Precision: %.3f\n", precision(test.row_count, test_target, test_predicted));
  printf("Recall: %.3f\n", recall(test.row_count, test_target, test_predicted));
  printf("F1 score: %.3f\n", f1(test.row_count, test_target, test_predicted));
  
  int pid = fork();
  if (pid == 0) {
    FILE* fptr = fopen(DATASET_ROC_PATH, "w");
    create_file_with_roc(model_result->row_count, model_result->test_target, model_result->test_predicted, fptr);
    fclose(fptr);
    char python_exec[100];
    sprintf(python_exec, "python3 src/visualizer.py %s ROC-Curve FP TP", DATASET_ROC_PATH);
    system(python_exec);
  } else {
    FILE* fptr = fopen(DATASET_PR_PATH, "w");
    create_file_with_pr(model_result->row_count, model_result->test_target, model_result->test_predicted, fptr);
    fclose(fptr);
    char python_exec[100];
    sprintf(python_exec, "python3 src/visualizer.py %s PR-Curve Precision Recall", DATASET_PR_PATH);
    system(python_exec);
    sleep(1);  
  }
  
  return 0;
}