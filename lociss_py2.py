import json
import time
import urllib2

req = urllib2.Request("http://api.open-notify.org/iss-now.json")

def locateiss():
    response = urllib2.urlopen(req)
    obj = json.loads(response.read())
    # print "Time : %s | Lat./Long. : %.16f,%.16f" % (obj['timestamp'], obj['iss_position']['latitude'], obj['iss_position']['longitude'])
    print "Time :", obj['timestamp'], "|", "Lat./Long. :", obj['iss_position']['latitude'], ",", obj['iss_position']['longitude']

while True:
    locateiss()
    time.sleep(5)
