import numpy as np

def generate_matrix(n, m):
    mat = np.random.randint(1,100, (n, m))
    return mat


def print_matrix(matrix, n, m):
    print()
    for i in range(n):
        for j in range(m):
            print(matrix[i][j], end=" ")
        print()
    