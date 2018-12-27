import foreman
def test_execquery(params,results):
  jsonRes = foreman.execute_query(params["q"])
  if jsonRes is None:
    return False
  for key, value in jsonRes.iteritems():
    results[key] = value
  return True
