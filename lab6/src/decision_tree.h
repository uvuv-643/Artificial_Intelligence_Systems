#ifndef DTREE_H
#define DTREE_H

#include "types.h"

void fill_node(struct dataset dataset, column_t target, struct dt_node* node) ;

struct dt_clf fit(struct dataset dataset, column_t target, struct dt_params params);
void print_node_entropy(struct dt_node* node, int level, char parent_char);
column_t predict(struct dt_clf clf, struct dataset dataset);

#endif