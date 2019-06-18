import foreman
def recursive_exec(params,results):
        for key, value in params.items():
                query = 'EXECUTE post_register (%s) VALUES ("%s")' % (key, value)
                print(query)
                foreman.execute_query(query)
        return True    
