# Golang upload image using 3 different approaches

This is a simple RESTful API for managing image uploads. It provides three different approaches for uploading images:

1. <b>Direct Upload:</b> Upload images directly via a form submission.

2. <b>Base64 Upload:</b> Upload images encoded in Base64 format.

3. <b>Multipart Upload:</b> Upload multiple images using multipart form data.

## Pros and Cons of Each Approach

| Approach         | Pros                           | Cons                                          |
| ---------------- | ------------------------------ | --------------------------------------------- |
| Direct Upload    | Simple and straightforward     | Limited browser support for large files       |
| Base64 Upload    | No file system access required | Increased payload size due to Base64 encoding |
| Multipart Upload | Supports multiple file uploads | Slightly more complex to implement            |

## Usage

### Direct Upload

Endpoint: `POST /upload-image/direct`

Uploads an image file directly via a form submission.

#### Request:

```bash
curl -X POST -F "image=@/path/to/image.jpg" http://localhost:8080/upload-image/direct
```

#### Response:

```json
{
  "message": "File uploaded successfully"
}
```

### Base64 Upload

Endpoint: `POST /upload-image/base64/:image-name`

Uploads an image encoded in Base64 format.

#### Request:

```bash
curl -X POST -F "image=base64string" http://localhost:8080/upload-image/base64/myimage
```

#### Response:

```json
{
  "message": "Image uploaded successfully"
}
```

### Multipart Upload

Endpoint: `POST /upload-image/multipart`

Uploads multiple images using multipart form data.

#### Request:

```bash
curl -X POST -F "images=@/path/to/image1.jpg" -F "images=@/path/to/image2.jpg" http://localhost:8080/upload-image/multipart
```

#### Response:

```json
{
  "message": "Files uploaded successfully"
}
```
