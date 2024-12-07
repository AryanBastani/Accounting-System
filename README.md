# Simple Accounting System Project

## Project Objective
The goal of this project is to create an accounting system that allows users to manage *DL (Detail Ledger)* and *SL (Subsidiary Ledger)* entities and use them to issue accounting vouchers. The system includes full CRUD functionality for each entity, validation rules, and tests to ensure proper behavior.

---

## Entities and Features

### 1. **Detail Ledger (DL)**  
- **Fields:**  
  - `Code (string)`  
  - `Title (string)`  

- **Validation Rules:**  
  - `Code` and `Title` cannot be empty.  
  - Maximum length: 64 characters for both `Code` and `Title`.  
  - Both `Code` and `Title` must be unique.  
  - Deletion is not allowed if referenced in a voucher.  

---

### 2. **Subsidiary Ledger (SL)**  
- **Fields:**  
  - `Code (string)`  
  - `Title (string)`  
  - `IsDetailable (boolean)`  

- **Validation Rules:**  
  - `Code` and `Title` cannot be empty.  
  - Maximum length: 64 characters for both `Code` and `Title`.  
  - Both `Code` and `Title` must be unique.  
  - Deletion or editing is not allowed if referenced in a voucher.  

---

### 3. **Accounting Voucher**  
- **Fields:**  
  - `Number (string)`  

- **Validation Rules:**  
  - `Number` cannot be empty.  
  - Maximum length: 64 characters.  
  - `Number` must be unique.  

---

### 4. **Voucher Item**  
- **Fields:**  
  - `SummaryID (int)`  
  - `DetailID (nullable int)`  
  - `Debit (int)`  
  - `Credit (int)`  

- **Validation Rules:**  
  - `SummaryID` is mandatory and must reference an existing record.  
  - If `SummaryID` is detailable, selecting a `DetailID` is mandatory.  
  - If `SummaryID` is not detailable, `DetailID` must be empty.  
  - If `DetailID` is selected, it must reference an existing record.  
  - Either `Debit` or `Credit` must be greater than 0 (but not both).  
  - A voucher must have a minimum of 2 and a maximum of 500 items.  
  - All voucher items must balance, meaning the total `Debit` equals the total `Credit`.  

---

## Project Notes
1. **Code Quality & Testing:**  
   The correctness of the program, clean code, and comprehensive testing are critical.  

2. **Validation Messages:**  
   Error messages must be clear and descriptive.  

3. **Versioning:**  
   Use a `Version` field in entities (DL, SL, Voucher) to ensure consistency during updates and deletions.  

4. **Testing Requirements:**  
   - Tests should not rollback database changes, allowing repeated test execution without errors.  
   - BDD tests must avoid mocking the database.  

5. **Submission Requirements:**  
   - Upload the project files, along with an explanatory voice or video recording, as a single zipped file.  

--- 
