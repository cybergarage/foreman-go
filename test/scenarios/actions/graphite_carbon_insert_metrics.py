import os
import tempfile
import datetime

def graphite_carbon_insert_metrics(params,results):
	plainReq = ""

	## echo is not working ?
	#for key, value in params.items():
	#	if 0 < len(plainReq):
	#		plainReq += "\\n"
	#	plainReq += "%s %s `date +%%s`" % (key, str(value))
	#
	#cmd = "echo -e \"%s\" | nc -c localhost 2003" % (plainReq)

	now = datetime.datetime.now().strftime('%s')
	for key, value in params.items():
		if 0 < len(plainReq):
			plainReq += "\n"
		plainReq += "%s %s %s" % (key, str(value), now)
	with tempfile.NamedTemporaryFile(delete=False, mode='w') as f:
		f.write(plainReq)
		plainReqFilename = f.name
	cmd = "more %s | nc -c localhost 2003" % (plainReqFilename)

	#print(cmd)

	if os.system(cmd) != 0:
		os.remove(plainReqFilename)
		return False

	os.remove(plainReqFilename)

	return True
