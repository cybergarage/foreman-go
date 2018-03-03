def test_echo(params,results):
	for key, value in params.iteritems():
		results[key] = value
	return True
