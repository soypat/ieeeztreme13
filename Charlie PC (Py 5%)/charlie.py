import itertools

def cana_input(iterable_list):
    iterable_list.join()
    eval()



T = int(input())

for _ in range(T):
    K = []
    B = int(input())
    nlen = int(input())
    line = input()
    N = [int(x) for x in line.split()]
    for n in range(nlen):
        line = input()
        K.append([int(x) for x in line.split()])

    iterable_list = ['range('+str(x)+'),' for x in N]
    stringy = ''.join(iterable_list)
    iterable = eval( 'itertools.product('+stringy+')' )
    king = 0;
    for i in iterable:

        cand=sum(   ( K[index][i] for index,i in enumerate(i)   ) )

        if king < cand <= B:
            king = cand

    print(king)