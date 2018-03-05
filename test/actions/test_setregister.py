import foreman
def test_setregister(params,results):
    for key, value in params.iteritems():
        foreman.set_register(key, str(value))
    return True
