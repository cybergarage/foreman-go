import foreman
# The action (test_qos_action_03) is not executed because the destination action (test_qos_action_02) returns a false
def test_qos_action_03(params,results):
    for key, value in params.iteritems():
        chkKey = 'qos_action_03_%s' % key
        foreman.set_register(chkKey, str(value))
    return True