    1.  Redis:
            Untuk Redis, karena Redis tidak mempunyai fitur built-in untuk query SQL,

    2.  Neo4j:
            Untuk Neo4j, query SELECT * FROM users; dapat ditulis sebagai berikut:

                MATCH (u:User)
                RETURN u;
                Perintah ini akan mengambil semua node yang memiliki label "User" pada database Neo4j.

    3.  Cassandra:
            Untuk Cassandra, query SELECT * FROM users; dapat ditulis sebagai berikut:

            SELECT * FROM users;
            Perintah ini akan mengambil semua data pada tabel "users" pada database Cassandra. 