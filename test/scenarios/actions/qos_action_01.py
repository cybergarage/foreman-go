import foreman
def qos_action_01(params,results):
    for key, value in params.items():
        chkKey = 'qos_action_01_%s' % key
        foreman.set_register(chkKey, str(value))
        # Pass all QoS parameters to qos_action_02
        results[key] = value
    return True
