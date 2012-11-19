import pymongo, simplejson, urllib, csv, time
from fish import ProgressFish


LIMIT = 1e4
LIMIT = int(LIMIT)

conn = pymongo.Connection()
db = conn.derisive

ks51 = db.ks5_0910
ks52 = db.ks5_1011
mappings = db.postcode_lsoa
mappings.drop()

def get_p(di): return set(map((lambda r: r['"SCH_POSTCODE"']), di))

records = list(get_p(ks51.find({},{'"SCH_POSTCODE"':1}).limit(LIMIT)).intersection(get_p(ks52.find({},{'"SCH_POSTCODE"':1}).limit(LIMIT))))
postcodes = map((lambda r: r[1:-1].replace(' ', '')), records)
 
print len(postcodes)

i = 0

fish = ProgressFish(total=len(postcodes))

for postcode in postcodes:
    time.sleep(0.1)
    fish.animate(amount=i)
    try:
        i += 1
        url = "http://mapit.mysociety.org/postcode/%s" % postcode
        result = simplejson.load(urllib.urlopen(url))
        lsoa_code = filter((lambda area: area['type_name'] == "Lower Layer Super Output Area (Full)"), result['areas'].values())[0]['name']
        mappings.insert({'lsoa_code': lsoa_code, 'postcode': postcode})

    except Exception as e:
        print "NOO"

mappings.create_index([("postcode", pymongo.ASCENDING)])
