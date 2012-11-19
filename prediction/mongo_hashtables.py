import pymongo, simplejson, urllib, csv, time
from fish import ProgressFish

conn = pymongo.Connection()
db = conn.derisive
coll = db.schools_data

fields = open('mongo_hashtables.txt', 'r').read().split("\n")[0:-1]
fish = ProgressFish(total=len(fields))
for i, field in enumerate(fields):
    fish.animate(amount=i)
    uniques = "\n".join(coll.distinct(field))
    open("hashtables/%s" % field.replace("/", ""), 'w').write(uniques)


LIMIT = 1e4
LIMIT = int(LIMIT)

conn = pymongo.Connection()
db = conn.derisive
