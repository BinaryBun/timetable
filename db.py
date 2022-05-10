from getpass import getpass
from mysql.connector import connect, Error

class DB():
    def __init__(self):
        pass

    def insert(self, data):
        print(data)
        try:
        	with connect(host="127.0.0.1",
        				 user="",
        				 password="") as connection:
        		with connection.cursor() as cursor:
        			'''
                        request
                    '''

        except Error as e:
        	return f"Error: {e}"
