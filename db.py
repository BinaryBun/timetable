from getpass import getpass
from mysql.connector import connect, Error

try:
	with connect(host="127.0.0.1",
				 user="",
				 password="") as connection:

		with connection.cursor() as cursor:
			cursor.execute("USE parser;")
			cursor.execute("SELECT * from passwdords;")
			for data in cursor:
				print(data)

except Error as e:
	print(f"Error: {e}")
