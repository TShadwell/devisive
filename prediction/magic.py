import numpy as np
import matplotlib.pyplot as plt
y_labels = open("y.txt").read().split("\n")[0:-1]

i = 53
y_predicted = np.array(map((lambda n: np.float64(n)), open("tmp/predicted%d" % i).read().split("\n")[0:-1]))
y_test = np.array(map((lambda n: np.float64(n.split(" ")[0])), open("tmp/test%d" % i).read().split("\n")[0:-1]))
deltas = np.subtract(y_predicted, y_test)
rms = np.sqrt(np.mean(deltas**2))

print "Subject %s: RMS %f" % (y_labels[i], rms)
# os.system("rm tmp/test%d tmp/train%d tmp/predicted%d" % (i, i, i))

plt.figure(1, (8, 6))
plt.ylabel('Freq.')
plt.title('Predicted vs actual (n = n whole grades difference) on diff. data')
plt.hist(deltas, bins=range(-10,11), histtype='stepfilled', stacked=False, align='mid')
plt.savefig("tmp/hist.png")
plt.show()
