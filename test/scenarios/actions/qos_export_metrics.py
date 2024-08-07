import datetime
import time
import json
import foreman

def qos_export_metrics(params,results):
  now = datetime.datetime.now()
  nowTs = int(time.mktime(now.timetuple()))
  fromTs = nowTs - (60 * 60)
  untilTs = nowTs
  for key, value in params.iteritems():
    q = 'SELECT * FROM METRICS WHERE ts >= %d AND ts <= %d' % (fromTs, untilTs)
    jsonRes = foreman.execute_query(q)
    if jsonRes is None:
      return False
    regKey = 'qos_export_metrics_%s' % key
    foreman.set_register(regKey, str(value))
  return True
