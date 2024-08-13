# BookHub

## 1. Quản lý Sách (Books Management)
### API:
- **GET /api/books**: Lấy danh sách tất cả các sách.
- **GET /api/books/{id}**: Lấy thông tin chi tiết của một sách.
- **POST /api/books**: Thêm một sách mới vào thư viện.
- **PUT /api/books/{id}**: Cập nhật thông tin của một sách.
- **DELETE /api/books/{id}**: Xóa một sách khỏi thư viện.

## 2. Quản lý Tác giả (Authors Management)
### API:
- **GET /api/authors**: Lấy danh sách tất cả các tác giả.
- **GET /api/authors/{id}**: Lấy thông tin chi tiết của một tác giả.
- **POST /api/authors**: Thêm một tác giả mới.
- **PUT /api/authors/{id}**: Cập nhật thông tin của một tác giả.
- **DELETE /api/authors/{id}**: Xóa một tác giả khỏi hệ thống.

## 3. Quản lý Danh mục (Categories Management)
### API:
- **GET /api/categories**: Lấy danh sách tất cả các danh mục.
- **POST /api/categories**: Thêm một danh mục mới.
- **PUT /api/categories/{id}**: Cập nhật thông tin của một danh mục.
- **DELETE /api/categories/{id}**: Xóa một danh mục khỏi hệ thống.

## 4. Quản lý Thành viên (Members Management)
### API:
- **GET /api/members**: Lấy danh sách tất cả các thành viên.
- **GET /api/members/{id}**: Lấy thông tin chi tiết của một thành viên.
- **POST /api/members**: Thêm một thành viên mới.
- **PUT /api/members/{id}**: Cập nhật thông tin của một thành viên.
- **DELETE /api/members/{id}**: Xóa một thành viên khỏi hệ thống.

## 5. Quản lý Mượn sách (Loan Management)
### API:
- **GET /api/loans**: Lấy danh sách tất cả các bản ghi mượn sách.
- **GET /api/loans/{id}**: Lấy thông tin chi tiết của một bản ghi mượn sách.
- **POST /api/loans**: Tạo một bản ghi mượn sách mới.
- **PUT /api/loans/{id}**: Cập nhật thông tin của một bản ghi mượn sách.
- **DELETE /api/loans/{id}**: Xóa một bản ghi mượn sách khỏi hệ thống.

## 6. Tìm kiếm và Thống kê (Search and Reports)
### API:
- **GET /api/search**: Tìm kiếm sách, tác giả, danh mục, v.v.
- **GET /api/reports/overdue**: Báo cáo các sách quá hạn.
- **GET /api/reports/popular**: Báo cáo các sách được mượn nhiều nhất.