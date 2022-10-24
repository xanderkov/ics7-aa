import numpy as np
from alg import simple_mult, winograd_mult, winograd_mult_opt
from io_m import generate_matrix, print_matrix
N1 = 3
M1 = 3
N2 = 3
M2 = 3


def main():
    m, n, q = N1, M1, M2
    mat_1 = generate_matrix(m, n)
    mat_2 = generate_matrix(n, q)
    
    print_matrix(mat_1, m, n)
    print_matrix(mat_2, n, q)
    
    res = simple_mult(mat_1, m, mat_2, n, q)
    print_matrix(res, m, q)
    
    res = winograd_mult(mat_1, m, mat_2, n, q)
    print_matrix(res, m, q)
    
    res = winograd_mult_opt(mat_1, m, mat_2, n, q)
    print_matrix(res, m, q)

if __name__ == '__main__':
    main()