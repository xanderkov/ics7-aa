# подсчет
# бинарное дерево
# поразрядная 

class TreeNode:
    def __init__(self, value=None):
        self.value = value
        self.left = None
        self.right = None
    
    def inpurt_in_tree(self, elem):
        if self.value is None:
            self.value = elem
            return
        
        if elem < self.value:
            if self.left is None:
                self.left = TreeNode(elem)
            else:
                self.left.inpurt_in_tree(elem)
        
        elif elem > self.value:
            if self.right is None:
                self.right = TreeNode(elem)
            else:
                self.right.inpurt_in_tree(elem)

    def pre_order(self):
        if self.value:
            print(self.value, end=" ")
            
            if self.right:
                self.right.pre_order()
            
            if self.left:
                self.left.pre_order()
            



def binary_sort(alist):
    tree = TreeNode()
    
    for i in alist:
        tree.inpurt_in_tree(i)
    
    tree.pre_order()


def counting_sort(alist, largest):
    c = [0]*(largest + 1)
    for i in range(len(alist)):
        c[alist[i]] = c[alist[i]] + 1
 
    c[0] = c[0] - 1
    for i in range(1, largest + 1):
        c[i] = c[i] + c[i - 1]
 
    result = [None]*len(alist)
 
    for x in reversed(alist):
        result[c[x]] = x
        c[x] = c[x] - 1
 
    return result

def input_arr():
    alist = input('Enter the list of (nonnegative) numbers: ').split()
    alist = [int(x) for x in alist]
    k = max(alist)
    return alist, k 


def main():
    
    # alist, k = input_arr()
    alist = [7, 8, 3, 4, 1, 6]
    k = max(alist)
    sorted_list = counting_sort(alist, k)
    print('Sorted list: ', end='')
    print(sorted_list)

    binary_sort(alist)

    

if __name__ == "__main__":
    main()