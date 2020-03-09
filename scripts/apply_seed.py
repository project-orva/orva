# python 2.7.16
import requests
import json

from postgres import Postgres

def toSqlType(value):
    s = {
        "unicode": "text",
        "int": "int"
    }

    return s[type(value).__name__] 

def applyPostgres(conn, seedData):
    db = Postgres(url=conn)

    for table in seedData.keys():
        # drop preexisting tables
        db.run("DROP TABLE {}".format(table))
        print("[postgresql] table dropped {}".format(table))

        # create new tables
        props = ""
        for prop in seedData[table][0]:
            cProp = seedData[table][0][prop]
            
            cq = ""
            if prop == "id":
                cq = "id varchar(40) PRIMARY KEY"
            else:
                cq = "{} {} NOT NULL".format(prop, toSqlType(cProp))

            if props is "":
                props = cq
            else:
                props = "{}, {}".format(props, cq)

        db.run("CREATE TABLE {} ({})".format(table, props))
        print("[postgresql] table created {}".format(table))

        # insert data into new tables
        for row in seedData[table]:
            p = ""

            for x in row.values():
                ux = x
                if type(x).__name__ == "unicode": 
                    ux = "'{}'".format(x)

                if p is "":
                    p = ux
                else:
                    p = "{}, {}".format(p, ux)

            db.run("INSERT INTO {} ({}) VALUES ({})".format(table, ', '.join(row.keys()), p))
            # print("[postgresql] added row")

def applyCouchDB(couchURL, seedData):  
    for k in seedData.keys():
        db = "{}/{}".format(couchURL, k)
        requests.delete(db)
        print('[couchDB] cleaned db: {}'.format(k))  

        requests.put(db)

        for keyIndicies in seedData[k]:
            url = "{}/{}".format(db, keyIndicies['id'])
            resp = requests.put(url, data=json.dumps(keyIndicies))

if __name__ == "__main__":
    seedData = ""
    with open("seed.json", 'r') as content_file:
        seedData = content_file.read()

    j = json.loads(seedData)

    applyCouchDB("http://0.0.0.0:5984", j["couch"])
    print('[couchDB] seed successfully applied')

    applyPostgres("postgresql://docker:docker@0.0.0.0:32771", j["tables"])
    print('[postgresql] seed succesffuly applied')