import foreman
def export_register(params,results):
        for key, _ in params.items():
                export_key = "export_%s" % (key)
                results[export_key] = foreman.get_register(key)
        return True    
