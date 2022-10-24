from time import process_time
from alg import simple_mult, winograd_mult, winograd_mult_opt
from io_m import generate_matrix, print_matrix

N_TEST = 100
N = 10
M = 10

def test_simple_mult(A, B):
    # Start the stopwatch / counter 
    t1_start = process_time() 
    for i in range(N_TEST):
        simple_mult(A, M, B, N, M)
    # Stop the stopwatch / counter
    t1_stop = process_time()
    
    print("Simple mult in sec:", t1_stop-t1_start)

def test_winograd_mult(A, B):
    # Start the stopwatch / counter 
    t1_start = process_time() 
    for i in range(N_TEST):
        winograd_mult(A, M, B, N, M)
    # Stop the stopwatch / counter
    t1_stop = process_time()
    
    print("Winograd mult in sec:", t1_stop-t1_start)

def test_winograd_mult_opt(A, B):
    # Start the stopwatch / counter 
    t1_start = process_time() 
    for i in range(N_TEST):
        winograd_mult_opt(A, M, B, N, M)
    # Stop the stopwatch / counter
    t1_stop = process_time()
    
    print("Winograd mult optimized in sec:", t1_stop-t1_start)


def main():
    n = 100
    A = generate_matrix(M, N)
    B = generate_matrix(N, M)
    
    test_simple_mult(A, B)
    test_winograd_mult(A, B)
    test_winograd_mult_opt(A, B)
    
if __name__=="__main__":
    main()