import foreman
def qos_unsatisfied_export_metric(params,results):
    for key, _ in params.items():
        res_json = foreman.execute_query('SELECT * FROM METRICS WHERE name == %s' % key)
        metrics_val = res_json[0]['datapoints'][-1][0]
        foreman.set_register('%s_metric' % key, str(metrics_val))
    return True
