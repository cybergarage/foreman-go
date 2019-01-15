import datetime
import time
import json
import foreman

def test_qos_analyze_correlation(params,results):
  nowTs = int(time.mktime(now.timetuple()))
  fromTs = nowTs - (60 * 60)
  untilTs = nowTs
  for key, value in params.iteritems():
    q = 'ANALYZE %s FROM METRICS WHERE ts >= %d AND ts <= %d' % (key, fromTs, untilTs)
    results['q'] = q
    jsonRes = foreman.execute_query(q)
    if jsonRes is None:
      return True
  return True
