package entity

//Represents leave_records table on the database
type LeaveRecord struct {
	Request_id  int    `json:"request_id"`
	Request_on  string `json:"request_on"`
	From_date   string `json:"from_date"`
	To_date     string `json:"to_date"`
	Return_date string `json:"return_date"`
	Reason      string `json:"reason"`
	Mobile      string `json:"mobile"`
	Address     string `json:"address"`
	Status_id   int    `json:"status_id"`
	Leave_id    int    `json:"leave_id"`
	User_id     int    `json:"user_id"`
}

//View model for leave_records and status table
type LeaveRecordListModel struct {
	Request_id  int    `json:"request_id"`
	Request_on  string `json:"request_on"`
	Leave_name  string `json:"leave_type"`
	Reason      string `json:"reason"`
	Status_name string `json:"status"`
	User_id     int    `json:"user_id"`
}

// -- TABLE DEFINITION FOR be_payroll db

// CREATE TABLE
// 	users (
// 		user_id serial PRIMARY KEY,
// 		name VARCHAR(50) NOT NULL,
// 		password VARCHAR NOT NULL,
// 		user_mail VARCHAR(50) UNIQUE NOT NULL,
// 		nik VARCHAR(20) NOT NULL,
// 		role_id INT NOT NULL,
// 		position_id INT NOT NULL
// 	);

// CREATE TABLE
// 	roles (
// 		role_id INT PRIMARY KEY,
// 		role_name VARCHAR(10) NOT NULL
// 	);

// CREATE TABLE
// 	positions (
// 		position_id INT PRIMARY KEY,
// 		position_name VARCHAR(30) NOT NULL
// 	);

// CREATE TABLE
// 	leave_records (
// 		request_id SERIAL PRIMARY KEY,
// 		request_on DATE NOT NULL,
// 		from_date DATE NOT NULL,
// 		to_date DATE NOT NULL,
// 		return_date DATE NOT NULL,
// 		reason VARCHAR(100) NOT NULL,
// 		mobile 	VARCHAR(15) NOT NULL,
// 		address VARCHAR(150) NOT NULL,
// 		status_id INT NOT NULL,
// 		leave_id INT NOT NULL,
// 		user_id INT NOT NULL
// 	);

// CREATE TABLE
// 	leave_types (
// 		leave_id INT PRIMARY KEY NOT NULL,
// 		leave_name VARCHAR(10) NOT NULL
// 	);

// CREATE TABLE
// 	status (
// 		status_id INT PRIMARY KEY NOT NULL,
// 		status_name VARCHAR(15) NOT NULL
// 	);

// CREATE TABLE
// 	leave_balance (
// 		balance_id SERIAL PRIMARY KEY,
// 		balance_year VARCHAR(4) NOT NULL,
// 		cuti_tahunan INT,
// 		cuti_diambill INT,
// 		cuti_balance INT,
// 		cuti_izin INT,
// 		cuti_sakit INT,
// 		user_id INT NOT NULL
// 	);

// SELECT * FROM users;
// SELECT * FROM roles;
// SELECT * FROM positions;
// SELECT * FROM roles;
// SELECT * FROM leave_records;
// SELECT * FROM leave_types;
// SELECT * FROM status;
// SELECT * FROM leave_balance;

// INSERT INTO
// 	roles (role_id, role_name)
// VALUES
// 	(1, 'admin'),
// 	(2, 'employee');

// INSERT INTO
// 	leave_types (leave_id, leave_name)
// VALUES
// 	(1, 'cuti'),
// 	(2, 'izin'),
// 	(3, 'sakit');

// INSERT INTO
// 	status (status_id, status_name)
// VALUES
// 	(1, 'Waiting Approval'),
// 	(2, 'Approved'),
// 	(3, 'Rejected'),
// 	(4, 'Scheduled'),
// 	(5, 'Paid')
// 	;

// ALTER TABLE status
// ALTER COLUMN status_name TYPE VARCHAR(20);