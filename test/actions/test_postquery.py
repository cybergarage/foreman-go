import foreman
def test_postquery(params,results):
  jsonRes = foreman.post_query("localhost", 8188, params["q"])
  if jsonRes is None:
    return False
  for key, value in jsonRes.iteritems():
    results[key] = value
  return True
