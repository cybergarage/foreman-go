import foreman
def qos_unsatisfied_export_register(params,results):
    for key, _ in params.items():
        res_json = foreman.execute_query('EXPORT FROM REGISTER WHERE name == %s' % key)
        register_val = res_json["register"][key]["value"]
        foreman.set_register('%s_register' % key, str(register_val))
    return True
