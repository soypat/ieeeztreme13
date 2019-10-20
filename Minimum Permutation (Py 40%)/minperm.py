from sys import stdin
input_list = []
for line in stdin:
    input_list.append([int(x) for x in line.split()])

A=input_list[1]
S=input_list[2]

N=len(A)
M=len(S)
#ordena la lista de S

SS=sorted(S)
for i in range(0,M):
    for j in range(0,len(A)):
        if SS[i]<A[j]:
            A.insert(j,SS[i])
            break

print(*A, sep=' ')
