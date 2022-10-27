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
def countingSort(array, place):
    size = len(array)
    output = [0] * size
    count = [0] * 10

    for i in range(0, size):
        index = array[i] // place
        count[index % 10] += 1

    for i in range(1, 10):
        count[i] += count[i - 1]

    i = size - 1
    while i >= 0:
        index = array[i] // place
        output[count[index % 10] - 1] = array[i]
        count[index % 10] -= 1
        i -= 1

    for i in range(0, size):
        array[i] = output[i]


def radixSort(array):
    max_element = max(array)

    place = 1
    while max_element // place > 0:
        countingSort(array, place)
        place *= 10
    return array