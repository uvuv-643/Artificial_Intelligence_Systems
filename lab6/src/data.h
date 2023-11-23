#ifndef DATA_H
#define DATA_H

struct dataset drop_random_columns(const struct dataset dataset);
column_t get_target_field(const struct dataset dataset);
struct train_test_split_data train_test_split(struct dataset dataset, float test_rate);
void fill_empty_values(struct dataset dataset);

#endif