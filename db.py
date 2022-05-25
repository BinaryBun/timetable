from getpass import getpass
from mysql.connector import connect, Error

class DB():
    def __init__(self):
        pass

    def insert(self, data, group):
        day = data[0]
        try:
            with connect(host="127.0.0.1",
                         user="root",
                         password="binarybun",
                         autocommit=True) as connection:
                with connection.cursor() as cursor:
                    cursor.execute("USE timetable;")
                    req_1 = "insert into `table` (`group`, `day`, nuber, `time`, audit, `name`, prepod, `data`, map)"
                    for i in data[1:]:
                        req_2 = f'("{group}", "{day}", "-", "{i[0]}", "{i[1]}", "{i[2]}", "{i[3]}", "{i[4]}", "-")'
                        print(f"{req_1} value {req_2};")
                        #cursor.execute(f"{req_1} value {req_2};")

        except Error as e:
            return f"Error: {e}"
