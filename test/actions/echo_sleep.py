import time
def echo(params,results):
	for key, value in params.items():
		#print(key + " = " + value + " (set)")
		results[key] = value
		#print(key + " = " + value + " (sleep)")
		time.sleep(1)
		#print(key + " = " + value + " (done)")
	return True
