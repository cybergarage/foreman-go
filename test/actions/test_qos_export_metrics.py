import datetime
import time
import json
import foreman

def test_qos_export_metrics(params,results):
  now = datetime.datetime.now()
  nowTs = int(time.mktime(now.timetuple()))
  fromTs = nowTs - (4 * 60)
  untilTs = nowTs
  for key, value in params.iteritems():
    q = 'EXPORT FROM METRICS WHERE ts >= %d AND ts <= %d' % (fromTs, untilTs)
    results['q'] = q
    jsonRes = foreman.execute_query(q)
    if jsonRes is None:
      return False
  return True
