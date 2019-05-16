import foreman
def getregister(params,results):
    for reg_key, value in params.items():
        reg_value = foreman.get_register(reg_key)
        results[reg_key] = reg_value
    return True
