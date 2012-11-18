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

LIMIT = 1e4
LIMIT = int(LIMIT)

conn = pymongo.Connection()
db = conn.derisive
coll = db.raw_data
newcoll = db.data

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
    for i, val in enumerate(vals): hash_mapping[val.strip(" ")] = i
    mappings[j] = hash_mapping
for j in nums:
    mappings[j] = "NUM"

def y_to_num(grade):
    if grade == "": return -1
    if type(grade) is int or type(grade) is float: return grade
    grade = grade.strip("\r").strip("\"").strip("%").strip(" ")

    if grade == "NP" or grade == "NA" or grade == "N/A" or grade == "Not applicable" or grade == "null" or grade == "NULL": return -1

    if len(grade) == 2 and grade[0] == grade[1]: grade = grade[0]
    
    grade_mappings = {'':-1, 'U':0, 'N':1, 'G':2, 'F':3, 'E':4, 'D':5, 'C':6, 'B':7, 'A':8, 'A*':9}
    if grade in grade_mappings: return grade_mappings[grade]
    elif grade == '*': return grade_mappings['A*']
    elif len(grade) == 2: return grade_mappings[grade[0]]
    else:
        print "Unknown grade %s" % grade
        return -1

def x_to_num(args):
    j, x = args
    if x == "": return -1
    if type(x) is int or type(x) is float: return x
    x = x.strip("\r").strip("\"").strip("%").strip(" ")

    if x == "NP" or x == "NA" or x == "N/A" or x == "Not applicable" or x == "null" or x == "NULL":
        return -1

    try:
        mapping = mappings[j]
    except KeyError as e:
        return -1

    if mapping == "NUM":
        if x == "SUPP": return 0
        try:
            x = float(x)
        except ValueError as e:
            return -1
    else:
        try:
            x = mapping[x]
        except KeyError as e:
            # print "skipping due to not in mapping"
            return -1

    return x

num = coll.count() if coll.count() < LIMIT else LIMIT
X = np.zeros((num, 266))
Y = np.zeros((num, 183))

data = coll.find().limit(LIMIT)
fish = ProgressFish(total=LIMIT)
for i, record in enumerate(data):
     fish.animate(amount=i)
     X[i] = map(x_to_num, enumerate(record['x']))
     Y[i] = map(y_to_num, record['y'])

data_x = open("DATA_X", "w")
cPickle.dump(X, data_x)
data_x.close()

data_y = open("DATA_Y", "w")
cPickle.dump(Y, data_y)
data_y.close()
