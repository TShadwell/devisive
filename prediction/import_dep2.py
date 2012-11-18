import pymongo, csv

conn = pymongo.Connection()
db = conn.derisive
mappings = db.dep_mappings_2

dep = open('deprivation.csv', 'r')
depreader = list(csv.reader(dep))
deps = map((lambda p: {'code': p[10], 'dep': p[14:-1]}), depreader)
mappings.insert(deps)

mappings.create_index([("code", pymongo.ASCENDING)])
