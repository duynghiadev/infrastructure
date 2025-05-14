Dựa vào cấu trúc dự án tôi thấy trong thư mục /internal , các route được tổ chức theo từng module và nằm trong thư mục delivery của mỗi module. Cụ thể:

1. Module Auth:

- /internal/auth/delivery/ - chứa các route liên quan đến xác thực

2. Module Todos:

- /internal/todos/delivery/ - chứa các route liên quan đến todos

3. File handler chính:

- /internal/server/handler.go - file này chứa hàm MapHandlers để đăng ký tất cả các routes
  Cấu trúc này tuân theo Clean Architecture, trong đó:
- Mỗi module (auth, todos) có thư mục riêng
- Trong mỗi module có các layer:

  - delivery: xử lý HTTP routes và requests
  - repository: tương tác với database
  - usecase: business logic
  - presenter: format response data
    Để xem chi tiết các routes, bạn cần kiểm tra:

1. /internal/auth/delivery/ cho các API authentication
2. /internal/todos/delivery/ cho các API todos
3. /internal/server/handler.go để xem cách các routes được map
   Đây là cấu trúc chuẩn của một project Go theo Clean Architecture, giúp code được tổ chức rõ ràng và dễ bảo trì.
