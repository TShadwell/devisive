import itertools, sys

o = open('OUTPUT', 'r').read().split("\n")[0:-1]
h = open('hashtables.txt', 'r').read().split("\n")[0:-1]

groups = []
for k, g in itertools.groupby(o, (lambda r: r in h)):
    groups.append(list(g))

for i in range(len(groups)/2):
    key = groups[i*2][0]
    vals = "\n".join(groups[i*2+1])

    open("hashtables/%s" % key.replace("/", ""), 'w').write(vals)
