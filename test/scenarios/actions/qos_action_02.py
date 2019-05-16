import foreman
def qos_action_02(params,results):
    for key, value in params.items():
        chkKey = 'qos_action_02_%s' % key
        foreman.set_register(chkKey, str(value))
        results[key] = value
    # Returns a false not to run the related routes
    return False