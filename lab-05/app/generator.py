import numpy as np
N = 200
if __name__ == '__main__':
    a = np.random.randint(2, size=(N, N))
    mat = np.matrix(a)
    with open('data/%s.txt' % N ,'wb') as f:
        for line in mat:
            np.savetxt(f, line, fmt='%d')