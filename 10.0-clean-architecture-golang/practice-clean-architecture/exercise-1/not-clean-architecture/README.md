[Ứng dụng Clean Architecture cho service Golang REST API](https://200lab.io/blog/ung-dung-clean-architecture-service-golang-rest-api)

---

# Ứng dụng Clean Architecture cho service Golang REST API

## Vấn đề của code gốc

- Toàn bộ code nằm trong một file main.go duy nhất
- Không có sự phân chia và tổ chức source code
- Mỗi handler đảm nhận nhiều nhiệm vụ: xử lý request, kiểm tra dữ liệu, thao tác DB, trả về JSON
- Hậu quả: khó bảo trì, dễ xung đột khi làm việc nhóm, khó thực hiện unit test

## Thiết kế kiến trúc mới

Áp dụng Clean Architecture với 3 tầng chính:

1. **Transport** : tiếp nhận HTTP request, xử lý data, trả về JSON cho client
2. **Business** : thực hiện logic nghiệp vụ
3. **Storage** : lưu trữ và truy xuất dữ liệu

## Cấu trúc thư mục đề xuất

```
modules/
  item/
    model/     (chứa model dữ liệu)
    transport/ (tầng tiếp nhận request)
    biz/       (tầng xử lý logic)
    storage/   (tầng lưu trữ dữ liệu)
```

## Triển khai thực tế

Bài viết minh họa cách áp dụng kiến trúc này với API Create Item:

1. **Model** : định nghĩa cấu trúc dữ liệu ToDoItem
2. **Business** : định nghĩa interface lưu trữ, logic kiểm tra dữ liệu
3. **Storage** : triển khai interface lưu trữ với MySQL/GORM
4. **Transport** : tiếp nhận request, xử lý dữ liệu và kết nối các tầng

## Điểm quan trọng

- Sử dụng interface để tạo sự độc lập giữa các tầng
- Áp dụng nguyên tắc "Interface dùng ở đâu thì khai báo ở đó"
- Encapsulation (bao đóng) các thành phần trong mỗi tầng
- Tầng Business không phụ thuộc vào chi tiết triển khai của Storage

Kết quả là có được mã nguồn có cấu trúc rõ ràng, dễ mở rộng, dễ test và bảo trì trong tương lai.
