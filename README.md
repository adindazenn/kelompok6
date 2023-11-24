# # Finalproject3 Hacktiv8 MSIB5 - Kanban Board
## Kelompok 6
#### Tech Stack
- Languange: Go
- Framework : GIN
- Database : Postgre
- ORM : GORM
- Token : JWT
- Deploy : Railway

Link : https://finalproject3-kelompok6.up.railway.app

Dokumentasi API

A. Users

1. POST /users/register
   
- request body 
```
{
	"full_name": "string",
	"email": "string",
	"password": "string"
}
```

- response
- status 201
- data :
```
{
	"id": "integer",
	"full_name": "string":,
	"email": "string",
	"created_at": "date"
}
```
note untuk endpoint ini, role dari data user akan otomatis menjadi member. boleh langsung dihardcode di controller sebelum disimpan ke database

3. POST /users/login
- request body
```
{
	"email": "string",
	"password": "string"
}
```
- response
- status 200
```
data :
{
	"token": "jwt string"
}
```
note untuk endpoint ini wajib melakukan pengecekan email dan password user. pengecekan password wajib dengan library/package Bcrypt 

4. PUT /users/update-account
- request headers:  Authorization (Bearer token string)
- request body :
```
{
	"full_name": "string",
	"email": "string"
}
```
- response
- status 200
- body :
```
{
	"id": "integer",
	"full_name": "string",
	"email": "string",
	"updated_at": "date"
}
```
note endpoint ini wajib melakukan autentikasi dengan package/library jsonwebtoken.

5. DELETE /users/delete-account
- request headers: Authorization (Bearer token string)

- response
- status 200
- body :
```
{
	"message": "Your account has been succesfully deleted"
}
```
note endpoint ini ajib melakukan autentikasi dengan package/library jsonwebtoken.

B. Categories
1. POST /categories
- request 
- headers: Authorization (Bearer token string)
- body :
```
{
	"type": "string"
}
```

- response 
- status 201
- data :
```
{
	"id": "integer"
	"type": "string"
	"created_at": "date"
}
```

2. GET /categories
- request
- headers: Authorization (Bearer token string)
- response
- status 200
- data :
```
{
	{
		"id": "integer"
		"type": "string"
		"updated_at": "date"
		"created_at": "date"
		"Tasks": [
			{
				"id": "integer"
				"title": "string"
				"description": "string"
				"user_id": "integer"
				"category_id": "integer"
				"created_at": "date"
				"updated_at": "date"
			}
		]
	}
}
```

3. PATCH /categories/:categoryId
- request
- headers: Authorization (Bearer token string)
- params: categoryId (integer)
- body:
```
{
	"type": "string"
}
```
- response
- status 200
- data :
```
{
	"id": "integer"
	"type": "string"
	"sold_product_amount": "integer"
	"updated_at": "date"
}
```

4. DELETE /categories/:categoryId
- request
- headers: Authorization (Bearer token string)
- params: categoryId (integer)
- response
- status 200
- data :
```
{
	"message": "Category has been successfully deleted"
}
```

C. Tasks 
1. POST /tasks
- request 
- header : Authorization (Bearer token string)
- body : 
```
{
	"title": "string",
	"description": "string",
	"category_id": "integer"
}
```
- response :
- status 201
- data :
```
{
	"id": "integer",
	"title": "string",
	"status": "boolean",
	"description": "string",
	"user_id": "integer",
	"category_id": "integer",
	"created_at": "date"
}
```
note : pada endpoint ini harus dilakukan pengecekan data category dengan id yang diberikan pada request body dengan field categoryid ada atau tidak pada database. jika ada maka boleh disimpan di database namun jika tidak ada maka harus melempar error.  kemudian untuk field status nilai awalnya akan otomatis menjadi false, bisa langsung hardcode di controller.

2. GET /tasks
- request 
- header : Authorization (Bearer token string)
- response :
- status 200
- data :
```
{
	"id": "integer",
	"title": "string",
	"status": "boolean",
	"description": "string",
	"user_id": "integer",
	"category_id": "integer",
	"created_at": "date",
	"User": {
		"id": "integer",
		"email": "string",
		"full_name": "string"
	}
}
```

3. PUT /tasks/:taskId
- request 
- header : Authorization (Bearer token string)
- params : taskId (integer)
- body :
```
{
	"title": "string",
	"description": "string"
}
```
- response
- status 200
- data :
```
{
	"id": "integer",
	"title": "string",
	"description": "string",
	"status": "boolean",
	"user_id": "integer",
	"category_id": "integer",
	"updated_at": "date"
}
```
note : pada endpoint ini, perlu dilakukan proses autorisasi dimana user hanya boleh mengupdate task miliknya sendiri

4. PATCH /tasks/update-status/:taskId
- request 
- header : Authorization (Bearer token string)
- params : taskId (integer)
- body :
```
{
	"status": "boolean"
}
```
- response :
- status 200
- data :
```
{
	"id": "integer",
	"title": "string",
	"description": "string",
	"status": "boolean",
	"user_id": "integer",
	"category_id": "integer",
	"updated_at": "date"
}
```
note : pada endpoint ini, perlu dilakukan proses autorisasi dimana user hanya boleh mengupdate status task miliknya sendiri

5. PATCH /tasks/update-category/:taskId
- request 
- header : Authorization (Bearer token string)
- params : taskId (integer)
- body :
```
{
	"category_id": "integer"
}
```
- response :
- status 200
- data :
```
{
	"id": "integer",
	"title": "string",
	"description": "string",
	"status": "boolean",
	"user_id": "integer",
	"category_id": "integer",
	"updated_at": "date"
}
```
note : pada endpoint ini, perlu dilakukan proses autorisasi dimana user hanya boleh mengupdate category id task miliknya sendiri. lalu perlu dilakuakn pengecekan jika category dengan id yang diinput user ada dalam database, maka proses dilanjut. namun jika tidak ada langsung melempar error.

6. DELETE /tasks/:taskId
- request 
- header : Authorization (Bearer token string)
- params : taskId (integer)
- response
- status 200
- data :
```
{
	"message": "Task has been succesfully deleted"
}
```
note : pada endpoint ini, perlu dilakukan proses autorisasi dimana user hanya boleh menghapus task miliknya sendiri.

  
