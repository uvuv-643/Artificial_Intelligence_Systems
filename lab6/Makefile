CC = gcc
CFLAGS = -Wall -fno-stack-protector -g
LDFLAGS = -lm

OBJ_DIR = obj
SRC_DIR = src
OBJ = $(OBJ_DIR)/files.o $(OBJ_DIR)/main.o $(OBJ_DIR)/data.o $(OBJ_DIR)/decision_tree.o $(OBJ_DIR)/utils.o $(OBJ_DIR)/metrics.o

$(OBJ_DIR)/%.o: $(SRC_DIR)/%.c | $(OBJ_DIR)
	$(CC) $(CFLAGS) -c $< -o $@

run: $(OBJ)
	$(CC) $(CFLAGS) -o main $(OBJ) $(LDFLAGS)

.PHONY: clean

$(OBJ_DIR):
	mkdir -p $(OBJ_DIR)
	
clean:
	rm -f obj/*.o main

debug: CFLAGS += -DDEBUG
debug: run