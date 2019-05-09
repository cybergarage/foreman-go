import foreman
def qos_unsatisfied_export(params,results):
    names = ["m01", "m02"]
    for name in names:
        # Export from metrics store to registry
        res_json = foreman.execute_query('SELECT * FROM METRICS WHERE name == %s' % name)
        metrics_val = res_json[0]['datapoints'][-1][0]
        foreman.set_register('%s_metric' % name, str(metrics_val))
        # Export from registry store to registry
        res_json = foreman.execute_query('EXPORT FROM REGISTER WHERE name == %s' % name)
        register_val = res_json["register"][name]["value"]
        foreman.set_register('%s_register' % name, str(register_val))
    return True
