import foreman
def post_register(params,results):
        for key, value in params.items():
                foreman.set_register(key, str(value))
                results[key] = value
        return True    
