import pymongo, csv, code, sys
from fish import ProgressFish

LIMIT=1e7
LIMIT = int(LIMIT)

conn = pymongo.Connection()
db = conn.derisive

schools_data = db.schools_data
dep_mappings = db.dep_mappings
dep_2_mappings = db.dep_mappings_2
postcode_mappings = db.postcode_lsoa
postcodes = db.postcodes
ks51 = db.ks5_0910
ks52 = db.ks5_1011
coll = db.raw_data

coll.drop()

y_keys = map((lambda r: "\"%s\"" % r.rstrip()), list(open('y.txt', 'r')))
x_keys = map((lambda r: "\"%s\"" % r.rstrip()), list(open('x.txt', 'r')))

records1 = ks51.find({"\"KS5_ALEV\"":1}).limit(LIMIT)
records2 = ks52.find({"\"KS5_ALEV\"":1}).limit(LIMIT)

def get_o_list(record, li):
    record['"KS2_SCILEV"'] = record['"KS2_SCILEV"\r']
    return map((lambda k: record[k.upper()] if (k.upper() in record) else ""), li)

def do_records(records):
    num_records = LIMIT
    fish = ProgressFish(total=num_records)
    for i, record in enumerate(records):
        fish.animate(amount=i)
        keys = record.keys()
        x = get_o_list(record, x_keys)
        y = get_o_list(record, y_keys)

        y_count = len(filter((lambda r: r != "\"\"" and r != "''" and r != ""), y))
        x_count = len(filter((lambda r: r != "\"\"" and r != "''" and r != ""), x))

        if y_count == 0 or x_count == 0: continue

        home_dep_details = [""]*16

        try:
            home_lsoa_code = record['"CEN_LSOA"'][1:-1]
            if home_lsoa_code != None and home_lsoa_code != '':
                home_dep_details = dep_2_mappings.find_one({'code': home_lsoa_code})['dep']
        except Exception as e:
            print "home - probably wales/scotland"
            print e

        x += home_dep_details

        sd = [""]*133

        try:
            sd2 = schools_data.find_one({'KS5_11SCHNAME': record["\"SCH_SCHOOLNAME\""][1:-1]}, {'_id': 0})
            if sd2 != None: sd = map((lambda k: sd2[k]), ["LURN", "LLA", "LESTAB", "LLAESTAB", "LSCHNAME", "LSTREET", "LLOCALITY", "LADDRESS3", "LTOWN", "LPOSTCODE", "LTELNUM", "LICLOSE", "LISNEW", "LMINORGROUP", "LNFTYPE", "LISPRIMARY", "LISSECONDARY", "LISPOST16", "LAGEL", "LAGEH", "LGENDER", "LSFGENDER", "LRELDENOM", "LADMPOL", "LNEWACFLAG", "KS5_11RECTYPE", "KS5_11ALPHAIND", "KS5_11REGION", "KS5_11LASORT", "KS5_11LEA", "KS5_11ESTAB", "KS5_11URN", "KS5_11SCHNAME_AC", "KS5_11SCHNAME", "KS5_11ADDRESS1", "KS5_11ADDRESS2", "KS5_11ADDRESS3", "KS5_11TOWN", "KS5_11PCODE", "KS5_11TELNUM", "KS5_11CONTFLAG", "KS5_11NFTYPE", "KS5_11RELDENOM", "KS5_11ADMPOL", "KS5_11GENDER1618", "KS5_11FEEDER", "KS5_11AGERANGE", "KS5_11ICLOSE", "KS5_11TABKS2", "KS5_11TAB15", "KS5_11EXAMCONF", "KS5_11DUMMY1", "KS5_11TPUP1618", "KS5_11TALLPUPA", "KS5_11TALLPPSA", "KS5_11TALLPPEA", "KS5_11PTPASS1L3", "KS5_11PTPASS2LV3", "KS5_11PTPASS3LV3", "KS5_11TALLPPS08", "KS5_11TALLPPS09", "KS5_11TALLPPS10", "KS5_11TALLPPE08", "KS5_11TALLPPE09", "KS5_11TALLPPE10", "ABS_11LA", "ABS_11ESTAB", "ABS_11URN", "ABS_11PERCTOT", "ABS_11PERCUA", "ABS_11PPERSABS15", "ABS_11PPERSABS20", "CFR_11URN", "CFR_11LANUMBER", "CFR_11LONDON/NON-LONDON", "CFR_11MEDIAN", "CFR_11PUPILS", "CFR_11FSM", "CFR_11FSMBAND", "CFR_11GRANTFUNDING", "CFR_11SELFGENINCOME", "CFR_11TOTALINCOME", "CFR_11TEACHINGSTAFF", "CFR_11SUPPLYTEACHERS", "CFR_11EDUCATIONSUPPORTSTAFF", "CFR_11PREMISES", "CFR_11BACKOFFICE", "CFR_11CATERING", "CFR_11OTHERSTAFF", "CFR_11ENERGY", "CFR_11LEARNINGRESOURCES", "CFR_11ICT", "CFR_11BOUGHTIN", "CFR_11OTHER", "CFR_11TOTALEXPENDITURE", "SWF_11LA", "SWF_11URN", "SWF_11NTEA", "SWF_11NTEAAS", "SWF_11NNONTEA", "SWF_11NFTETEA", "SWF_11NFTETEAAS", "SWF_11RATPUPTEA", "SWF_11SALARY", "CENSUS_11URN", "CENSUS_11LAESTAB", "CENSUS_11NUMFTE", "CENSUS_11TOTPUPSENDN", "CENSUS_11TSENSAP", "CENSUS_11TSENA", "CENSUS_11TOTSENST", "CENSUS_11TOTSENAP", "CENSUS_11PSENSAP", "CENSUS_11PSENA", "CENSUS_11PTOTSENST", "CENSUS_11PTOTSENAP", "CENSUS_11TOTPUPEALDN", "CENSUS_11NUMEAL", "CENSUS_11NUMENGFL", "CENSUS_11NUMUNCFL", "CENSUS_11PNUMEAL", "CENSUS_11PNUMENGFL", "CENSUS_11PNUMUNCFL", "CENSUS_11TOTPUPFSMDN", "CENSUS_11NUMFSM", "CENSUS_11NUMNOFSM", "CENSUS_11PNUMFSM", "CENSUS_11PNUMNOFSM", "OLA", "OURN", "OSCHOOLNAME", "OPHASE", "OREPORTURL"])
        except Exception as e:
            print "school details"
            print e

        x += sd

        school_dep_details = [""]*16

        try:
            school_postcode = record["\"SCH_POSTCODE\""][1:-1]
            school_lsoa_code = postcodes.find_one({'Postcode2': school_postcode}, {'Code':1})
            if school_lsoa_code != None:
                school_lsoa_code = school_lsoa_code['Code']
                school_dep_details = dep_2_mappings.find_one({'code': school_lsoa_code})['dep']
        except Exception as e:
            print "school deps"
            print e

        x += school_dep_details

        coll.insert({'x': x, 'y': y})

do_records(records1)
do_records(records2)
