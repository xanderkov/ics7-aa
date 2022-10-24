import numpy as np

def simple_mult(mat1, m, mat2, n, q):
    res = np.zeros([m, q])
    for i in range(m):
        for j in range(q):
            for k in range(n):
                res[i][j] = res[i][j] + mat1[i][k] * mat2[k][j]
    return res

def precompile_rows_win(mat, n, m):
    mh = np.zeros([n])
    
    for i in range(n):
        for j in range(m // 2):
            mh[i] = mh[i] + mat[i][j * 2] * mat[i][j * 2 + 1]
    return mh

def precompile_cols_win(mat, n, m):
    mv = np.zeros([m])
    
    for i in range(m):
        for j in range(n // 2):
            mv[i] = mv[i] + mat[j * 2][i] * mat[j * 2 + 1][i]
    return mv

def winograd_mult(A, m, B, n, q):
    res = np.zeros([m, q])
    mh = precompile_rows_win(A, m, n)
    mv = precompile_cols_win(B, n, q)
    for i in range(m):
        for j in range(q):
            res[i][j] = -mh[i] - mv[j]
            for k in range(n // 2):
                res[i][j] = res[i][j] + (A[i][k*2] + B[k*2+1][j])*(A[i][k*2+1] + B[k*2][j])
    if n % 2 != 0:
        for i in range(n):
            for j in range(m):
                res[i][j] = res[i][j] + A[i][n-1]*B[n-1][j]
    return res

def precompile_rows_win_opt(mat, n, m):
    mh = np.zeros([n])
    
    opt = m // 2
    for i in range(n):
        for j in range(opt):
            t = j << 1
            mh[i] += mat[i][t] * mat[i][t + 1]
    return mh

def precompile_cols_win(mat, n, m):
    mv = np.zeros([m])
    
    opt = n // 2
    for i in range(m):
        for j in range(opt):
            t = j << 1
            mv[i] += mat[t][i] * mat[t + 1][i]
    return mv

def winograd_mult_opt(A, m, B, n, q):
    res = np.zeros([m, q])
    mh = precompile_rows_win(A, m, n)
    mv = precompile_cols_win(B, n, q)
    
    opt = n // 2
    for i in range(m):
        for j in range(q):
            res[i][j] = -mh[i] - mv[j]
            for k in range(n // 2):
                t = k << 1
                res[i][j] += (A[i][t] + B[t+1][j])*(A[i][t+1] + B[t][j])
    if n % 2 != 0:
        for i in range(n):
            for j in range(m):
                res[i][j] += A[i][n-1]*B[n-1][j]
    return res



