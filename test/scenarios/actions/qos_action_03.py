import foreman
# The action (qos_action_03) is not executed because the destination action (qos_action_02) returns a false
def qos_action_03(params,results):
    for key, value in params.items():
        chkKey = 'qos_action_03_%s' % key
        foreman.set_register(chkKey, str(value))
        results[key] = value
    return True