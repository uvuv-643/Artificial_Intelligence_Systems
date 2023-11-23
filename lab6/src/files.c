#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "defs.h"
#include "types.h"
#include "files.h"

struct dataset read_dataset(const char* dataset_path, const char* label_path) {

  FILE *fileptr = fopen(dataset_path, "r");
  if (fileptr == NULL) {
    perror("Error opening dataset file");
    exit(EXIT_FAILURE);
  }

  int row_count = 0;
  char databuff[2 * DATASET_LABELS_MAX_ROW_LENGTH - 1];
  char** values = (char**)malloc(DATASET_ROW_COUNT * sizeof(char*)); 
  while (fgets(databuff, 2 * DATASET_LABELS_MAX_ROW_LENGTH - 1, fileptr) != NULL) {
    char* valuesbuff; int col_count = 0;
    char* values_in_row = (char*)malloc(DATASET_COLUMN_COUNT * sizeof(char));
    valuesbuff = strtok(databuff, DATASET_SEPARATOR);
    do {
      values_in_row[col_count++] = valuesbuff[0];
    } while ((valuesbuff = strtok(NULL, DATASET_SEPARATOR)) != NULL);
    values[row_count++] = values_in_row;
  }
  fclose(fileptr);

  fileptr = fopen(label_path, "r");
  if (fileptr == NULL) {
    perror("Error opening dataset labels file");
    exit(EXIT_FAILURE);
  }

  char** labels = malloc(DATASET_COLUMN_COUNT * sizeof(char*)), *label; 
  int column_number = 0;
  char labelsbuff[DATASET_LABELS_MAX_ROW_LENGTH];
  if (fgets(labelsbuff, DATASET_LABELS_MAX_ROW_LENGTH, fileptr) != NULL) {
    label = strtok(labelsbuff, DATASET_SEPARATOR);
    do {
      labels[column_number] = malloc(DATASET_COLUMN_NAME_LENGTH);
      strcpy(labels[column_number++], label);
    } while ((label = strtok(NULL, DATASET_SEPARATOR)) != NULL);
  }
  fclose(fileptr);
  return (struct dataset) {
    labels,
    values,
    DATASET_ROW_COUNT,
    DATASET_COLUMN_COUNT
  };

}