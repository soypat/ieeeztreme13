def int2base(n, b,dig):
    digits = []
    for i in range(dig):
        if n:
            digits.append(n % b)
            n //= b
        else:
            digits.append(0)
    return digits[::-1]

t = int(input())

for _ in range(t):
    line = input()
    line = line.split()
    B = int(line[0])
    X = int(line[1])

    n = 1
    while X >= n*B**n:
         X -= n*B**n
         n += 1

    d = X //n
    r = X % n
    alfabet = 'abcdefghijklmnopqrstuvwxyz'
    v = int2base(d,B,n)
    print(*alfabet[v[r]])

