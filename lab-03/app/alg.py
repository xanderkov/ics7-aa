# Сортировка мудрая
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
        
        elif elem >= self.value:
            if self.right is None:
                self.right = TreeNode(elem)
            else:
                self.right.inpurt_in_tree(elem)

    def pre_order(self, arr):
        if self.value:
            if self.left:
                self.left.pre_order(arr)
            if self.value: 
                arr.append(self.value)
                
            if self.right:
                self.right.pre_order(arr)
            
def binary_sort(alist):
    tree = TreeNode()
    
    for i in alist:
        tree.inpurt_in_tree(i)
    arr = []
    tree.pre_order(arr)
    return arr


# Сортировка подсчетом
def counting_sort(alist, largest):
    c = [0]*(largest + 1)
    for i in range(len(alist)):
        c[alist[i]] += 1
 
    c[0] -= 1
    for i in range(1, largest + 1):
        c[i] += c[i - 1]
 
    result = [0]*len(alist)
 
    for x in reversed(alist):
        result[c[x]] = x
        c[x] -= 1
 
    return result

# Сортировка поразрадяная
def range_sort(arr):
    length = len(str(max(arr)))
    rang = 10
    for i in range(length):
        B = [[] for k in range(rang)] #список длины range, состоящий из пустых списков
        for x in arr:
            figure = x // 10**i % 10
            B[figure].append(x)
        arr= []
        for k in range(rang):
            arr = arr + B[k]
    
    return arr