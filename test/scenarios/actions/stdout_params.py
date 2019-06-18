import foreman
def stdout_params(params,results):
        for key, value in params.items():
                msg = '[%s] = %s' % (key, value)
                print(msg)
        return True    
