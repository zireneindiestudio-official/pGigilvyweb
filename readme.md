# Bookings and Reservations

This is the repository for my bookings and reservations project.

- Built in Go version 1.15
- Uses the [chi router](https://github.com/go-chi/chi)
- Uses [alex edwards SCS](https://github.com/alexedwards/scs/v2) session management
- Uses [nosurf](https://github.com/justinas/nosurf)

# Hai NguyenVan

Project được xây dựng từ một project từ khoá học trên udemy trước đây mình có học nên vẫn còn nhiều code cũ mình chưa kịp xoá, đang để tham khảo code. Những function có từ khoá liên quan đến các template & link này là của khoá học:
- http://127.0.0.1:8080/
- http://127.0.0.1:8080/home
- http://127.0.0.1:8080/about
- http://127.0.0.1:8080/generals-quarters
- http://127.0.0.1:8080/majors-suite
- http://127.0.0.1:8080/search-availability
- http://127.0.0.1:8080/contact

Các lệnh để chạy localhost & setup DB (mình hay chạy lệnh trên bash Terminal)
- Lệnh chạy chính:
    $ chmod +x run.sh (một lần khi mở project trên máy). Rồi chạy lệnh "./run.sh", sau này chỉ cần chạy ".run.sh"
- Lệnh chạy setup postgres DB:
    Mình sử dụng  Soda của Buffalo (https://gobuffalo.io/en/docs/db/migrations/)
    Lệnh chạy: 
        $ soda migrate
        $ soda migrate up
    Lệnh gỡ:
        $ soda migrate down
    (Nhớ tạo một Bookings rỗng trước trong postgres DB)
    //"20211231060246_create_create_product_tables.up.fizz" đang không đúng kiểu dữ liệu như file models của product, mình đang chỉnh tay trong database qua DBeaver (driver quản lí postgres trên windown) nên chưa chỉnh file này lại cho chuẩn.


    
