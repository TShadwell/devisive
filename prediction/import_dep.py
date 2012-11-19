import pymongo, csv

conn = pymongo.Connection()
db = conn.derisive
mappings = db.dep_mappings

dep = open('deprivation.csv', 'r')
depreader = list(csv.reader(dep))
deps = map((lambda p: {'postcode': p[11], 'dep': p[14:-1]}), depreader)
mappings.insert(deps)

mappings.create_index([("postcode", pymongo.ASCENDING)])
