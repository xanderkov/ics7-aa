from random import randint, random

def input_arr():
    alist = input('Enter the list of (nonnegative) numbers: ').split()
    alist = [int(x) for x in alist]
    k = max(alist)
    return alist, k 

def generate_arr(n):
    arr = [0]*n
    for i in range(n):
        arr[i] = randint(0, 1000000)
    return arr