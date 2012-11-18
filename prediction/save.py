import pymongo, csv, code, sys, os, shutil, code, pprint
import numpy as np
from fish import ProgressFish
from sklearn import preprocessing, svm, datasets
# from matplotlib.pyplot import draw, figure, show
import matplotlib.pyplot as plt
from sklearn.cross_validation import train_test_split
import cPickle

pp = pprint.PrettyPrinter(indent=4)

y_labels = open("y.txt").read().split("\n")[0:-1]
x_labels = open("x_real.txt").read().split("\n")[0:-1]
conversions = map((lambda r: r.split(",")), open("conversions.csv").read().split("\n"))[0:-1]
nums = map((lambda r: int(r[0])), filter((lambda r: r[1] == 'Num'), conversions))
hashtables = map((lambda r: int(r[0])), filter((lambda r: r[1] == 'Hashtable' or r[1] == 'Grade'), conversions))

def get_unique_vals_col(j):
    # print x_labels[j]
    try:
        return filter((lambda r: r != ""), open("hashtables/%s" % x_labels[j].replace("/",""), "r").read().split("\n"))
    except IOError as e:
        return []
    except Exception as e:
        print "%s: %s" % (x_labels[j], e)
        return []


mappings = {}
for j in hashtables:
    hash_mapping = {}
    vals = get_unique_vals_col(j)
    hash_mapping[""] = -1
    for i, val in enumerate(vals): hash_mapping[val.strip(" ")] = str(i)
    hash_mapping = ["%s,%s" % (h,k) for h,k in hash_mapping.items()]
    open("saved/%d" % j, "w").write("\n".join(hash_mapping))
