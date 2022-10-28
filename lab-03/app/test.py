from time import process_time
from input import generate_arr
from alg import counting_sort, radixSort, binary_sort

N = 100

def test_count_sort(arr):
    # Start the stopwatch / counter 
    t1_start = process_time() 
    k = max(arr)
    for i in range(N):
        counting_sort(arr, k)
    # Stop the stopwatch / counter
    t1_stop = process_time()
    
    print("Count sort in sec:", t1_stop-t1_start)

def test_range_sort(arr):
    # Start the stopwatch / counter 
    t1_start = process_time() 
    for i in range(N):
        radixSort(arr)
    # Stop the stopwatch / counter
    t1_stop = process_time()
    
    print("Radix sort in sec:", t1_stop-t1_start)


def test_binary_sort(arr):
    # Start the stopwatch / counter 
    t1_start = process_time() 
    for i in range(N):
        binary_sort(arr)
    # Stop the stopwatch / counter
    t1_stop = process_time()
    
    print("Binary sort in sec:", t1_stop-t1_start)

def generate_arr_posl(n):
    arr = [0]*n
    for i in range(n):
        arr[i] = i

def generate_arr_vniz(n):
    arr = [0]*n
    for i in range(n):
        arr[i] = n - i

def main():
    n = 100
    arr = generate_arr(n)
    
    test_count_sort(arr)
    test_binary_sort(arr)
    test_range_sort(arr)
    
    
if __name__=="__main__":
    main()