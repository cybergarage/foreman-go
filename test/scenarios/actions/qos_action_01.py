import foreman
def test_qos_action_01(params,results):
    for key, value in params.iteritems():
        chkKey = 'qos_action_01_%s' % key
        foreman.set_register(chkKey, str(value))
        # Pass all QoS parameters to test_qos_action_02
        results[key] = value
    return True
