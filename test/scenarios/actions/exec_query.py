import foreman
def exec_query(params,results):
  jsonRes = foreman.execute_query(params["q"])
  if jsonRes is None:
    return False
  for key, value in jsonRes.items():
    results[key] = value
  return True
