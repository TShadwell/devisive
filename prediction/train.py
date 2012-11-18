import pymongo, csv, code, sys, os, shutil, code, pprint
import numpy as np
from fish import ProgressFish
from sklearn import preprocessing, svm, datasets
# from matplotlib.pyplot import draw, figure, show
import matplotlib.pyplot as plt
from sklearn.cross_validation import train_test_split
import cPickle

COL_LIMIT = 1

pp = pprint.PrettyPrinter(indent=4)

y_labels = open("y.txt").read().split("\n")[0:-1]
x_labels = open("x_real.txt").read().split("\n")[0:-1]

data_x = open("DATA_X", "r")
X = cPickle.load(data_x)
data_x.close()

data_y = open("DATA_Y", "r")
Y = cPickle.load(data_y)
data_y.close()

print "Scaling"

scaler = preprocessing.Scaler().fit(X)
X = scaler.transform(X)

print "Normalising"

normaliser = preprocessing.Normalizer().fit(X)
X = normaliser.transform(X)

print "Training"

DELTAS = []

imax = COL_LIMIT if COL_LIMIT < Y.shape[1] else Y.shape[1]

for i in range(imax):
    i = 53
    y = Y[:,i]
    took_indexes = (y != -1).nonzero()[0]
    print "Got indexes"
    if len(took_indexes) < 1000: continue
    y = y[took_indexes]
    x = X[took_indexes]

    x_train, x_test, y_train, y_test = train_test_split(x, y, test_size=0.5, random_state=0)

    print "Dumping"
    datasets.dump_svmlight_file(x_train, y_train, "tmp/train%d" % i)
    datasets.dump_svmlight_file(x_test, y_test, "tmp/test%d" % i)
    os.system("cd tmp && csplit -s test%d 3 && mv xx01 test%d && rm xx00" % (i, i))
    os.system("cd tmp && csplit -s train%d 3 && mv xx01 train%d && rm xx00" % (i, i))
    # os.system("svm-train -g 2 -c 8 -q tmp/train%d tmp/model%d" % (i, i))
    os.system("svm-train -g 0.03125 -c 32 tmp/train%d tmp/model%d" % (i, i))
    os.system("svm-predict tmp/test%d tmp/model%d tmp/predicted%d" % (i, i, i))

    y_predicted = np.array(map((lambda n: np.float64(n)), open("tmp/predicted%d" % i).read().split("\n")[0:-1]))
    deltas = np.subtract(y_predicted, y_test)
    rms = np.sqrt(np.mean(deltas**2))

    print "Subject %s: RMS %f" % (y_labels[i], rms)
    # os.system("rm tmp/test%d tmp/train%d tmp/predicted%d" % (i, i, i))
    DELTAS.append(deltas)

DELTAS = [item for sublist in DELTAS for item in sublist]

plt.figure(1, (8, 6))
plt.ylabel('Freq.')
plt.title('Predicted vs actual (n = n whole grades difference) on diff. data')
plt.hist(DELTAS, bins=range(-10,11), histtype='stepfilled', stacked=False, align='mid')
plt.savefig("tmp/hist.png")
plt.show()
