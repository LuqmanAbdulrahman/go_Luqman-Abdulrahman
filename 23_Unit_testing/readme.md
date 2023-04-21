UNIT TESTING

Unit test di Go dituliskan dalam bentuk fungsi, yang memiliki parameter yang bertipe *testing.T, dengan nama fungsi harus diawali kata Test. Lewat parameter tersebut, kita bisa mengakses method-method untuk keperluan testing.

go menyediakan package testing, berisikan banyak sekali tools untuk keperluan unit test.

Pada section ini belajar mengenai testing, benchmark, dan juga testing menggunakan testify