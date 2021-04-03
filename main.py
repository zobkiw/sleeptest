import time
from datetime import datetime

while 1:
    print("ABOUT TO SLEEP: ", datetime.utcnow())
    time.sleep(60*60)
    print("AWAKE:          ", datetime.utcnow())
    print("----------------")
