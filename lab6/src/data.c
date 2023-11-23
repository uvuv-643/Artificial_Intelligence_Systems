#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <time.h>
#include <math.h>
#include "defs.h"
#include "types.h"
#include "data.h"

// in given dataset last column is target predictor
// from other columns will be taken RANDOM_COLUMN_COUNT items
// copy dataset
struct dataset drop_random_columns(const struct dataset dataset) {
  srand((unsigned int)time(NULL));
  int generated[DATASET_COLUMN_COUNT - 1] = {0};
  int indexes[RANDOM_COLUMN_COUNT] = {0};
  for (int i = 0; i < RANDOM_COLUMN_COUNT; ++i) {
    int random_number;
    do {
        random_number = rand() % (DATASET_COLUMN_COUNT - 1) + 1;
    } while (generated[random_number] == 1);
    generated[random_number] = 1;
    indexes[i] = random_number;
  }
  string_t* labels = (string_t*)malloc(sizeof(string_t) * (RANDOM_COLUMN_COUNT + 1));
  row_t* values = (row_t*)malloc(sizeof(row_t) * (dataset.row_count + 1));
  for (int j = 0; j < RANDOM_COLUMN_COUNT; j++) {
    labels[j] = (string_t)malloc(DATASET_COLUMN_NAME_LENGTH);
    memcpy(labels[j], dataset.labels[indexes[j]], DATASET_COLUMN_NAME_LENGTH);
  }
  for (int i = 0; i < dataset.row_count; i++) {
    values[i] = (row_t)malloc(RANDOM_COLUMN_COUNT + 1);
    for (int j = 0; j < RANDOM_COLUMN_COUNT; j++) {
      values[i][j] = dataset.values[i][indexes[j]];
    }
    values[i][RANDOM_COLUMN_COUNT] = dataset.values[i][0];
  }
  labels[RANDOM_COLUMN_COUNT] = (string_t)malloc(DATASET_COLUMN_NAME_LENGTH);
  memcpy(labels[RANDOM_COLUMN_COUNT], dataset.labels[0], DATASET_COLUMN_NAME_LENGTH);
  return (struct dataset) {
    labels,
    values,
    dataset.row_count,
    RANDOM_COLUMN_COUNT + 1
  };
}

column_t get_target_field(const struct dataset dataset) {
  column_t column = (column_t)malloc(dataset.row_count);
  for (int i = 0; i < dataset.row_count; i++) {
    column[i] = dataset.values[i][dataset.column_count - 1];
  }
  return column;
}

struct train_test_split_data train_test_split(const struct dataset dataset, float test_rate) {
  if (test_rate < 0.01 || test_rate >= 0.5) return (struct train_test_split_data) {0};
  string_t* labels_train = (string_t*)malloc(sizeof(string_t) * dataset.column_count);
  string_t* labels_test = (string_t*)malloc(sizeof(string_t) * dataset.column_count);
  for (int j = 0; j < dataset.column_count; j++) {
    labels_train[j] = (string_t)malloc(DATASET_COLUMN_NAME_LENGTH);
    labels_test[j] = (string_t)malloc(DATASET_COLUMN_NAME_LENGTH);
    memcpy(labels_train[j], dataset.labels[j], DATASET_COLUMN_NAME_LENGTH);
    memcpy(labels_test[j], dataset.labels[j], DATASET_COLUMN_NAME_LENGTH);
  }
  size_t train_size = floor(dataset.row_count * (1 - test_rate));
  size_t test_size = ceil(dataset.row_count * test_rate);
  row_t* values_train = (row_t*)malloc(sizeof(row_t) * train_size);
  row_t* values_test = (row_t*)malloc(sizeof(row_t) * test_size);

  // 0 -> train, 1 -> test;
  int rows_distribution[DATASET_ROW_COUNT] = {0};
  for (int i = 0; i < test_size; i++) {
    int random_number;
    do {
        random_number = rand() % dataset.row_count;
    } while (rows_distribution[random_number] == 1);
    rows_distribution[random_number] = 1;
  }
  int test_p = 0; int train_p = 0;
  for (int i = 0; i < dataset.row_count; i++) {
    if (rows_distribution[i] == 1) {
      values_test[test_p] = (row_t)malloc(dataset.column_count);
      for (int j = 0; j < dataset.column_count; j++) {
        values_test[test_p][j] = dataset.values[i][j];
      }
      test_p++;
    } else {
      values_train[train_p] = (row_t)malloc(dataset.column_count);
      for (int j = 0; j < dataset.column_count; j++) {
        values_train[train_p][j] = dataset.values[i][j];
      }
      train_p++;
    }
  }
  return (struct train_test_split_data) {
    (struct dataset) {
      labels_train,
      values_train,
      train_size,
      dataset.column_count
    },
    (struct dataset) {
      labels_test,
      values_test,
      test_size,
      dataset.column_count
    }
  };
}

// changes current values in given dataset
void fill_empty_values(struct dataset dataset) {
  for (int col = 0; col < dataset.column_count; col++) {
    int count[DATASET_UNIQUE_VALUES] = {0};
    for (int row = 0; row < dataset.row_count; row++) {
      if (dataset.values[row][col] != '?') count[dataset.values[row][col] - 'a']++;
    }
    int most_popular_count = 0; int most_popular_index = -1;
    for (int i = 0; i < DATASET_UNIQUE_VALUES; i++) {
      if (count[i] > most_popular_count) {
        most_popular_index = i;
        most_popular_count = count[i];
      }
    }
    for (int row = 0; row < dataset.row_count; row++) {
      if (dataset.values[row][col] == '?') {
        dataset.values[row][col] = most_popular_index + 'a';
      }
    }
  }
}
