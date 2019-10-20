from sys import stdin
import numpy as np

n = int(stdin.readline())

inputs = [[int(x) for x in line.split() ] for line in stdin]

board = np.array(inputs[:n])
target = np.array(inputs[n:])

board = board[np.argsort(board[:,0]),:]
target = target[np.argsort(target[:,0]),:]

board_index = np.argsort(board[:,1])
target_index = np.argsort(target[:,1])

mensaje = []
count = 0;

for i in range(n):
    while board[i,0] != -2*(n-i-1):
        mensaje.append(str(board[i,0]) + ' ' + str(board[i,1]) + ' L \n')
        count += 1
        board[i,0] -= 1

for current,i in enumerate(board_index):
    while board[i,1] != -2*(n-current-1):
        mensaje.append(str(board[i,0]) + ' ' + str(board[i,1]) + ' D\n')
        count += 1
        board[i,1] -= 1

for cursor in range(n):
    bi = np.flatnonzero(target_index[cursor] == board_index)
    bi = bi[0]

    while cursor != bi:
        a = board_index[bi]
        b = board_index[bi-1]
        mensaje.append(str(board[a, 0]) + ' ' + str(board[a, 1])    + ' D\n')
        mensaje.append(str(board[a, 0]) + ' ' + str(board[a, 1]-1)  + ' T\n')
        mensaje.append(str(board[b, 0]) + ' ' + str(board[b, 1])    + ' U\n')
        mensaje.append(str(board[b, 0]) + ' ' + str(board[b, 1]+ 1) + ' U\n')
        mensaje.append(str(board[a, 0]) + ' ' + str(board[a, 1]-1)  + ' P\n')
        mensaje.append(str(board[a, 0]) + ' ' + str(board[a, 1]-1)  + ' D\n')

        count +=6
        board[a, 1] -= 2
        board[b, 1] += 2

        board_index[bi],board_index[bi-1] = board_index[bi-1],board_index[bi]

        bi -= 1


for i in reversed(range(n)):
    while board[i,0] != target[i,0]:
        mensaje.append(str(board[i,0]) + ' ' + str(board[i,1]) + ' R\n')
        count +=1
        board[i,0] += 1

for i in reversed(target_index):
    while board[i,1] != target[i,1]:
        mensaje.append(str(board[i,0]) + ' ' + str(board[i,1]) + ' U\n')
        count += 1
        board[i,1] += 1

print(count)
print(''.join(mensaje))