-- Tạo bảng Categories
CREATE TABLE Categories (
    CategoryID INT AUTO_INCREMENT PRIMARY KEY,
    CategoryName VARCHAR(255) NOT NULL
);

-- Tạo bảng Authors
CREATE TABLE Authors (
    AuthorID INT AUTO_INCREMENT PRIMARY KEY,
    FirstName VARCHAR(255) NOT NULL,
    LastName VARCHAR(255) NOT NULL
);

-- Tạo bảng Books
CREATE TABLE Books (
    BookID INT AUTO_INCREMENT PRIMARY KEY,
    Title VARCHAR(255) NOT NULL,
    ISBN VARCHAR(13) NOT NULL,
    PublishedYear INT NOT NULL,
    CategoryID INT,
    FOREIGN KEY (CategoryID) REFERENCES Categories(CategoryID)
);

-- Tạo bảng BookAuthors
CREATE TABLE BookAuthors (
    BookID INT,
    AuthorID INT,
    FOREIGN KEY (BookID) REFERENCES Books(BookID),
    FOREIGN KEY (AuthorID) REFERENCES Authors(AuthorID),
    PRIMARY KEY (BookID, AuthorID)
);

-- Tạo bảng Members
CREATE TABLE Members (
    MemberID INT AUTO_INCREMENT PRIMARY KEY,
    FirstName VARCHAR(255) NOT NULL,
    LastName VARCHAR(255) NOT NULL,
    Email VARCHAR(255) NOT NULL,
    PhoneNumber VARCHAR(20),
    JoinDate DATE NOT NULL
);

-- Tạo bảng Loans
CREATE TABLE Loans (
    LoanID INT AUTO_INCREMENT PRIMARY KEY,
    BookID INT,
    MemberID INT,
    LoanDate DATE NOT NULL,
    DueDate DATE NOT NULL,
    ReturnDate DATE,
    FOREIGN KEY (BookID) REFERENCES Books(BookID),
    FOREIGN KEY (MemberID) REFERENCES Members(MemberID)
);

-- Tạo bảng Reviews
CREATE TABLE Reviews (
    ReviewID INT AUTO_INCREMENT PRIMARY KEY,
    BookID INT,
    MemberID INT,
    ReviewDate DATE NOT NULL,
    Rating INT NOT NULL CHECK (Rating >= 1 AND Rating <= 5),
    Comment TEXT,
    FOREIGN KEY (BookID) REFERENCES Books(BookID),
    FOREIGN KEY (MemberID) REFERENCES Members(MemberID)
);

ALTER TABLE Books ADD COLUMN Deleted BOOLEAN DEFAULT FALSE;
ALTER TABLE Authors ADD COLUMN Deleted BOOLEAN DEFAULT FALSE;
ALTER TABLE Categories ADD COLUMN Deleted BOOLEAN DEFAULT FALSE;
ALTER TABLE Members ADD COLUMN Deleted BOOLEAN DEFAULT FALSE;
ALTER TABLE Loans ADD COLUMN Deleted BOOLEAN DEFAULT FALSE;
ALTER TABLE Reviews ADD COLUMN Deleted BOOLEAN DEFAULT FALSE;
ALTER TABLE BookAuthors ADD COLUMN Deleted BOOLEAN DEFAULT FALSE;