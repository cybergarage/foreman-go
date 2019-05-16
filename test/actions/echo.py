def echo(params,results):
	for key, value in params.items():
		results[key] = value
	return True
