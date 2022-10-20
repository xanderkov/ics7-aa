from alg import counting_sort, binary_sort, range_sort
from input import *
# подсчет
# бинарное дерево
# поразрядная 
ARR_LEN = 10


def main():
    # alist, k = input_arr()
    alist = generate_arr(ARR_LEN)
    print(alist)
    k = max(alist)
    sorted_list = counting_sort(alist, k)
    print('Count sort: ', end='')
    print(sorted_list)
    
    print("Binary sort: ", end="")
    arr = binary_sort(alist)
    print(arr)
    
    arr = range_sort(alist)
    print("Range sort: ", end="")
    print(arr)

    

if __name__ == "__main__":
    main()