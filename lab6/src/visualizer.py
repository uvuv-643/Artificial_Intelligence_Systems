import json
import matplotlib.pyplot as plt
import sys

if len(sys.argv) != 5:
  print("bad usage python script")
  sys.exit(1)

file_path_roc = sys.argv[1]
title = sys.argv[2]
X = sys.argv[3]
Y = sys.argv[4]


# Read data from the JSON file
with open(file_path_roc, 'r') as file:
    data = json.load(file)

# Separate x and y coordinates
x_coordinates = [point[0] for point in data]
y_coordinates = [point[1] for point in data]

# Create a line plot
plt.plot(x_coordinates, y_coordinates, marker='o', linestyle='-')

# Set labels and title
plt.xlabel(X)
plt.ylabel(Y)
plt.title(title)

# Display the plot (optional)
plt.show()
