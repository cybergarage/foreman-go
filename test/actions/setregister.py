import foreman
def setregister(params,results):
    for key, value in params.items():
        foreman.set_register(key, str(value))
    return True
