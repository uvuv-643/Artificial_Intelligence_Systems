#include <stdlib.h>
#include <stdio.h>
#include "defs.h"

int* range(int from, int to, int step) {
  int number_of_elements = (to - from + step - 1) / step;
  int* range = (int*)malloc(sizeof(int) * number_of_elements);
  int count = 0;
  for (int i = from; i < to; i += step) {
    range[count++] = i; 
  }
  return range;
}